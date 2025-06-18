package egoModels

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
	egoclientReq "github.com/flipped-aurora/gin-vue-admin/server/model/egoclient/request"
	"github.com/liusuxian/go-aisdk"
	"github.com/liusuxian/go-aisdk/consts"
	"github.com/liusuxian/go-aisdk/httpclient"
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
	s.ModelAssemble = map[consts.ModelType]map[string]AssembleFunc{
		consts.ChatModel: {
			"any": s.DeepSeekAssemble,
		},
	}
}

func (s *DeepseekService) DeepSeekAssemble(ED *egoclient.EgoDialogue, Req *egoclientReq.EgoDialoguePostUserMsg) (histories []*egoclient.EgoDialogueHistory, err error) {
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
		Messages: []models.ChatMessage{
			&models.UserMessage{
				Content: Req.Text,
			},
		},
		Stream:              true,
		MaxCompletionTokens: 4096,
	}

	resp, err := global.AiSDK.CreateChatCompletionStream(ctx, chatReq, httpclient.WithTimeout(time.Minute*2))
	if err != nil {
		return nil, err
	}
	var Contents []models.ChatStreamContentBlock

	for {
		var (
			item       models.ChatBaseResponse
			isFinished bool
		)
		if item, isFinished, err = resp.StreamReader.Recv(); err != nil {
			log.Printf("createChatCompletionStream error = %v, request_id = %s", err, aisdk.RequestID(err))
			break
		}
		if isFinished {
			for _, v := range Contents {
				history := egoclient.EgoDialogueHistory{
					ConversationID:   ED.ID,
					Role:             egoclient.AssistantRole,
					ReasoningContent: v.ReasoningBuffer.String(),
					Content:          v.ContentBuffer.String(),
				}
				histories = append(histories, &history)
			}
			break
		}
		log.Printf("createChatCompletionStream item = %+v", item)
		for _, v := range item.Choices {
			for v.Index >= len(Contents) {
				Contents = append(Contents, models.ChatStreamContentBlock{})
			}
			Contents[v.Index].ReasoningBuffer.WriteString(v.Delta.ReasoningContent)
			Contents[v.Index].ContentBuffer.WriteString(v.Delta.Content)
		}
		if item.Usage != nil && item.StreamStats != nil {
			log.Printf("createChatCompletionStream usage = %+v", item.Usage)
			log.Printf("createChatCompletionStream stream_stats = %+v", item.StreamStats)
		}
	}

	log.Println(*histories[0])
	return
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
