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

	rt.GET("/:pages", ctrl.GetListWOSub)
	rt.GET("/wsub/:pages", ctrl.GetListWithSub)
	rt.GET("/wsub", ctrl.GetAllWithSub)
	rt.GET("/", ctrl.GetAllWOSub)

	rt.DELETE("/:id", ctrl.DeleteList)

	list := rt.Group("/list")
	list.Use(middleware.UploadFile)
	list.POST("", ctrl.AddList)
	list.PUT("/:id", ctrl.UpdateList)
}
