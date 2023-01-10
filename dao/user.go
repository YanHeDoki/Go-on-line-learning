package dao

import (
	"database/sql"
	"goadmin/dao/mysql"
	"goadmin/utils"
	"time"
)

func GetUserById(Id string) (*User, error) {
	db := mysql.GetDb()
	sqlstr := "select id,openid,mobile,password,nickname,sex,age,avatar,sign,is_disabled from ucenter_member where id=? "
	res := &User{}
	err := db.Get(res, sqlstr, Id)
	if err != nil {
		return nil, err
	}
	return res, err
}

func GetUserByMobile(mobile string) (*User, error) {
	db := mysql.GetDb()
	sqlstr := "select id,openid,mobile,password,nickname,sex,age,avatar,sign,is_disabled from ucenter_member where mobile=? "
	res := &User{}
	err := db.Get(res, sqlstr, mobile)
	if err != nil {
		return nil, err
	}
	return res, err
}

func UserLogin(u *User) (*User, error) {
	db := mysql.GetDb()
	sqlstr := "select id,openid,mobile,password,nickname,sex,age,avatar,sign,is_disabled from ucenter_member where is_deleted !=1 and mobile=? "
	res := &User{}
	err := db.Get(res, sqlstr, u.Mobile)
	if err != nil {
		return nil, err
	}
	return res, err
}

func AdminUserLogin(u *AdminUser) (*AdminUser, error) {
	db := mysql.GetDb()
	sqlstr := "select id,password,auth from admin_user where username = ? "
	res := &AdminUser{}
	err := db.Get(res, sqlstr, u.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, nil
		}
		return nil, err
	}
	return res, err
}

func Register(u *User) error {
	db := mysql.GetDb()
	sqlstr := "insert into ucenter_member(id,mobile,password,nickname,gmt_create,gmt_modified) values (?,?,?,?,?,?)"
	date := time.Now()
	datestr := date.Format("2006-01-02 15:04:05")
	_, err := db.Exec(sqlstr, utils.GenID(), u.Mobile, u.Password, u.Nickname, datestr, datestr)
	if err != nil {
		return err
	}
	return nil
}

func AddOrd(order *Order) error {
	db := mysql.GetDb()
	sqlstr := "insert into" +
		" t_order(id,order_no,course_id,course_title,course_cover,teacher_name,member_id,nickname,mobile,total_fee,pay_type,status,is_deleted,gmt_create,gmt_modified) " +
		"values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	date := time.Now()
	datestr := date.Format("2006-01-02 15:04:05")
	_, err := db.Exec(sqlstr, order.Id, order.OrderNo, order.CourseId, order.CourseTitle, order.CourseCover, order.TeacherName, order.MemberId, order.NickName, order.Mobile, order.TotalFee, order.PayType, order.Status, order.IsDeleted, datestr, datestr)
	if err != nil {
		return err
	}
	return nil
}

func GetOrdById(ordid string) (*Order, error) {
	db := mysql.GetDb()
	sqlstr := "select id,order_no,course_id,course_title,course_cover,teacher_name,member_id,nickname,mobile,total_fee,pay_type,status from t_order where order_no=?  "
	ord := &Order{}
	err := db.Get(ord, sqlstr, ordid)
	if err != nil {
		return nil, err
	}
	return ord, nil
}

func GetOrdStatus(ordid string) (int, error) {
	db := mysql.GetDb()
	sqlstr := "select status from t_order where order_no=?  "
	status := 0
	err := db.Get(&status, sqlstr, ordid)
	if err != nil {
		return 0, err
	}
	return status, nil
}

func SetOrdStatus(ordid string) error {
	db := mysql.GetDb()
	sqlstr := "update t_order set status=1  where order_no=?  "
	_, err := db.Exec(sqlstr, ordid)
	if err != nil {
		return err
	}
	return nil
}

func CreatPayLog(o *Order) error {
	db := mysql.GetDb()
	sqlstr := "insert into" +
		" t_pay_log(id,order_no,total_fee,pay_type,trade_state,transaction_id,attr,pay_time,gmt_create,gmt_modified) " +
		"values (?,?,?,?,?,?,?,?,?,?)"
	date := time.Now()
	datestr := date.Format("2006-01-02 15:04:05")
	_, err := db.Exec(sqlstr, utils.GenID(), o.OrderNo, o.TotalFee, 1, "SUCCESS", "123456789", "nil", datestr, datestr, datestr)
	if err != nil {
		return err
	}
	return nil
}

func UserIsBuyCourse(userid, courseid string) (bool, error) {
	db := mysql.GetDb()
	sqlstr := "select count(1) from t_order where course_id=? and member_id =? and status=1  "
	num := 0
	err := db.Get(&num, sqlstr, courseid, userid)
	if err != nil {
		return false, err
	}
	return num > 0, nil
}
