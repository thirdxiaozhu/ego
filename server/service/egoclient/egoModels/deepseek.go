package egoModels

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
	egoclientReq "github.com/flipped-aurora/gin-vue-admin/server/model/egoclient/request"
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
	s.ModelAssemble = map[string]map[string]AssembleFunc{
		"chat": {
			"any": s.DeepSeekAssemble,
		},
	}
}

func (s *DeepseekService) DeepSeekAssemble(ED *egoclient.EgoDialogue, Req *egoclientReq.EgoDialoguePostUserMsg) {
	fmt.Println("DeepSeekChatAssemble", *ED.User.UserID, Req.DialogueID, Req.Text)
	//ctx := context.Background()
	//_, err := global.AiSDK.CreateChatCompletionStream(ctx, *ED.User.UserID, models.ChatRequest{}, nil, httpclient.WithTimeout(time.Minute*2))
	//if err != nil {
	//	return
	//}
}
