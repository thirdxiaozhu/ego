package egoModels

import (
	"context"
	"fmt"
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
			"any": s.DeepSeekAssemble,
		},
	}
}

func (s *DeepseekService) DeepSeekAssemble(ED *egoclient.EgoDialogue, Req *egoclientReq.EgoDialoguePostUserMsg) (httpclient.Response, error) {
	fmt.Println("DeepSeekChatAssemble", *ED.User.UserID, Req.DialogueID, Req.Text)

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

	return global.AiSDK.CreateChatCompletionStream(ctx, chatReq, httpclient.WithTimeout(time.Minute*2))
}
