package egoModels

import (
	"context"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
	egoclientReq "github.com/flipped-aurora/gin-vue-admin/server/model/egoclient/request"
	"github.com/liusuxian/go-aisdk/consts"
	"github.com/liusuxian/go-aisdk/httpclient"
	"github.com/liusuxian/go-aisdk/models"
	"time"
)

func init() {
	RegisterService("ark", func() Service {
		return NewArkService()
	})
}

type ArkService struct {
	BasicService
}

func NewArkService() *ArkService {
	service := &ArkService{}
	service.initAssemblers()
	return service
}

func (s *ArkService) initAssemblers() {
	s.ModelHandlers = map[consts.ModelType]map[string]*ModelHandler{
		consts.ChatModel: {
			consts.Doubaoseed1_6: &ModelHandler{s.DoubaoSeedAssemble, DefaultChatHandler},
		},
		consts.ImageModel: {
			consts.Doubaoseedream3: &ModelHandler{s.DoubaoSeedReamAssemble, s.DoubaoSeedReamHandler},
		},
	}
}

func (s *ArkService) ParseChatModal(ModelName string, Text string, modals []egoclientReq.EgoDialogueMultiModal) (*models.UserMessage, error) {
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
			userMsgPart.Text = modal.Text
		case models.ChatUserMsgPartTypeImageURL:
			userMsgPart.ImageURL = &models.ChatUserMsgImageURL{URL: modal.Url}
		case models.ChatUserMsgPartTypeVideoURL:
			userMsgPart.VideoURL = &models.ChatUserMsgVideoURL{URL: modal.Url}
		default:
			return nil, errors.New("不支持的多模态类型")
		}

		userMsg.MultimodalContent = append(userMsg.MultimodalContent, userMsgPart)
	}

	return userMsg, nil
}

func (s *ArkService) DoubaoSeedAssemble(ED *egoclient.EgoDialogue, Req *egoclientReq.EgoDialoguePostUserMsg) (httpclient.Response, error) {
	if Req.ChatOption == nil {
		return nil, errors.New("错误的请求格式")
	}
	ctx := context.WithValue(context.Background(), "Dialogue", ED)
	chatReq := models.ChatRequest{
		Provider: ED.Model.ModelProvider,
		Model:    *ED.Model.ModelName,
		UserInfo: models.UserInfo{
			UserID: *ED.User.UserID,
		},
		Stream:              true,
		MaxCompletionTokens: 4096,
		StreamOptions: &models.ChatStreamOptions{
			IncludeUsage: true,
		},
	}

	if Req.ChatOption.Reasoning == true {
		chatReq.Thinking = &models.ChatThinkingOptions{
			Type: Req.ChatOption.ReasoningMode,
		}
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

func (s *ArkService) DoubaoSeedReamAssemble(ED *egoclient.EgoDialogue, Req *egoclientReq.EgoDialoguePostUserMsg) (httpclient.Response, error) {
	if Req.ImageOption == nil {
		return nil, errors.New("错误的请求格式")
	}
	ctx := context.WithValue(context.Background(), "Dialogue", ED)
	imageReq := models.ImageRequest{
		Provider: ED.Model.ModelProvider,
		Model:    *ED.Model.ModelName,
		UserInfo: models.UserInfo{
			UserID: *ED.User.UserID,
		},
		Prompt: Req.Text,
		Size:   models.ImageSize(Req.ImageOption.Size),
	}
	imgResp, err := global.AiSDK.CreateImage(ctx, imageReq, httpclient.WithTimeout(time.Minute*5), httpclient.WithStreamReturnIntervalTimeout(time.Second*5))
	if err != nil {
		return nil, err
	}
	return &imgResp, nil
}

func (s *ArkService) DoubaoSeedReamHandler(ctx context.Context, ED *egoclient.EgoDialogue, resp httpclient.Response) (err error) {
	imageResp := resp.(*models.ImageResponse)

	//存储token
	if err = ModelSer.CreateEgoDialogueItem(ctx, &egoclient.EgoDialogueItem{
		UUID:             ED.UUID.String(), //图片这里用对话的UUID代替
		DialogueID:       ED.ID,
		CompletionTokens: imageResp.Usage.OutputTokens,
		PromptTokens:     0,
	}); err != nil {
		return
	}

	for _, v := range imageResp.Data {
		if err = ModelSer.CreateEgoDialogueHistory(ctx, &egoclient.EgoDialogueHistory{
			Role:       egoclient.AssistantRole,
			Item:       ED.UUID.String(), // 与token处一致
			DialogueID: ED.ID,
			Content:    v.URL, //未来按需更改为B64
			IsChoice:   true,
		}); err != nil {
			return
		}
	}
	return nil
}
