package egoModels

import (
	"bytes"
	"context"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
	egoclientReq "github.com/flipped-aurora/gin-vue-admin/server/model/egoclient/request"
	"github.com/liusuxian/go-aisdk/consts"
	"github.com/liusuxian/go-aisdk/httpclient"
	"github.com/liusuxian/go-aisdk/models"
	"strings"
)

type ModelService struct {
}

var ModelSer ModelService

// CreateEgoDialogueHistory 创建Ego对话历史记录
// Author [yourname](https://github.com/yourname)
func (MS *ModelService) CreateEgoDialogueHistory(ctx context.Context, EDH *egoclient.EgoDialogueHistory) (err error) {
	err = global.GVA_DB.Create(EDH).Error
	return err
}

// CreateEgoDialogueItem 创建Ego对话历史记录
// Author [yourname](https://github.com/yourname)
func (MS *ModelService) CreateEgoDialogueItem(ctx context.Context, EDI *egoclient.EgoDialogueItem) (err error) {
	err = global.GVA_DB.Create(EDI).Error
	return err
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

func GetHandler(ED *egoclient.EgoDialogue) (*ModelHandler, error) {
	service, ok := GetService(ED.Model.ModelProvider)
	if !ok {
		return nil, errors.New("service not found")
	}
	return service.GetModelHandler(ED)
}

type AssembleRequest func(*egoclient.EgoDialogue, *egoclientReq.EgoDialoguePostRequest) (httpclient.Response, error)

type HandleResponse func(ctx context.Context, DialogueID uint) func(item models.ChatBaseResponse, isFinished bool) error

type ModelHandler struct {
	AssembleRqstFunc AssembleRequest
	HandleRespFunc   HandleResponse
}

type Service interface {
	GetModelHandler(*egoclient.EgoDialogue) (*ModelHandler, error)
	ParseChatRequest(ED *egoclient.EgoDialogue, req *egoclientReq.EgoDialoguePostRequest) (*models.ChatRequest, error)
	ParseChatModal(ModelName string, Text string, modals []egoclientReq.EgoDialogueMultiModal) (*models.UserMessage, error)
}

type BasicService struct {
	ModelHandlers map[consts.ModelType]map[string]*ModelHandler
}

func (s *BasicService) GetModelHandler(ED *egoclient.EgoDialogue) (*ModelHandler, error) {
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

func (s *BasicService) CheckModelValid(modelName string, toMatch ...string) bool {
	for _, match := range toMatch {
		if strings.Contains(modelName, match) {
			return true
		}
	}
	return false
}

type ChatStreamContentBlock struct {
	ContentID         string
	SystemFingerprint string
	ContentBuffer     bytes.Buffer
	ReasoningBuffer   bytes.Buffer
}
