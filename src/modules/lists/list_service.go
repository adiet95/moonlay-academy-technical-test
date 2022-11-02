package lists

import (
	"github.com/adiet95/moonlay-academy-technical-test/src/database"
	"github.com/adiet95/moonlay-academy-technical-test/src/interfaces"
	"github.com/adiet95/moonlay-academy-technical-test/src/libs"
)

type list_service struct {
	list_repo interfaces.ListRepo
}

func NewService(reps interfaces.ListRepo) *list_service {
	return &list_service{reps}
}

func (l *list_service) GetListWithSub(limit, offset int) *libs.Response {
	data, err := l.list_repo.GetListWithSub(limit, offset)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (l *list_service) GetListWOSub(limit, offset int) *libs.Response {
	data, err := l.list_repo.GetListWOSub(limit, offset)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (l *list_service) FindListId(id int) *libs.Response {
	data, err := l.list_repo.FindListId(id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (l *list_service) AddList(data *database.List) *libs.Response {
	data, err := l.list_repo.AddList(data)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (l *list_service) AddSub(data *database.SubList) *libs.Response {
	data, err := l.list_repo.AddSub(data)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (l *list_service) UpdateList(data *database.List, id int) *libs.Response {
	data, err := l.list_repo.UpdateList(data, id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (l *list_service) UpdateSub(data *database.SubList, id int) *libs.Response {
	data, err := l.list_repo.UpdateSub(data, id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (l *list_service) DeleteList(id int) *libs.Response {
	data := l.list_repo.DeleteList(id)
	if data != nil {
		return libs.New(data.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (l *list_service) DeleteSub(id int) *libs.Response {
	data := l.list_repo.DeleteSub(id)
	if data != nil {
		return libs.New(data.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}
