package sublists

import (
	"github.com/adiet95/moonlay-academy-technical-test/src/database"
	"github.com/adiet95/moonlay-academy-technical-test/src/interfaces"
	"github.com/adiet95/moonlay-academy-technical-test/src/libs"
)

type sublist_service struct {
	list_repo interfaces.SublistRepo
}

func NewService(reps interfaces.SublistRepo) *sublist_service {
	return &sublist_service{reps}
}

func (l *sublist_service) FindListId(id int) *libs.Response {
	data, err := l.list_repo.FindListId(id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (l *sublist_service) AddSub(data *database.SubList) *libs.Response {
	validate := libs.Validation(data.SubTitle, data.SubDescription)
	if validate != nil {
		return libs.New(validate.Error(), 400, true)
	}
	data, err := l.list_repo.AddSub(data)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (l *sublist_service) UpdateSub(data *database.SubList, id int) *libs.Response {
	validate := libs.Validation(data.SubTitle, data.SubDescription)
	if validate != nil {
		return libs.New(validate.Error(), 400, true)
	}
	data, err := l.list_repo.UpdateSub(data, id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (l *sublist_service) DeleteSub(id int) *libs.Response {
	data := l.list_repo.DeleteSub(id)
	if data != nil {
		return libs.New(data.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}
