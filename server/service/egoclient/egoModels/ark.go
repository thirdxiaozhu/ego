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
	ds := &ArkService{}
	ds.initAssemblers()
	return ds
}

func (s *ArkService) initAssemblers() {
	s.ModelAssemble = map[consts.ModelType]map[string]AssembleFunc{
		consts.ChatModel: {
			consts.Doubaoseed1_6: s.DoubaoSeedAssemble,
		},
	}
}

func (s *ArkService) ParseChatModal(ModelName string, Req *egoclientReq.EgoDialoguePostUserMsg) (*models.UserMessage, error) {
	userMsg := &models.UserMessage{}

	if len(Req.Multimodal) == 0 {
		userMsg.Content = Req.Text
	} else {
		userMsg.MultimodalContent = append(userMsg.MultimodalContent, models.ChatUserMsgPart{
			Type: models.ChatUserMsgPartTypeText,
			Text: Req.Text,
		})
	}

	for _, modal := range Req.Multimodal {
		userMsgPart := models.ChatUserMsgPart{
			Type: modal.Type,
		}

		switch modal.Type {
		case models.ChatUserMsgPartTypeText:
			userMsgPart.Text = modal.Text
		case models.ChatUserMsgPartTypeImageURL:
			userMsgPart.ImageURL = &models.ChatUserMsgImageURL{URL: modal.Url}
		default:
			return nil, errors.New("不支持的多模态类型")
		}

		userMsg.MultimodalContent = append(userMsg.MultimodalContent, userMsgPart)
	}

	return userMsg, nil
}

func (s *ArkService) DoubaoSeedAssemble(ED *egoclient.EgoDialogue, Req *egoclientReq.EgoDialoguePostUserMsg) (httpclient.Response, error) {
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

	if Req.Reasoning == true {
		chatReq.Thinking = &models.ChatThinkingOptions{
			Type: Req.ReasoningMode,
		}
	}

	//插入历史消息
	for _, v := range ED.Histories {
		chatReq.Messages = append(chatReq.Messages, v.Role.GetMessage(v.Content, v.ReasoningContent))
	}

	//插入用户当前消息
	var userMsg *models.UserMessage
	var err error
	if userMsg, err = s.ParseChatModal(*ED.Model.ModelName, Req); err != nil {
		return nil, err
	}
	chatReq.Messages = append(chatReq.Messages, userMsg)

	return global.AiSDK.CreateChatCompletionStream(ctx, chatReq, httpclient.WithTimeout(time.Minute*5), httpclient.WithStreamReturnIntervalTimeout(time.Second*5))
}
