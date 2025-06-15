package egoModels

type ArkService struct {
	BasicService
}

func init() {
	RegisterService("ark", func() Service {
		return &ArkService{
			BasicService: BasicService{
				ModelAssemble: map[string]map[string]AssembleFunc{
					"chat": {
						"deepseek-chat":     nil,
						"deepseek-reasoner": nil,
					},
				},
			},
		}
	})
}
