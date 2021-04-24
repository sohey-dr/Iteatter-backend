package model

type Post struct {
	Id       int `uri:"id" binding:"requiered,uuid"`
	Title    string
	Body     string
	UserFrom int
	// CreatedAt ???
	// UpdatedAt ???
}
