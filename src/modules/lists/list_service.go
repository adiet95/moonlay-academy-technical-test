package lists

import (
	"os"

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

func (l *list_service) GetListAll() *libs.Response {
	data, err := l.list_repo.GetListAll()
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (l *list_service) GetListAllWOSub() *libs.Response {
	data, err := l.list_repo.GetListAllWOSub()
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (l *list_service) AddList(data *database.List) *libs.Response {

	validate := libs.Validation(data.Title, data.Description)
	if validate != nil {
		return libs.New(validate.Error(), 400, true)
	}

	data, err := l.list_repo.AddList(data)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (l *list_service) UpdateList(data *database.List, id int) *libs.Response {
	oldData, err := l.list_repo.FindId(id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	oldPath := "./uploads/" + oldData.File
	if oldPath != "" {
		os.Remove(oldPath)
	}

	validate := libs.Validation(data.Title, data.Description)
	if validate != nil {
		return libs.New(validate.Error(), 400, true)
	}

	data, err1 := l.list_repo.UpdateList(data, id)
	if err1 != nil {
		return libs.New(err1.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (l *list_service) DeleteList(id int) *libs.Response {
	oldData, err := l.list_repo.FindId(id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	oldPath := "./uploads/" + oldData.File
	if oldPath != "" {
		os.Remove(oldPath)
	}

	data := l.list_repo.DeleteList(id)
	if data != nil {
		return libs.New(data.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}
