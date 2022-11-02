package interfaces

import (
	"github.com/adiet95/moonlay-academy-technical-test/src/database"
	"github.com/adiet95/moonlay-academy-technical-test/src/libs"
)

type ListRepo interface {
	GetListWithSub(limit, offset int) (*database.Lists, error)
	GetListWOSub(limit, offset int) (*database.Lists, error)
	AddList(data *database.List) (*database.List, error)
	UpdateList(data *database.List, id int) (*database.List, error)
	DeleteList(id int) error
	FindId(id int) (*database.List, error)
	GetListAll() (*database.Lists, error)
	GetListAllWOSub() (*database.Lists, error)
}

type ListService interface {
	GetListWithSub(limit, offset int) *libs.Response
	GetListWOSub(limit, offset int) *libs.Response
	AddList(data *database.List) *libs.Response
	UpdateList(data *database.List, id int) *libs.Response
	DeleteList(id int) *libs.Response
	GetListAll() *libs.Response
	GetListAllWOSub() *libs.Response
}
