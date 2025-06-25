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
	s.ModelAssemble = map[consts.ModelType]map[string]AssembleFunc{
		consts.ChatModel: {
			"qwq-plus": s.AliBLQwQPlusAssemble,
		},
	}
}

func (s *AliBLService) ParseRequestModal(Req *egoclientReq.EgoDialoguePostUserMsg) (*models.UserMessage, error) {
	userMsg := &models.UserMessage{
		Content: Req.Text,
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
			return nil, errors.New("多模态类型错误")
		}

		userMsg.MultimodalContent = append(userMsg.MultimodalContent, userMsgPart)
	}

	return userMsg, nil
}

func (s *AliBLService) AliBLQwQPlusAssemble(ED *egoclient.EgoDialogue, Req *egoclientReq.EgoDialoguePostUserMsg) (httpclient.Response, error) {

	ctx := context.WithValue(context.Background(), "Dialogue", ED)
	chatReq := models.ChatRequest{
		Provider: ED.Model.ModelProvider,
		Model:    *ED.Model.ModelName,
		UserInfo: models.UserInfo{
			UserID: *ED.User.UserID,
		},
		Stream:              true,
		EnableThinking:      Req.Reasoning, //Qwen3 默认开启thinking
		MaxCompletionTokens: 4096,
		StreamOptions: &models.ChatStreamOptions{
			IncludeUsage: true,
		},
	}

	//插入历史消息
	for _, v := range ED.Histories {
		chatReq.Messages = append(chatReq.Messages, v.Role.GetMessage(v.Content, v.ReasoningContent))
	}

	//插入用户当前消息
	var userMsg *models.UserMessage
	var err error
	if userMsg, err = s.ParseRequestModal(Req); err != nil {
		return nil, err
	}
	chatReq.Messages = append(chatReq.Messages, userMsg)

	return global.AiSDK.CreateChatCompletionStream(ctx, chatReq, httpclient.WithTimeout(time.Minute*5), httpclient.WithStreamReturnIntervalTimeout(time.Second*5))
}
