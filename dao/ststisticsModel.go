package dao

type StatisticsDaily struct {
	Id             string `json:"id"  db:"id"`
	DateCalculated string `json:"date_calculated" db:"date_calculated"`
	RegisterNum    int    `json:"register_num" db:"register_num"`
	LoginNum       int    `json:"login_num" db:"login_num"`
	VideoViewNum   int    `json:"video_view_num" db:"video_view_num"`
	CourseNum      int    `json:"course_num" db:"course_num"`
}
