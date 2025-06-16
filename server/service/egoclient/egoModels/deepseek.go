package egoModels

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
	egoclientReq "github.com/flipped-aurora/gin-vue-admin/server/model/egoclient/request"
	"github.com/liusuxian/go-aisdk/consts"
	"github.com/liusuxian/go-aisdk/httpclient"
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

func (s *DeepseekService) DeepSeekAssemble(ED *egoclient.EgoDialogue, Req *egoclientReq.EgoDialoguePostUserMsg) {
	fmt.Println("DeepSeekChatAssemble", *ED.User.UserID, Req.DialogueID, Req.Text)

	model := consts.DeepSeekChat
	if Req.Reasoning {
		model = consts.DeepSeekReasoner
	}

	ctx := context.WithValue(context.Background(), "Dialogue", ED)
	_, err := global.AiSDK.CreateChatCompletionStream(ctx, models.ChatRequest{
		ModelInfo: models.ModelInfo{
			Provider:  consts.DeepSeek,
			ModelType: consts.ChatModel,
			Model:     model,
		},
		UserInfo: models.UserInfo{
			UserID: *ED.User.UserID,
		},
		Messages: []models.ChatMessage{
			&models.UserMessage{
				Content: Req.Text,
			},
		},
		Stream:              true,
		MaxCompletionTokens: 4096,
	}, streamCallback, httpclient.WithTimeout(time.Minute*2))
	if err != nil {
		return
	}
}
func streamCallback(ctx context.Context, response models.ChatResponse) error {
	//if response.Choices[0].Delta.Content == "" {
	//	fmt.Print(response.Choices[0].Delta.ReasoningContent)
	//} else {
	//	fmt.Print(response.Choices[0].Delta.Content)
	//}
	dialogue := ctx.Value("Dialogue").(*egoclient.EgoDialogue)
	fmt.Println(dialogue.UUID)
	return nil
}
