package dao

import (
	"fmt"
	"goadmin/dao/mysql"
)

func PageFrontCourse(page, limit int, cfv *CourseFrontVo) ([]*CourseInfo, int, error) {

	Clist := make([]*CourseInfo, 0, 10)
	db := mysql.GetDb()
	qs := "select id,teacher_id,subject_id,subject_parent_id,title,price,lesson_num,cover,gmt_create,buy_count from edu_course where is_deleted!=1 "

	numqs := "select count(1) from edu_course where is_deleted!=1 "

	whereqs := ""
	if cfv.SubjectParentId != "" || cfv.SubjectId != "" {
		whereqs += " and  "
		if cfv.SubjectParentId != "" {
			n := "\"" + cfv.SubjectParentId + "\""
			whereqs += " subject_parent_id= " + n + " and "
		}
		if cfv.SubjectId != "" {
			n := "\"" + cfv.SubjectId + "\""
			whereqs += " subject_id = " + n + " and "
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
	err = db.Select(&Clist, qs, (page-1)*limit, limit)
	if err != nil {
		fmt.Println("select Teacher err:", err)
		return nil, 0, err
	}
	return Clist, total, nil
}
