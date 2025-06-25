package egoModels

import (
	"context"
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

func (s *AliBLService) AliBLQwQPlusAssemble(ED *egoclient.EgoDialogue, Req *egoclientReq.EgoDialoguePostUserMsg) (httpclient.Response, error) {

	ctx := context.WithValue(context.Background(), "Dialogue", ED)
	chatReq := models.ChatRequest{
		Provider: ED.Model.ModelProvider,
		Model:    *ED.Model.ModelName,
		UserInfo: models.UserInfo{
			UserID: *ED.User.UserID,
		},
		Stream:              true,
		EnableThinking:      Req.Reasoning,
		MaxCompletionTokens: 4096,
		StreamOptions: &models.ChatStreamOptions{
			IncludeUsage: true,
		},
	}
	//插入历史消息
	for _, v := range ED.Histories {
		chatReq.Messages = append(chatReq.Messages, v.Role.GetMessage(v.Content, v.ReasoningContent))
	}

	return global.AiSDK.CreateChatCompletionStream(ctx, chatReq, httpclient.WithTimeout(time.Minute*5), httpclient.WithStreamReturnIntervalTimeout(time.Second*5))
}
