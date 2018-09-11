package blog

import "strconv"

type ID int64
type Counter uint32

type Post struct {
	Id          ID      `json:"id"`
	ViewsNumber Counter `json:"views_number"`
	Name        string  `json:"name"`
	ShortDescr  string  `json:"short_descr"`
	Preview     string  `json:"preview"`
	Content     string  `json:"content"`
	URI         string  `json:"uri"`
}

func (id ID) String() string {
	return strconv.Itoa(int(id))
}
