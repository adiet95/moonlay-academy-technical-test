package lists

import (
	"errors"

	"github.com/adiet95/moonlay-academy-technical-test/src/database"
	"gorm.io/gorm"
)

type list_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *list_repo {
	return &list_repo{db}
}

func (l *list_repo) GetListWithSub(limit, offset int) (*database.Lists, error) {
	var datas database.Lists

	result := l.db.Preload("Sub").Limit(limit).Offset(offset).Find(&datas)

	if result.RowsAffected == 0 {
		return nil, errors.New("list datas is empty")
	}
	if result.Error != nil {
		return nil, errors.New("failed obtain datas list")
	}
	return &datas, nil
}

func (l *list_repo) GetListWOSub(limit, offset int) (*database.Lists, error) {
	var datas database.Lists

	result := l.db.Limit(limit).Offset(offset).Find(&datas)

	if result.RowsAffected == 0 {
		return nil, errors.New("list datas is empty")
	}
	if result.Error != nil {
		return nil, errors.New("failed obtain datas list")
	}
	return &datas, nil
}

func (l *list_repo) FindListId(id int) (*database.List, error) {
	var datas database.List

	result := l.db.Preload("Sublist").Find(&datas)

	if result.RowsAffected == 0 {
		return nil, errors.New("list datas is empty")
	}
	if result.Error != nil {
		return nil, errors.New("failed obtain datas list")
	}
	return &datas, nil
}

func (l *list_repo) AddList(data *database.List) (*database.List, error) {

	res := l.db.Create(data)
	if res.Error != nil {
		return nil, errors.New("failed create data list")
	}

	return data, nil
}

func (l *list_repo) AddSub(data *database.SubList) (*database.SubList, error) {
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

func (l *list_repo) UpdateList(data *database.List, id int) (*database.List, error) {
	var datas database.List

	result := l.db.Where("list_id = ?", id).Find(&datas)

	if result.RowsAffected == 0 {
		return nil, errors.New("list datas is empty")
	}
	if result.Error != nil {
		return nil, errors.New("failed obtain datas list")
	}

	res := l.db.Where("list_id = ?", id).Updates(data)
	if res.Error != nil {
		return nil, errors.New("failed obtain datas list")
	}

	return data, nil
}

func (l *list_repo) UpdateSub(data *database.SubList, id int) (*database.SubList, error) {
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

func (l *list_repo) DeleteList(id int) error {
	var datas database.Lists
	var data database.List

	result := l.db.Where("list_id = ?", id).Find(&datas)

	if result.RowsAffected == 0 {
		return errors.New("list datas is empty")
	}
	if result.Error != nil {
		return errors.New("failed obtain datas list")
	}

	res := l.db.Where("list_id = ?", id).Delete(&data)
	if res.Error != nil {
		return errors.New("failed obtain datas list")
	}

	return nil
}

func (l *list_repo) DeleteSub(id int) error {
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
