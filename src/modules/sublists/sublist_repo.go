package sublists

import (
	"errors"

	"github.com/adiet95/moonlay-academy-technical-test/src/database"
	"gorm.io/gorm"
)

type sublist_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *sublist_repo {
	return &sublist_repo{db}
}

func (l *sublist_repo) FindListId(id int) (*database.Sublists, error) {
	var datas database.Sublists

	result := l.db.Where("list_id = ?", id).Find(&datas)

	if result.RowsAffected == 0 {
		return nil, errors.New("list datas is empty")
	}
	if result.Error != nil {
		return nil, errors.New("failed obtain datas list")
	}
	return &datas, nil
}

func (l *sublist_repo) AddSub(data *database.SubList) (*database.SubList, error) {
	var list database.Lists
	result := l.db.Where("list_id = ?", data.ListId).Find(&list)

	if result.RowsAffected == 0 {
		return nil, errors.New("list data not found")
	}
	if result.Error != nil {
		return nil, errors.New("failed obtain datas list")
	}

	res := l.db.Create(data)
	if res.Error != nil {
		return nil, errors.New("failed to create data sublist")
	}

	return data, nil
}

func (l *sublist_repo) UpdateSub(data *database.SubList, id int) (*database.SubList, error) {
	var datas database.Sublists

	result := l.db.Where("sublist_id = ?", id).Find(&datas)

	if result.RowsAffected == 0 {
		return nil, errors.New("list datas is empty")
	}
	if result.Error != nil {
		return nil, errors.New("failed obtain datas list")
	}

	res := l.db.Where("sublist_id = ?", id).Updates(data)
	if res.Error != nil {
		return nil, errors.New("failed obtain datas list")
	}

	return data, nil
}

func (l *sublist_repo) DeleteSub(id int) error {
	var datas database.Sublists
	var data database.SubList

	result := l.db.Where("sublist_id = ?", id).Find(&datas)

	if result.RowsAffected == 0 {
		return errors.New("list datas is empty")
	}
	if result.Error != nil {
		return errors.New("failed obtain datas list")
	}

	res := l.db.Where("sublist_id = ?", id).Delete(&data)
	if res.Error != nil {
		return errors.New("failed obtain datas list")
	}
	return nil
}
