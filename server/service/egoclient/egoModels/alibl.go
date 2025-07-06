package egoModels

import (
	"context"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
	egoclientReq "github.com/flipped-aurora/gin-vue-admin/server/model/egoclient/request"
	"github.com/liusuxian/go-aisdk/consts"
	"github.com/liusuxian/go-aisdk/httpclient"
	"github.com/liusuxian/go-aisdk/models"
	"strings"
	"time"
)

func init() {
	RegisterService("alibl", func() Service {
		return NewAliBLService()
	})
}

type AliBLService struct {
	BasicService
}

func NewAliBLService() *AliBLService {
	ds := &AliBLService{}
	ds.initAssemblers()
	return ds
}

func (s *AliBLService) initAssemblers() {
	s.ModelHandlers = map[consts.ModelType]map[string]*ModelHandler{
		consts.ChatModel: {
			consts.AliBLQwqPlus: &ModelHandler{s.AliBLAssemble, s.AliBLChatHandler},
			consts.AliBLQvqMax:  &ModelHandler{s.AliBLAssemble, s.AliBLChatHandler},
		},
	}
}

func CheckModalValid(modelName string, toMatch ...string) bool {
	for _, match := range toMatch {
		if strings.Contains(modelName, match) {
			return true
		}
	}
	return false
}

func (s *AliBLService) ParseChatModal(ModelName string, Text string, modals []egoclientReq.EgoDialogueMultiModal) (*models.UserMessage, error) {
	userMsg := &models.UserMessage{}

	if len(modals) == 0 {
		userMsg.Content = Text
	} else {
		userMsg.MultimodalContent = append(userMsg.MultimodalContent, models.ChatUserMsgPart{
			Type: models.ChatUserMsgPartTypeText,
			Text: Text,
		})
	}

	for _, modal := range modals {
		userMsgPart := models.ChatUserMsgPart{
			Type: modal.Type,
		}

		switch modal.Type {
		case models.ChatUserMsgPartTypeText:
			if CheckModalValid(ModelName, "qwen-vl", "qvq", "qwen-omni") != true {
				return nil, errors.New(fmt.Sprintf("该模型不支持多模态类型 %s", modal.Type))
			}
			userMsgPart.Text = modal.Text
		case models.ChatUserMsgPartTypeImageURL:
			if CheckModalValid(ModelName, "qwen-vl", "qvq", "qwen-omni") != true {
				return nil, errors.New(fmt.Sprintf("该模型不支持多模态类型 %s", modal.Type))
			}
			userMsgPart.ImageURL = &models.ChatUserMsgImageURL{URL: modal.Url}
		default:
			return nil, errors.New("不支持的多模态类型")
		}

		userMsg.MultimodalContent = append(userMsg.MultimodalContent, userMsgPart)
	}

	return userMsg, nil
}

func (s *AliBLService) AliBLAssemble(ED *egoclient.EgoDialogue, Req *egoclientReq.EgoDialoguePostUserMsg) (httpclient.Response, error) {
	if Req.ChatOption == nil {
		return nil, errors.New("错误的请求格式")
	}
	ctx := context.WithValue(context.Background(), "Dialogue", ED)
	chatReq := models.ChatRequest{
		//UserInfo: models.UserInfo{
		//	User: *ED.User.UserID,
		//},
		//Provider:            ED.Model.ModelProvider,
		//Model:               *ED.Model.ModelName,
		//MaxCompletionTokens: models.Int(4096),
		//FrequencyPenalty:    models.Float32(1.0),
		//EnableThinking:      models.Bool(Req.ChatOption.Reasoning), //Qwen3 默认开启thinking
		//Stream:              models.Bool(true),                     // 不会被序列化，会放到请求头中
		//StreamOptions: &models.ChatStreamOptions{
		//	IncludeUsage: models.Bool(true),
		//}, // 不会被序列化
		UserInfo: models.UserInfo{
			User: "123456",
		},
		Provider:            consts.AliBL,
		Model:               consts.AliBLQwenVlMax,
		FrequencyPenalty:    models.Float32(1.0),
		MaxCompletionTokens: models.Int(4096),
		// Metadata:            map[string]string{"X-DashScope-DataInspection": "{\"input\": \"cip\", \"output\": \"cip\"}"},
		WebSearchOptions: &models.ChatWebSearchOptions{
			EnableSource:   models.Bool(true),
			EnableCitation: models.Bool(true),
			CitationFormat: models.ChatCitationFormatRefNumber,
			ForcedSearch:   models.Bool(true),
			SearchStrategy: models.ChatSearchStrategyPro,
		},
		EnableThinking: models.Bool(true),
		Stream:         models.Bool(true), // 不会被序列化，会放到请求头中
		StreamOptions: &models.ChatStreamOptions{
			IncludeUsage: models.Bool(true),
		}, // 不会被序列化
	}

	//插入历史消息
	for _, v := range ED.Histories {
		chatReq.Messages = append(chatReq.Messages, v.Role.GetMessage(v.Content, v.ReasoningContent))
	}

	//插入用户当前消息
	var userMsg *models.UserMessage
	var err error
	if userMsg, err = s.ParseChatModal(*ED.Model.ModelName, Req.Text, Req.ChatOption.Multimodal); err != nil {
		return nil, err
	}
	chatReq.Messages = append(chatReq.Messages, userMsg)

	return global.AiSDK.CreateChatCompletionStream(ctx, chatReq, httpclient.WithTimeout(time.Minute*5), httpclient.WithStreamReturnIntervalTimeout(time.Second*5))
}

func (s *AliBLService) AliBLChatHandler(ctx context.Context, DialogueID uint) func(item models.ChatBaseResponse, isFinished bool) error {
	var Contents []ChatStreamContentBlock
	var DialogueItem egoclient.EgoDialogueItem
	return func(item models.ChatBaseResponse, isFinished bool) (err error) {
		if isFinished {
			// 存储历史记录
			for _, v := range Contents {
				if err = ModelSer.CreateEgoDialogueHistory(ctx, &egoclient.EgoDialogueHistory{
					Role:             egoclient.AssistantRole,
					Item:             DialogueItem.UUID,
					DialogueID:       DialogueID,
					ReasoningContent: v.ReasoningBuffer.String(),
					Content:          v.ContentBuffer.String(),
					IsChoice:         true,
				}); err != nil {
					return err
				}
			}

			// 存储token用量
			if err = ModelSer.CreateEgoDialogueItem(ctx, &DialogueItem); err != nil {
				return
			}
			return nil
		}
		for _, v := range item.Choices {
			for v.Index >= len(Contents) {
				Contents = append(Contents, ChatStreamContentBlock{})
			}
			Contents[v.Index].ContentID = item.ID
			Contents[v.Index].ReasoningBuffer.WriteString(v.Delta.ReasoningContent)
			Contents[v.Index].ContentBuffer.WriteString(v.Delta.Content)
		}
		if item.Usage != nil && item.StreamStats != nil {
			//更新token用量
			DialogueItem.UUID = item.ID
			DialogueItem.PromptTokens = item.Usage.PromptTokens
			DialogueItem.DialogueID = DialogueID
			DialogueItem.CompletionTokens = item.Usage.CompletionTokens
			//log.Printf("createChatCompletionStream usage = %+v", item.Usage)
			//log.Printf("createChatCompletionStream stream_stats = %+v", item.StreamStats)
		}
		return nil
	}
}
