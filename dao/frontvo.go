package dao

type CourseFrontVo struct {
	Title           string `json:"title" db:"title"`
	TeacherId       string `json:"teacher_id" db:"teacher_id"`
	SubjectParentId string `json:"subjectParentId" db:"subject_parent_id"`
	SubjectId       string `json:"subjectId" db:"subject_id"`
	BuyCountSort    string `json:"buyCountSort" db:"buy_count_sort"`
	GmtCreateSort   string `json:"gmtCreateSort" db:"gmt_create_sort"`
	PriceSort       string `json:"priceSort" db:"price_sort"`
}
