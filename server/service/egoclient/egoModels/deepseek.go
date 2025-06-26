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
	s.ModelAssemble = map[consts.ModelType]map[string]AssembleFunc{
		consts.ChatModel: {
			"any": s.DeepSeekReasonerAssemble,
		},
	}
}

func (s *DeepseekService) ParseChatModal(ModelName string, Req *egoclientReq.EgoDialoguePostUserMsg) (*models.UserMessage, error) {
	userMsg := &models.UserMessage{
		Content: Req.Text,
	}
	if len(Req.Multimodal) != 0 {
		return nil, errors.New("deepseek 不支持多模态输入")
	}

	return userMsg, nil
}

func (s *DeepseekService) DeepSeekReasonerAssemble(ED *egoclient.EgoDialogue, Req *egoclientReq.EgoDialoguePostUserMsg) (httpclient.Response, error) {
	model := consts.DeepSeekChat
	if Req.Reasoning {
		model = consts.DeepSeekReasoner
	}

	ctx := context.WithValue(context.Background(), "Dialogue", ED)
	chatReq := models.ChatRequest{
		Provider: consts.DeepSeek,
		Model:    model,
		UserInfo: models.UserInfo{
			UserID: *ED.User.UserID,
		},
		Stream:              true,
		MaxCompletionTokens: 4096,
	}
	//插入历史消息
	for _, v := range ED.Histories {
		chatReq.Messages = append(chatReq.Messages, v.Role.GetMessage(v.Content, v.ReasoningContent))
	}

	//插入用户当前消息
	var userMsg *models.UserMessage
	var err error
	if userMsg, err = s.ParseChatModal("", Req); err != nil {
		return nil, err
	}
	chatReq.Messages = append(chatReq.Messages, userMsg)

	return global.AiSDK.CreateChatCompletionStream(ctx, chatReq, httpclient.WithTimeout(time.Minute*5), httpclient.WithStreamReturnIntervalTimeout(time.Second*5))
}
