package dao

type Banner struct {
	Id       string `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	ImageUrl string `json:"imageUrl" db:"image_url"`
	LinkUrl  string `json:"linkUrl" db:"link_url"`
	Sort     int    `json:"sort" db:"sort"`
}
