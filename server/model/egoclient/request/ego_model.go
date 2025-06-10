
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type EgoModelSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      ModelProvider  *string `json:"modelProvider" form:"modelProvider"` 
      ModelType  *string `json:"modelType" form:"modelType"` 
      ModelName  *string `json:"modelName" form:"modelName"` 
    request.PageInfo
}
