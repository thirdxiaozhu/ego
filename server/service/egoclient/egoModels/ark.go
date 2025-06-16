package egoModels

import "github.com/liusuxian/go-aisdk/consts"

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
