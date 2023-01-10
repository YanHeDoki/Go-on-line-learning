package dao

import (
	"fmt"
	"goadmin/dao/mysql"
	"goadmin/utils"
	"strconv"
	"time"
)

func TeacherTotal() (int, error) {
	db := mysql.GetDb()
	qs := "select count(1) from edu_teacher"
	res := 0
	err := db.Get(&res, qs)
	if err != nil {
		fmt.Println("select Teacher err:", err)
		return 0, err
	}
	return res, nil
}

func TeacherList() ([]*Teacher, error) {
	tlist := make([]*Teacher, 0, 10)
	qs := "select id,name,intro,career,sort,level from edu_teacher where is_deleted!=1 "
	db := mysql.GetDb()
	err := db.Select(&tlist, qs)
	if err != nil {
		fmt.Println("get teacher err:", err)
		return nil, err
	}
	return tlist, nil
}

func GetPageTeacherCondition(page, limit int, t *Teacher) ([]*Teacher, int, error) {
	tlist := make([]*Teacher, 0, 10)
	db := mysql.GetDb()
	qs := "select id,name,intro,career,sort,level from edu_teacher where is_deleted!=1 "
	numqs := "select count(1) from edu_teacher where is_deleted!=1 "

	whereqs := ""
	if t.Name != "" || t.Level != 0 {
		whereqs += " and  "
		if t.Name != "" {
			n := "\"%" + t.Name + "%\""
			whereqs += " name like " + n + " and "
		}
		if t.Level != 0 {
			l := strconv.Itoa(t.Level)
			whereqs += " level = " + l + " and "
		}
		whereqs += "1"
		qs += whereqs
		numqs += whereqs
	}
	total := 0
	err := db.Get(&total, numqs)
	if err != nil {
		fmt.Println(qs)
		fmt.Println("select totalnum err:", err)
		return nil, 0, err
	}

	qs += " limit ?,?"
	//qs := "select id,name,intro,career,sort,level from edu_teacher limit ?,? "
	err = db.Select(&tlist, qs, (page-1)*limit, limit)
	if err != nil {
		fmt.Println("select Teacher err:", err)
		return nil, 0, err
	}
	return tlist, total, nil
}

func DelTeacherId(id int) error {
	db := mysql.GetDb()
	sqlstr := "update edu_teacher set is_deleted = 1 where id=?"

	_, err := db.Exec(sqlstr, id)
	if err != nil {
		fmt.Println("del teacher err:", err)
		return err
	}
	return nil
}

func AddTeacher(t *Teacher) error {
	db := mysql.GetDb()
	sqlstr := "insert into edu_teacher (id,name,avatar,sort,level,career,intro,gmt_create,gmt_modified) values(?,?,?,?,?,?,?,?,?)"
	id := utils.GenID()
	date := time.Now()
	datestr := date.Format("2006-01-02 15:04:05")
	_, err := db.Exec(sqlstr, id, t.Name, t.Avatar, t.Sort, t.Level, t.Career, t.Intro, datestr, datestr)
	if err != nil {
		fmt.Println("add err:", err)
		return err
	}
	return nil
}

func GetTeacherById(Id int) (*Teacher, error) {
	db := mysql.GetDb()
	sqlstr := "select id,name,intro,career,sort,level,avatar from edu_teacher where id=?"
	t := &Teacher{}
	err := db.Get(t, sqlstr, Id)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func UpdateTeacher(t *Teacher) error {
	db := mysql.GetDb()
	sqlstr := "update edu_teacher set name =?,sort=?,level=?,intro=?,career=? where id=?"

	_, err := db.Exec(sqlstr, t.Name, t.Sort, t.Level, t.Intro, t.Career, t.Id)
	if err != nil {
		return err
	}
	return nil

}

func PageTeacher(page, limit int) ([]*Teacher, int, error) {
	tlist := make([]*Teacher, 0, 10)
	db := mysql.GetDb()

	qs := "select id,name,intro,career,sort,level from edu_teacher where is_deleted!=1 "
	numqs := "select count(1) from edu_teacher where is_deleted!=1 "

	total := 0
	err := db.Get(&total, numqs)
	if err != nil {
		fmt.Println(qs)
		fmt.Println("select totalnum err:", err)
		return nil, 0, err
	}

	qs += " limit ?,?"
	//qs := "select id,name,intro,career,sort,level from edu_teacher limit ?,? "
	err = db.Select(&tlist, qs, (page-1)*limit, limit)
	if err != nil {
		fmt.Println("select Teacher err:", err)
		return nil, 0, err
	}
	return tlist, total, nil
}
