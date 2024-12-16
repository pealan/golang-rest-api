package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pealan/golang-rest-api/handler"
)

func DeviceRoutes(deviceRouter *gin.RouterGroup, deviceHandler *handler.DeviceHandler) {
	deviceRouter.GET("", deviceHandler.GetAllDevices)
	deviceRouter.POST("", deviceHandler.AddDevice)
	deviceRouter.GET("/:id", deviceHandler.GetDeviceById)
	deviceRouter.PUT("/:id", deviceHandler.UpdateDevice)
	deviceRouter.DELETE("/:id", deviceHandler.DeleteDevice)
}
