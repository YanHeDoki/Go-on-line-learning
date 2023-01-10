package dao

type User struct {
	Id         string `json:"id" db:"id"`
	Openid     string `json:"openid" db:"openid"`
	Mobile     string `json:"mobile" db:"mobile"`
	Password   string `json:"password" db:"password"`
	Nickname   string `json:"nickname" db:"nickname"`
	Avatar     string `json:"avatar" db:"avatar"`
	Sign       string `json:"sign" db:"sign"`
	IsDisabled int    `json:"is_disabled" db:"is_disabled"`
	IsDeleted  int    `json:"is_deleted" db:"is_deleted"`
	Sex        int    `json:"sex" db:"sex"`
	Age        int    `json:"age" db:"age"`
}

type Order struct {
	Id          string  `json:"id" db:"id"`
	OrderNo     string  `json:"orderNo" db:"order_no"`
	CourseId    string  `json:"courseId" db:"course_id"`
	CourseTitle string  `json:"courseTitle" db:"course_title"`
	CourseCover string  `json:"courseCover" db:"course_cover"`
	TeacherName string  `json:"teacherName" db:"teacher_name"`
	MemberId    string  `json:"memberId" db:"member_id"`
	NickName    string  `json:"nickname" db:"nickname"`
	Mobile      string  `json:"mobile" db:"mobile"`
	TotalFee    float64 `json:"totalFee" db:"total_fee"`
	PayType     int8    `json:"payType" db:"pay_type"`
	Status      int8    `json:"status" db:"status"`
	IsDeleted   int8    `json:"isDeleted" db:"is_deleted"`
}

type OrderLog struct {
	Id            string  `json:"id" db:"id"`
	OrderNo       string  `json:"order_no" db:"order_no"`
	TradeState    string  ` json:"trade_state"  db:"trade_state"`
	Attr          string  `json:"attr" db:"attr"`
	TransactionId string  `json:"transaction_id" db:"transaction_id"`
	PayTime       string  `json:"pay_time" db:"pay_time"`
	TotalFee      float64 `json:"total_fee" db:"total_fee"`
	PayType       int8    `json:"pay_type" db:"pay_type"`
	Status        int8    `json:"status" db:"status"`
	IsDeleted     int8    `json:"is_deleted" db:"is_deleted"`
}
type test struct {
}

type AdminUser struct {
	Id       string `json:"id" db:"id"`
	Username string ` json:"username"  db:"username"`
	Password string `json:"password" db:"password"`
	Auth     string `json:"auth" db:"auth"`
}
