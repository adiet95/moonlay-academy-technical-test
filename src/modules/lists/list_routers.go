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

	rt.GET("/", ctrl.GetListWOSub)
	rt.GET("/sub", ctrl.GetListWithSub)
	rt.GET("/find", ctrl.FindListId)

	list := rt.Group("/list")
	list.Use(middleware.UploadFile)
	list.POST("", ctrl.AddList)
	list.PUT("", ctrl.UpdateList, middleware.UploadFile)

	sub := rt.Group("/sub")
	sub.POST("", ctrl.AddSub)
	sub.PUT("", ctrl.UpdateSub)

	rt.DELETE("/", ctrl.DeleteList)
	rt.DELETE("/sub", ctrl.DeleteSub)
}
