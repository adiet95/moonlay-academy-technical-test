package lists

import (
	"strconv"

	"github.com/adiet95/moonlay-academy-technical-test/src/database"
	"github.com/adiet95/moonlay-academy-technical-test/src/interfaces"
	"github.com/adiet95/moonlay-academy-technical-test/src/libs"
	"github.com/gorilla/schema"
	"github.com/labstack/echo/v4"
)

type list_ctrl struct {
	svc interfaces.ListService
}

func NewCtrl(reps interfaces.ListService) *list_ctrl {
	return &list_ctrl{svc: reps}
}

func (l *list_ctrl) GetListWithSub(e echo.Context) error {
	v := e.Request().URL.Query().Get("pages")
	pages, _ := strconv.Atoi(v)
	limit := 5
	var offset int

	if pages == 1 {
		offset = 0
	}
	offset = limit * (pages - 1)
	l.svc.GetListWithSub(limit, offset).Send(e)
	return nil
}

func (l *list_ctrl) GetListWOSub(e echo.Context) error {
	v := e.Request().URL.Query().Get("pages")
	pages, _ := strconv.Atoi(v)
	limit := 5
	var offset int

	if pages == 1 {
		offset = 0
	}
	offset = limit * (pages - 1)
	l.svc.GetListWOSub(limit, offset).Send(e)
	return nil

}

func (l *list_ctrl) FindListId(e echo.Context) error {
	v := e.Request().URL.Query().Get("id")
	id, _ := strconv.Atoi(v)

	l.svc.FindListId(id).Send(e)
	return nil

}

func (l *list_ctrl) AddList(e echo.Context) error {
	e.Logger().SetHeader("multipart/form-data")
	var decoder = schema.NewDecoder()
	var data database.List
	file := e.Get("image")

	//file upload
	files := file.(string)
	data.File = files

	err := decoder.Decode(&data, e.Request().PostForm)
	if err != nil {
		libs.New(err.Error(), 400, true)
		return err
	}
	l.svc.AddList(&data).Send(e)
	return nil

}

func (l *list_ctrl) AddSub(e echo.Context) error {
	e.Logger().SetHeader("multipart/form-data")
	var decoder = schema.NewDecoder()
	var data database.SubList
	file := e.Get("image")

	//file upload
	files := file.(string)
	data.SubFile = files

	err := decoder.Decode(&data, e.Request().PostForm)
	if err != nil {
		libs.New(err.Error(), 400, true)
		return err
	}
	l.svc.AddSub(&data).Send(e)
	return nil
}

func (l *list_ctrl) UpdateList(e echo.Context) error {
	e.Logger().SetHeader("multipart/form-data")
	v := e.Request().URL.Query().Get("id")
	id, _ := strconv.Atoi(v)

	var decoder = schema.NewDecoder()
	var data database.List
	file := e.Get("image")

	//file upload
	files := file.(string)
	data.File = files

	err := decoder.Decode(&data, e.Request().PostForm)
	if err != nil {
		libs.New(err.Error(), 400, true)
		return err
	}
	l.svc.UpdateList(&data, id).Send(e)
	return nil

}

func (l *list_ctrl) UpdateSub(e echo.Context) error {
	e.Logger().SetHeader("multipart/form-data")
	v := e.Request().URL.Query().Get("id")
	id, _ := strconv.Atoi(v)

	var decoder = schema.NewDecoder()
	var data database.SubList
	file := e.Get("image")

	//file upload
	files := file.(string)
	data.SubFile = files

	err := decoder.Decode(&data, e.Request().PostForm)
	if err != nil {
		libs.New(err.Error(), 400, true)
		return err
	}
	l.svc.UpdateSub(&data, id).Send(e)
	return nil

}

func (l *list_ctrl) DeleteList(e echo.Context) error {
	v := e.Request().URL.Query().Get("id")
	id, _ := strconv.Atoi(v)

	l.svc.DeleteList(id).Send(e)
	return nil

}

func (l *list_ctrl) DeleteSub(e echo.Context) error {
	v := e.Request().URL.Query().Get("id")
	id, _ := strconv.Atoi(v)

	l.svc.DeleteSub(id).Send(e)
	return nil

}
