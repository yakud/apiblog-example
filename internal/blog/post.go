package blog

type ID int64
type Counter uint32

type Post struct {
	Id          ID
	ViewsNumber Counter

	Name       string
	ShortDescr string
	Preview    string
	Content    string
	URI        string
}
