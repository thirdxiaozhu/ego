package egoclient

import (
	"context"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
	egoclientReq "github.com/flipped-aurora/gin-vue-admin/server/model/egoclient/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service/egoclient/egoModels"
	"github.com/google/uuid"
	"github.com/liusuxian/go-aisdk"
	"github.com/liusuxian/go-aisdk/httpclient"
	"github.com/liusuxian/go-aisdk/models"
	"log"
)

type EgoDialogueService struct{}

// CreateEgoDialogue 创建Ego对话记录
// Author [yourname](https://github.com/yourname)
func (EDService *EgoDialogueService) CreateEgoDialogue(ctx context.Context, userid uint, ED *egoclient.EgoDialogue) (err error) {
	ED.UserID = userid
	ED.UUID, _ = uuid.NewV6()
	err = global.GVA_DB.Create(ED).Error
	return err
}

// DeleteEgoDialogue 删除Ego对话记录
// Author [yourname](https://github.com/yourname)
func (EDService *EgoDialogueService) DeleteEgoDialogue(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&egoclient.EgoDialogue{}, "id = ?", ID).Error
	return err
}

// DeleteEgoDialogueByIds 批量删除Ego对话记录
// Author [yourname](https://github.com/yourname)
func (EDService *EgoDialogueService) DeleteEgoDialogueByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]egoclient.EgoDialogue{}, "id in ?", IDs).Error
	return err
}

// UpdateEgoDialogue 更新Ego对话记录
// Author [yourname](https://github.com/yourname)
func (EDService *EgoDialogueService) UpdateEgoDialogue(ctx context.Context, ED egoclient.EgoDialogue) (err error) {
	fmt.Println(ED)
	err = global.GVA_DB.Model(&egoclient.EgoDialogue{}).Where("id = ?", ED.ID).Updates(&ED).Error
	return err
}

// GetEgoDialogue 根据ID获取Ego对话记录
// Author [yourname](https://github.com/yourname)
func (EDService *EgoDialogueService) GetEgoDialogue(ctx context.Context, ID string) (ED egoclient.EgoDialogue, err error) {
	err = global.GVA_DB.Where("id = ?", ID).Preload("Model").Preload("User").Preload("Histories").Preload("Items").First(&ED).Error
	return
}

// GetEgoDialogueByUuid 根据ID获取Ego对话记录
// Author [yourname](https://github.com/yourname)
func (EDService *EgoDialogueService) GetEgoDialogueByUuid(ctx context.Context, Uuid string) (ED egoclient.EgoDialogue, err error) {
	err = global.GVA_DB.Where("uuid = ?", Uuid).Preload("Model").Preload("User").Preload("Histories").Preload("Items").First(&ED).Error
	return
}

// GetEgoDialogueInfoList 分页获取Ego对话记录
// Author [yourname](https://github.com/yourname)
func (EDService *EgoDialogueService) GetEgoDialogueInfoList(ctx context.Context, info egoclientReq.EgoDialogueSearch) (list []egoclient.EgoDialogue, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&egoclient.EgoDialogue{})
	var EDs []egoclient.EgoDialogue
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.UUID != nil && *info.UUID != "" {
		db = db.Where("uuid = ?", *info.UUID)
	}
	if info.User != nil {
		db = db.Where("user = ?", *info.User)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Preload("Model").Preload("User")
	}

	err = db.Find(&EDs).Error
	return EDs, total, err
}
func (EDService *EgoDialogueService) GetEgoDialoguePublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// PostEgoDialogueUserMsg 创建Ego对话记录
// Author [yourname](https://github.com/yourname)
func (EDService *EgoDialogueService) PostEgoDialogueUserMsg(ctx context.Context, Req *egoclientReq.EgoDialoguePostUserMsg) error {
	ED, err := EDService.GetEgoDialogueByUuid(ctx, Req.DialogueID)
	if err != nil {
		return errors.New("无法找到对话")
	}

	//新的用户消息存入历史数据中
	newHistory := egoclient.EgoDialogueHistory{
		Role:             egoclient.UserRole,
		ItemID:           "",
		IsChoice:         true,
		ConversationID:   ED.ID,
		ReasoningContent: "",
		Content:          Req.Text,
	}
	if err = EDService.CreateEgoDialogueHistory(ctx, &newHistory); err != nil {
		return err
	}

	ED.Histories = append(ED.Histories, newHistory)

	//发送请求
	var resp httpclient.Response
	if resp, err = egoModels.AssembleRequest(&ED, Req); err != nil {
		return err
	}
	streamResp := resp.(models.ChatResponseStream)

	go func() {
		var Contents []models.ChatStreamContentBlock
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
					history := egoclient.EgoDialogueHistory{
						Role:             egoclient.AssistantRole,
						ItemID:           item.ID,
						ConversationID:   ED.ID,
						ReasoningContent: v.ReasoningBuffer.String(),
						Content:          v.ContentBuffer.String(),
						IsChoice:         true,
					}

					if err = EDService.CreateEgoDialogueHistory(ctx, &history); err != nil {
						return
					}
				}
				break
			}
			log.Printf("createChatCompletionStream item = %+v", item)

			//TODO: 在这里做返回前端的SSE
			for _, v := range item.Choices {
				for v.Index >= len(Contents) {
					Contents = append(Contents, models.ChatStreamContentBlock{})
				}
				Contents[v.Index].ContentID = item.ID
				Contents[v.Index].ContentID = item.ID
				Contents[v.Index].ReasoningBuffer.WriteString(v.Delta.ReasoningContent)
				Contents[v.Index].ContentBuffer.WriteString(v.Delta.Content)
			}
			if item.Usage != nil && item.StreamStats != nil {
				//存储token用量
				Item.UUID = item.ID
				Item.PromptTokens = item.Usage.PromptTokens
				Item.ConversationID = ED.ID
				Item.CompletionTokens = item.Usage.CompletionTokens
				if err = EDService.CreateEgoDialogueItem(ctx, &Item); err != nil {
					return
				}
				log.Printf("createChatCompletionStream usage = %+v", item.Usage)
				log.Printf("createChatCompletionStream stream_stats = %+v", item.StreamStats)
			}
		}
	}()

	return nil
}

// CreateEgoDialogueHistory 创建Ego对话历史记录
// Author [yourname](https://github.com/yourname)
func (EDService *EgoDialogueService) CreateEgoDialogueHistory(ctx context.Context, EDH *egoclient.EgoDialogueHistory) (err error) {
	err = global.GVA_DB.Create(EDH).Error
	return err
}

// CreateEgoDialogueItem 创建Ego对话历史记录
// Author [yourname](https://github.com/yourname)
func (EDService *EgoDialogueService) CreateEgoDialogueItem(ctx context.Context, EDI *egoclient.EgoDialogueItem) (err error) {
	err = global.GVA_DB.Create(EDI).Error
	return err
}
