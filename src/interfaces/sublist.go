package interfaces

import (
	"github.com/adiet95/moonlay-academy-technical-test/src/database"
	"github.com/adiet95/moonlay-academy-technical-test/src/libs"
)

type SublistRepo interface {
	FindListId(id int) (*database.Sublists, error)
	AddSub(data *database.SubList) (*database.SubList, error)
	UpdateSub(data *database.SubList, id int) (*database.SubList, error)
	DeleteSub(id int) error
}

type SublistService interface {
	FindListId(id int) *libs.Response
	AddSub(data *database.SubList) *libs.Response
	UpdateSub(data *database.SubList, id int) *libs.Response
	DeleteSub(id int) *libs.Response
}
