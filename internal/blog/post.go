package blog

import (
	"strconv"
)

type ID int32

type Post struct {
	Id          ID     `json:"id"`
	ViewsNumber int32  `json:"views_number" sql:"default:0,notnull"`
	Name        string `json:"name"`
	ShortDescr  string `json:"short_descr"`
	Preview     string `json:"preview"`
	Content     string `json:"content"`
	URI         string `json:"uri"`
}

func (id ID) String() string {
	return strconv.Itoa(int(id))
}
