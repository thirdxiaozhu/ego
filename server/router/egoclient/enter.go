package egoclient

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ EgoClientUserRouter }

var ECUApi = api.ApiGroupApp.EgoclientApiGroup.EgoClientUserApi
