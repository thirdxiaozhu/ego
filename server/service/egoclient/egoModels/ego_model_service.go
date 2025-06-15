package egoModels

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
	egoclientReq "github.com/flipped-aurora/gin-vue-admin/server/model/egoclient/request"
)

type Service interface {
	AssembleRequest(*egoclient.EgoDialogue, *egoclientReq.EgoDialoguePostUserMsg) (err error)
}

type AssembleFunc func(*egoclient.EgoDialogue, *egoclientReq.EgoDialoguePostUserMsg)

type BasicService struct {
	ModelAssemble map[string]map[string]AssembleFunc
}

func (s *BasicService) AssembleRequest(ED *egoclient.EgoDialogue, Req *egoclientReq.EgoDialoguePostUserMsg) (err error) {
	var modelType map[string]AssembleFunc
	var exists bool
	var fn AssembleFunc
	if modelType, exists = s.ModelAssemble[*ED.Model.ModelType]; !exists {
		return errors.New("model type not exists")
	}
	if fn, exists = modelType[*ED.Model.ModelName]; !exists || fn == nil {
		return errors.New("assemble Function not exists")
	}
	fn(ED, Req)
	return
}

var serviceRegistry = make(map[string]func() Service)

func RegisterService(name string, factory func() Service) {
	serviceRegistry[name] = factory
}

func GetService(name string) (Service, bool) {
	if factory, exists := serviceRegistry[name]; exists {
		return factory(), true
	}
	return nil, false
}

func AssembleRequest(ED *egoclient.EgoDialogue, Req *egoclientReq.EgoDialoguePostUserMsg) (err error) {
	service, ok := GetService(*ED.Model.ModelProvider)
	if !ok {
		err = errors.New("service not found")
	}
	err = service.AssembleRequest(ED, Req)
	if err != nil {
		return err
	}
	return
}
