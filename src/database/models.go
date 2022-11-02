package database

import "time"

type List struct {
	ListId      uint      `gorm:"primaryKey" json:"list_id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	File        string    `json:"file,omitempty"`
	SubList     []SubList `json:"sublist,omitempty"`
	CreatedAt   time.Time `gorm:"default:now(); not null" json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type SubList struct {
	SublistId      uint      `gorm:"primaryKey" json:"sublist_id,omitempty"`
	ListId         uint      `json:"list_id,omitempty"`
	List           List      `json:"list,omitempty"`
	SubTitle       string    `json:"sub_title,omitempty"`
	SubDescription string    `json:"sub_description,omitempty"`
	SubFile        string    `json:"sub_file,omitempty"`
	CreatedAt      time.Time `gorm:"default:now(); not null" json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}

type Lists []List
type Sublists []SubList
