package sublists

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(rt *echo.Echo, db *gorm.DB) {
	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	rt.GET("/find", ctrl.FindListId)
	sub := rt.Group("/sub")
	sub.POST("", ctrl.AddSub)
	sub.PUT("", ctrl.UpdateSub)

	rt.DELETE("/sub", ctrl.DeleteSub)

}