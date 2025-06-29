package egoModels

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
	egoclientReq "github.com/flipped-aurora/gin-vue-admin/server/model/egoclient/request"
	"github.com/liusuxian/go-aisdk/consts"
	"github.com/liusuxian/go-aisdk/httpclient"
	"github.com/liusuxian/go-aisdk/models"
)

type Service interface {
	GetModelHandler(*egoclient.EgoDialogue, *egoclientReq.EgoDialoguePostUserMsg) (*ModelHandler, error)
	ParseChatModal(ModelName string, Text string, modals []egoclientReq.EgoDialogueMultiModal) (*models.UserMessage, error)
}

type AssembleFunc func(*egoclient.EgoDialogue, *egoclientReq.EgoDialoguePostUserMsg) (httpclient.Response, error)
type HandleFunc func(*httpclient.Response) error

type ModelHandler struct {
	AssembleFunc AssembleFunc
	HandleFunc   HandleFunc
}

type BasicService struct {
	ModelHandlers map[consts.ModelType]map[string]*ModelHandler
}

func (s *BasicService) GetModelHandler(ED *egoclient.EgoDialogue, Req *egoclientReq.EgoDialoguePostUserMsg) (*ModelHandler, error) {
	var modelType map[string]*ModelHandler
	var exists bool
	var handler *ModelHandler
	if modelType, exists = s.ModelHandlers[ED.Model.ModelType]; !exists {
		return nil, errors.New("model type not exists")
	}
	if handler, exists = modelType[*ED.Model.ModelName]; !exists || handler == nil {
		return nil, errors.New("assemble Function not exists")
	}
	return handler, nil
}

var serviceRegistry = make(map[consts.Provider]func() Service)

func RegisterService(name consts.Provider, factory func() Service) {
	serviceRegistry[name] = factory
}

func GetService(name consts.Provider) (Service, bool) {
	if factory, exists := serviceRegistry[name]; exists {
		return factory(), true
	}
	return nil, false
}

func GetHandler(ED *egoclient.EgoDialogue, Req *egoclientReq.EgoDialoguePostUserMsg) (*ModelHandler, error) {
	service, ok := GetService(ED.Model.ModelProvider)
	if !ok {
		return nil, errors.New("service not found")
	}
	return service.GetModelHandler(ED, Req)
}
