package egoclient

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ EgoClientUserApi }

var ECUService = service.ServiceGroupApp.EgoclientServiceGroup.EgoClientUserService
