package lists

import (
	"github.com/adiet95/moonlay-academy-technical-test/src/middleware"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(rt *echo.Echo, db *gorm.DB) {
	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	rt.GET("/", ctrl.GetListWithSub)
	rt.GET("/sub", ctrl.GetListWOSub)
	rt.GET("/find", ctrl.FindListId)

	rt.POST("/", ctrl.AddList, middleware.UploadFile)
	rt.POST("/sub", ctrl.AddSub, middleware.UploadFile)

	rt.PUT("/", ctrl.UpdateList, middleware.UploadFile)
	rt.PUT("/sub", ctrl.UpdateSub, middleware.UploadFile)

	rt.DELETE("/", ctrl.DeleteList)
	rt.DELETE("/sub", ctrl.DeleteSub)
}
