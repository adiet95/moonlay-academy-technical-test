package sublists

import (
	"encoding/json"
	"strconv"

	"github.com/adiet95/moonlay-academy-technical-test/src/database"
	"github.com/adiet95/moonlay-academy-technical-test/src/interfaces"
	"github.com/adiet95/moonlay-academy-technical-test/src/libs"
	"github.com/labstack/echo/v4"
)

type sublist_ctrl struct {
	svc interfaces.SublistService
}

func NewCtrl(reps interfaces.SublistService) *sublist_ctrl {
	return &sublist_ctrl{svc: reps}
}

func (l *sublist_ctrl) FindListId(e echo.Context) error {
	v := e.Request().URL.Query().Get("id")
	id, _ := strconv.Atoi(v)

	l.svc.FindListId(id).Send(e)
	return nil

}

func (l *sublist_ctrl) AddSub(e echo.Context) error {
	var data database.SubList
	err := json.NewDecoder(e.Request().Body).Decode(&data)
	if err != nil {
		libs.New(err.Error(), 400, true)
		return err
	}

	l.svc.AddSub(&data).Send(e)
	return nil
}

func (l *sublist_ctrl) UpdateSub(e echo.Context) error {
	v := e.Request().URL.Query().Get("id")
	id, _ := strconv.Atoi(v)

	var data database.SubList
	err := json.NewDecoder(e.Request().Body).Decode(&data)
	if err != nil {
		libs.New(err.Error(), 400, true)
		return err
	}

	l.svc.UpdateSub(&data, id).Send(e)
	return nil

}

func (l *sublist_ctrl) DeleteSub(e echo.Context) error {
	v := e.Request().URL.Query().Get("id")
	id, _ := strconv.Atoi(v)

	l.svc.DeleteSub(id).Send(e)
	return nil

}
