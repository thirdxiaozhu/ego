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
	AssembleRequest(*egoclient.EgoDialogue, *egoclientReq.EgoDialoguePostUserMsg) (httpclient.Response, error)
	ParseRequestModal(Req *egoclientReq.EgoDialoguePostUserMsg) (*models.UserMessage, error)
}

type AssembleFunc func(*egoclient.EgoDialogue, *egoclientReq.EgoDialoguePostUserMsg) (httpclient.Response, error)

type BasicService struct {
	ModelAssemble map[consts.ModelType]map[string]AssembleFunc
}

func (s *BasicService) AssembleRequest(ED *egoclient.EgoDialogue, Req *egoclientReq.EgoDialoguePostUserMsg) (httpclient.Response, error) {
	var modelType map[string]AssembleFunc
	var exists bool
	var fn AssembleFunc
	if modelType, exists = s.ModelAssemble[ED.Model.ModelType]; !exists {
		return nil, errors.New("model type not exists")
	}
	if fn, exists = modelType[*ED.Model.ModelName]; !exists || fn == nil {
		return nil, errors.New("assemble Function not exists")
	}
	return fn(ED, Req)
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

func AssembleRequest(ED *egoclient.EgoDialogue, Req *egoclientReq.EgoDialoguePostUserMsg) (httpclient.Response, error) {
	service, ok := GetService(ED.Model.ModelProvider)
	if !ok {
		return nil, errors.New("service not found")
	}
	return service.AssembleRequest(ED, Req)
}
