package egoModels

import (
	egoclientReq "github.com/flipped-aurora/gin-vue-admin/server/model/egoclient/request"
	"github.com/liusuxian/go-aisdk/consts"
	"github.com/liusuxian/go-aisdk/models"
)

type ArkService struct {
	BasicService
}

func init() {
	RegisterService("ark", func() Service {
		return &ArkService{
			BasicService: BasicService{
				ModelAssemble: map[consts.ModelType]map[string]AssembleFunc{
					consts.ChatModel: {
						"deepseek-chat":     nil,
						"deepseek-reasoner": nil,
					},
				},
			},
		}
	})
}

func (s *ArkService) ParseChatModal(ModelName string, Req *egoclientReq.EgoDialoguePostUserMsg) (*models.UserMessage, error) {

	return nil, nil
}
