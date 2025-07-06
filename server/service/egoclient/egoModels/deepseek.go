package egoModels

import (
	"context"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
	egoclientReq "github.com/flipped-aurora/gin-vue-admin/server/model/egoclient/request"
	"github.com/liusuxian/go-aisdk/consts"
	httpclient "github.com/liusuxian/go-aisdk/httpclient"
	"github.com/liusuxian/go-aisdk/models"
	"log"
	"time"
)

func init() {
	RegisterService("deepseek", func() Service {
		return NewDeepseekService()
	})
}

type DeepseekService struct {
	BasicService
}

func NewDeepseekService() *DeepseekService {
	ds := &DeepseekService{}
	ds.initAssemblers()
	return ds
}

func (s *DeepseekService) initAssemblers() {
	s.ModelHandlers = map[consts.ModelType]map[string]*ModelHandler{
		consts.ChatModel: {
			"any": &ModelHandler{s.DeepSeekReasonerAssemble, nil},
		},
	}
}

func (s *DeepseekService) ParseChatModal(ModelName string, Text string, modals []egoclientReq.EgoDialogueMultiModal) (*models.UserMessage, error) {
	userMsg := &models.UserMessage{
		Content: Text,
	}
	if len(modals) != 0 {
		return nil, errors.New("deepseek 不支持多模态输入")
	}

	return userMsg, nil
}

func (s *DeepseekService) DeepSeekReasonerAssemble(ED *egoclient.EgoDialogue, Req *egoclientReq.EgoDialoguePostUserMsg) (httpclient.Response, error) {
	if Req.ChatOption == nil {
		return nil, errors.New("错误的请求格式")
	}
	model := consts.DeepSeekChat
	if Req.ChatOption.Reasoning {
		model = consts.DeepSeekReasoner
	}

	ctx := context.WithValue(context.Background(), "Dialogue", ED)
	chatReq := models.ChatRequest{
		Provider: consts.DeepSeek,
		Model:    model,
		UserInfo: models.UserInfo{
			User: *ED.User.UserID,
		},
		Stream:              models.Bool(true),
		MaxCompletionTokens: models.Int(4096),
	}
	//插入历史消息
	for _, v := range ED.Histories {
		chatReq.Messages = append(chatReq.Messages, v.Role.GetMessage(v.Content, v.ReasoningContent))
	}

	//插入用户当前消息
	var userMsg *models.UserMessage
	var err error
	if userMsg, err = s.ParseChatModal("", Req.Text, Req.ChatOption.Multimodal); err != nil {
		return nil, err
	}
	chatReq.Messages = append(chatReq.Messages, userMsg)

	return global.AiSDK.CreateChatCompletionStream(ctx, chatReq, httpclient.WithTimeout(time.Minute*5), httpclient.WithStreamReturnIntervalTimeout(time.Second*5))
}

func (s *DeepseekService) DeepSeekChatHandler(ctx context.Context, DialogueID uint, Contents *[]ChatStreamContentBlock, ItemUUID *string) func(item models.ChatBaseResponse, isFinished bool) error {
	return func(item models.ChatBaseResponse, isFinished bool) (err error) {
		if isFinished {
			for _, v := range *Contents {
				if err = ModelSer.CreateEgoDialogueHistory(ctx, &egoclient.EgoDialogueHistory{
					Role:             egoclient.AssistantRole,
					Item:             *ItemUUID,
					DialogueID:       DialogueID,
					ReasoningContent: v.ReasoningBuffer.String(),
					Content:          v.ContentBuffer.String(),
					IsChoice:         true,
				}); err != nil {
					return err
				}
			}
			return nil
		}
		for _, v := range item.Choices {
			for v.Index >= len(*Contents) {
				*Contents = append(*Contents, ChatStreamContentBlock{})
			}
			(*Contents)[v.Index].ContentID = item.ID
			(*Contents)[v.Index].ReasoningBuffer.WriteString(v.Delta.ReasoningContent)
			(*Contents)[v.Index].ContentBuffer.WriteString(v.Delta.Content)
		}
		if item.Usage != nil && item.StreamStats != nil {
			//存储token用量
			*ItemUUID = item.ID
			if err = ModelSer.CreateEgoDialogueItem(ctx, &egoclient.EgoDialogueItem{
				UUID:             item.ID,
				PromptTokens:     item.Usage.PromptTokens,
				DialogueID:       DialogueID,
				CompletionTokens: item.Usage.CompletionTokens,
			}); err != nil {
				return
			}
			log.Printf("createChatCompletionStream usage = %+v", item.Usage)
			log.Printf("createChatCompletionStream stream_stats = %+v", item.StreamStats)
		}
		return nil
	}
}
