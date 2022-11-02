package interfaces

import (
	"github.com/adiet95/moonlay-academy-technical-test/src/database"
	"github.com/adiet95/moonlay-academy-technical-test/src/libs"
)

type ListRepo interface {
	GetListWithSub(limit, offset int) (*database.Lists, error)
	GetListWOSub(limit, offset int) (*database.Lists, error)
	FindListId(id int) (*database.List, error)
	AddList(data *database.List) (*database.List, error)
	AddSub(data *database.SubList) (*database.SubList, error)
	UpdateList(data *database.List, id int) (*database.List, error)
	UpdateSub(data *database.SubList, id int) (*database.SubList, error)
	DeleteList(id int) error
	DeleteSub(id int) error
}

type ListService interface {
	GetListWithSub(limit, offset int) *libs.Response
	GetListWOSub(limit, offset int) *libs.Response
	FindListId(id int) *libs.Response
	AddList(data *database.List) *libs.Response
	AddSub(data *database.SubList) *libs.Response
	UpdateList(data *database.List, id int) *libs.Response
	UpdateSub(data *database.SubList, id int) *libs.Response
	DeleteList(id int) *libs.Response
	DeleteSub(id int) *libs.Response
}
