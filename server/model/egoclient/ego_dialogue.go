// 自动生成模板EgoDialogue
package egoclient

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/google/uuid"
	"github.com/liusuxian/go-aisdk/models"
)

type RoleType string

const (
	SystemRole    RoleType = "system"
	UserRole      RoleType = "user"
	AssistantRole RoleType = "assistant"
	ToolRole      RoleType = "tool"
)

func (role RoleType) GetMessage(content, reasonContent string) models.ChatMessage {
	switch role {
	case SystemRole:
		return &models.SystemMessage{
			Content: content,
		}
	case UserRole:
		return &models.UserMessage{
			Content: content,
		}
	case AssistantRole:
		return &models.AssistantMessage{
			Content: content,
			//Prefix:           true,
			//ReasoningContent: reasonContent,
		}
	case ToolRole:
		return &models.ToolMessage{
			Content: content,
		}
	default:
		return nil
	}
}

// Ego对话 结构体  EgoDialogue
type EgoDialogue struct {
	global.GVA_MODEL
	UUID      uuid.UUID            `json:"uuid" form:"uuid" gorm:"column:uuid;"`   //对话UUID
	UserID    uint                 `json:"userID" form:"user" gorm:"column:user;"` //所属用户
	User      EgoClientUser        `json:"user" gorm:"foreignKey:ID;references:UserID;"`
	ModelID   int                  `json:"modelID" form:"model" gorm:"column:model;"` //模型
	Model     EgoModel             `json:"model" gorm:"foreignKey:ID;references:ModelID;"`
	Items     []EgoDialogueItem    `json:"items" gorm:"foreignKey:ConversationID"`      //token使用情况
	Histories []EgoDialogueHistory `json:"histories" gorm:"foreignKey:ConversationID;"` //历史记录
}

// TableName Ego对话 EgoDialogue自定义表名 ego_dialogue
func (EgoDialogue) TableName() string {
	return "ego_dialogue"
}

type EgoDialogueItem struct {
	global.GVA_MODEL
	UUID             string `json:"uuid" form:"uuid" gorm:"column:uuid;"`
	ConversationID   uint   `json:"conversation_id" gorm:"conversation-id;comment:关联对话ID"`
	CompletionTokens int    `json:"completion_tokens" form:"completion_tokens" gorm:"column:completion_tokens;"`
	PromptTokens     int    `json:"prompt_tokens" form:"prompt_tokens" gorm:"column:prompt_tokens;"`
}

func (EgoDialogueItem) TableName() string {
	return "ego_dialogue_item"
}

type EgoDialogueHistory struct {
	global.GVA_MODEL
	Role             RoleType `json:"role" form:"role" gorm:"column:role;"`
	Item             string   `json:"item" form:"item" gorm:"column:item;"`                  //所属Item UUID
	ConversationID   uint     `json:"conversation_id" gorm:"conversation-id;comment:关联对话ID"` //所属用户的ID（自增主键）
	ReasoningContent string   `json:"reasoning_content" form:"reasoning-content" gorm:"type:text;column:reasoning-content;"`
	Content          string   `json:"content" form:"content" gorm:"type:text;column:content;"`
	IsChoice         bool     `json:"isChoice" form:"isChoice" gorm:"column:is_choice;"`
}

func (EgoDialogueHistory) TableName() string {
	return "ego_dialogue_history"
}
