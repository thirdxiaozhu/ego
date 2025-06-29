package egoModels

import (
	"bytes"
	"context"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
	egoclientReq "github.com/flipped-aurora/gin-vue-admin/server/model/egoclient/request"
	"github.com/liusuxian/go-aisdk"
	"github.com/liusuxian/go-aisdk/consts"
	"github.com/liusuxian/go-aisdk/httpclient"
	"github.com/liusuxian/go-aisdk/models"
	"log"
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

type AssembleRequest func(*egoclient.EgoDialogue, *egoclientReq.EgoDialoguePostUserMsg) (httpclient.Response, error)

type HandleResponse func(ctx context.Context, ED *egoclient.EgoDialogue, resp httpclient.Response) error

type ModelHandler struct {
	AssembleRqstFunc AssembleRequest
	HandleRespFunc   HandleResponse
}

type Service interface {
	GetModelHandler(*egoclient.EgoDialogue) (*ModelHandler, error)
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

type ChatStreamContentBlock struct {
	ContentID         string
	SystemFingerprint string
	ContentBuffer     bytes.Buffer
	ReasoningBuffer   bytes.Buffer
}

func DefaultChatHandler(ctx context.Context, ED *egoclient.EgoDialogue, resp httpclient.Response) (err error) {
	streamResp := resp.(models.ChatResponseStream)

	var Contents []ChatStreamContentBlock
	var Item egoclient.EgoDialogueItem

	for {
		var (
			item       models.ChatBaseResponse
			isFinished bool
		)
		if item, isFinished, err = streamResp.StreamReader.Recv(); err != nil {
			log.Printf("createChatCompletionStream error = %v, request_id = %s", err, aisdk.RequestID(err))
			break
		}
		if isFinished {
			//把choices里的内容逐条插入history库里
			for _, v := range Contents {
				if err = ModelSer.CreateEgoDialogueHistory(ctx, &egoclient.EgoDialogueHistory{
					Role:             egoclient.AssistantRole,
					Item:             Item.UUID,
					DialogueID:       ED.ID,
					ReasoningContent: v.ReasoningBuffer.String(),
					Content:          v.ContentBuffer.String(),
					IsChoice:         true,
				}); err != nil {
					return
				}
			}
			break
		}
		log.Printf("createChatCompletionStream item = %+v", item)

		//TODO: 在这里做返回前端的SSE
		for _, v := range item.Choices {
			for v.Index >= len(Contents) {
				Contents = append(Contents, ChatStreamContentBlock{})
			}
			Contents[v.Index].ContentID = item.ID
			Contents[v.Index].ReasoningBuffer.WriteString(v.Delta.ReasoningContent)
			Contents[v.Index].ContentBuffer.WriteString(v.Delta.Content)
		}
		if item.Usage != nil && item.StreamStats != nil {
			//存储token用量
			Item.UUID = item.ID
			Item.PromptTokens = item.Usage.PromptTokens
			Item.DialogueID = ED.ID
			Item.CompletionTokens = item.Usage.CompletionTokens
			if err = ModelSer.CreateEgoDialogueItem(ctx, &Item); err != nil {
				return
			}
			log.Printf("createChatCompletionStream usage = %+v", item.Usage)
			log.Printf("createChatCompletionStream stream_stats = %+v", item.StreamStats)
		}
	}
	return nil
}
