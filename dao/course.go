package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"goadmin/constants"
	"goadmin/dao/mysql"
	"goadmin/utils"
	"strconv"
	"time"
)

func GetCourseByTile(title string) (*Course, error) {
	db := mysql.GetDb()
	sqlstr := "select * from edu_subject where title =?"
	c := &Course{}
	err := db.Get(c, sqlstr, title)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println("Get Course err:", err)
		return nil, err
	}
	return c, nil
}

func AddCourse(c *Course) error {
	db := mysql.GetDb()
	sqlstr := "insert into edu_subject (id,title,parent_id,sort,gmt_create,gmt_modified) values(?,?,?,?,?,?)"
	date := time.Now()
	datestr := date.Format("2006-01-02 15:04:05")
	_, err := db.Exec(sqlstr, c.Id, c.Title, c.ParentId, c.Sort, datestr, datestr)
	if err != nil {
		fmt.Println("Course insert err:", err)
		return err
	}
	return nil
}

func AddCourseForExcel(cm map[string][]string) error {

	//从map里添加找出然后逐个添加
	for k, v := range cm {
		parentCourse, err := GetCourseByTile(k)
		if err != nil {
			fmt.Println("Insert Course err:", err)
			return err
		}
		//一级目录存在直接构造插入数据库
		if parentCourse != nil {
			for _, title := range v {
				vc := &Course{
					Id:       utils.GenID(),
					Title:    title,
					ParentId: parentCourse.Id,
					Sort:     0,
				}
				err := AddCourse(vc)
				if err != nil {
					fmt.Println("Insert Course err:", err)
					return err
				}
			}
		} else {
			parentc := &Course{
				Id:       utils.GenID(),
				Title:    k,
				ParentId: "0",
				Sort:     0,
			}
			err := AddCourse(parentc)
			if err != nil {
				fmt.Println("Insert Course err:", err)
				return err
			}

			for _, title := range v {
				vc := &Course{
					Id:       utils.GenID(),
					Title:    title,
					ParentId: parentc.Id,
					Sort:     0,
				}
				err := AddCourse(vc)
				if err != nil {
					fmt.Println("Insert Course err:", err)
					return err
				}
			}
		}
	}

	return nil
}

func GetCourseList() ([]*CourseTree, error) {

	db := mysql.GetDb()

	CTL := make([]*CourseTree, 0, 5)
	CNL := make([]*CourseNode, 0, 50)
	onemap := make(map[string]*CourseTree)
	//查询出所有一级标题
	sqlstr := "select id,title from edu_subject where parent_id = 0"
	err := db.Select(&CTL, sqlstr)
	if err != nil {
		fmt.Println("select err:", err)
		return nil, err
	}
	for _, v := range CTL {
		onemap[v.Id] = v
		onemap[v.Id].Children = make([]*CourseNode, 0, 5)
	}
	//查出所以二级标题
	sqlstr = "select id,title,parent_id from edu_subject where parent_id != 0"
	err = db.Select(&CNL, sqlstr)
	if err != nil {
		fmt.Println("select err:", err)
		return nil, err
	}

	//二合一
	for _, v := range CNL {
		onemap[v.ParentId].Children = append(onemap[v.ParentId].Children, v)
	}

	return CTL, err
}

func AddCourseInfo(CIF *CourseInfo) (string, error) {
	db := mysql.GetDb()
	date := time.Now()
	datestr := date.Format("2006-01-02 15:04:05")

	cid := utils.GenID()
	sqlstr := "insert into edu_course (id,teacher_id,subject_id,subject_parent_id,title,price,lesson_num,cover,gmt_create,gmt_modified) values(?,?,?,?,?,?,?,?,?,?)"
	_, err := db.Exec(sqlstr, cid, CIF.TeacherId, CIF.SubjectId, CIF.SubjectParentId, CIF.Title, CIF.Price, CIF.LessonNum, CIF.Cover, datestr, datestr)

	if err != nil {
		fmt.Println("insert into edu_course err:", err)
		return "", err
	}

	sqlstr = "insert into edu_course_description (id,description,gmt_create,gmt_modified) values (?,?,?,?)"

	_, err = db.Exec(sqlstr, cid, CIF.Description, datestr, datestr)

	if err != nil {
		fmt.Println("insert into edu_course_description err:", err)
		return "", err
	}
	return cid, nil
}

func GetCourseInfoByid(courseid string) (*CourseInfo, error) {
	db := mysql.GetDb()
	resCIF := &CourseInfo{}
	sqlstr := "select id,teacher_id,subject_id,subject_parent_id,title,price,lesson_num,cover from edu_course where id =?"
	err := db.Get(resCIF, sqlstr, courseid)
	if err != nil {
		fmt.Println("select err:", err)
		return nil, err
	}
	sqlstr = "select description from edu_course_description where id=?"
	cur := &CourseDescription{}
	err = db.Get(cur, sqlstr, courseid)
	if err != nil {
		fmt.Println("select err:", err)
		return nil, err
	}
	resCIF.Description = cur.Description
	return resCIF, nil
}

func GetCourseFront(CourseId string) (*CourseFront, error) {
	db := mysql.GetDb()

	res := &CourseFront{}
	sqlstr := "select id,teacher_id,subject_id,subject_parent_id,title,price,lesson_num,cover,buy_count,view_count from edu_course where id =?"
	err := db.Get(res, sqlstr, CourseId)
	if err != nil {
		return nil, err
	}
	sqlstr = "select description from edu_course_description where id=?"
	err = db.Get(&res.Description, sqlstr, CourseId)
	if err != nil {
		return nil, err
	}
	sqlstr = "select title from edu_subject where id=?"
	err = db.Get(&res.SubjectLevelOne, sqlstr, res.SubjectParentId)
	if err != nil {
		return nil, err
	}
	err = db.Get(&res.SubjectLevelTwo, sqlstr, res.SubjectId)
	if err != nil {
		return nil, err
	}
	tid, _ := strconv.Atoi(res.TeacherId)
	teacher, err := GetTeacherById(tid)
	if err != nil {
		return nil, err
	}
	res.TeacherName = teacher.Name
	res.TeacherAvatar = teacher.Avatar

	return res, nil
}

func GetCoursesByTeacherId(teacherid string) ([]*CourseInfo, error) {
	db := mysql.GetDb()
	res := make([]*CourseInfo, 0, 10)
	sqlstr := "select id,teacher_id,subject_id,subject_parent_id,title,price,lesson_num,cover from edu_course where teacher_id =?"
	err := db.Select(&res, sqlstr, teacherid)
	if err != nil {
		fmt.Println("select err:", err)
		return nil, err
	}
	//sqlstr = "select description from edu_course_description where id=?"
	//cur := &CourseDescription{}
	//err = db.Select(cur, sqlstr, teacherid)
	//if err != nil {
	//	fmt.Println("select err:", err)
	//	return nil, err
	//}
	//res.Description = cur.Description
	return res, nil
}

func UpdateCourseInfo(updatecif *CourseInfo) error {
	db := mysql.GetDb()
	sqlstr := "update edu_course set teacher_id=?,subject_id=?,subject_parent_id=?,title=?,price=?,lesson_num=?,cover=?  where id=?"
	_, err := db.Exec(sqlstr, updatecif.TeacherId, updatecif.SubjectId, updatecif.SubjectParentId, updatecif.Title, updatecif.Price, updatecif.LessonNum, updatecif.Cover, updatecif.Id)
	if err != nil {
		fmt.Println("update err:", err)
		return err
	}
	sqlstr = "update edu_course_description set description = ? where id =?"
	_, err = db.Exec(sqlstr, updatecif.Description, updatecif.Id)
	if err != nil {
		fmt.Println("update err:", err)
		return err
	}
	return nil
}

func GetChapterVideo(couserid string) ([]*ChapterTree, error) {

	ct := make([]*ChapterTree, 0, 5)
	db := mysql.GetDb()
	sqlstr := "select id,title from edu_chapter where course_id=?"
	err := db.Select(&ct, sqlstr, couserid)
	if err != nil {
		fmt.Println("select err:", err)
		return nil, err
	}
	chapterMap := make(map[string]*ChapterTree, len(ct))

	for _, v := range ct {
		chapterMap[v.Id] = v
		chapterMap[v.Id].Children = make([]*ChapterNode, 0, 5)
	}
	CNT := make([]*ChapterNode, 0, 10)
	sqlstr = "select id,title,chapter_id,video_source_id from edu_video where course_id=?"
	err = db.Select(&CNT, sqlstr, couserid)
	if err != nil {
		fmt.Println("select err:", err)
		return nil, err
	}

	for _, v := range CNT {
		chapterMap[v.ChapterId].Children = append(chapterMap[v.ChapterId].Children, v)
	}
	return ct, nil
}

func AddCourseChapter(ct *Chapter) error {
	db := mysql.GetDb()
	sqlstr := "insert into edu_chapter (id,course_id,title,sort,gmt_create,gmt_modified) values(?,?,?,?,?,?)"
	date := time.Now()
	datestr := date.Format("2006-01-02 15:04:05")
	_, err := db.Exec(sqlstr, utils.GenID(), ct.CourseId, ct.Title, ct.Sort, datestr, datestr)
	if err != nil {
		return err
	}
	return nil
}
func UpdateCourseChapter(ct *Chapter) error {
	db := mysql.GetDb()
	sqlstr := "update edu_chapter set course_id=?,title,sort=? , sort =? where id =?"
	_, err := db.Exec(sqlstr, ct.CourseId, ct.Title, ct.Sort, ct.Id)
	if err != nil {
		return err
	}
	return nil
}

func GetPublishCourseInfo(courseid string) (*PublicCourseInfo, error) {
	db := mysql.GetDb()
	courseInfo, err := GetCourseInfoByid(courseid)
	if err != nil {
		return nil, err
	}
	pcf := &PublicCourseInfo{
		Id:        courseInfo.Id,
		Title:     courseInfo.Title,
		Cover:     courseInfo.Cover,
		LessonNum: courseInfo.LessonNum,
		Price:     courseInfo.Price,
	}
	sqlstr := "select name from edu_teacher where id=?"
	err = db.Get(&pcf.TeacherName, sqlstr, courseInfo.TeacherId)
	if err != nil {
		fmt.Println("Get TeacherName err:", err)
		return nil, err
	}
	sqlstr = "select title from edu_subject where id=?"
	err = db.Get(&pcf.SubjectLevelOne, sqlstr, courseInfo.SubjectParentId)
	if err != nil {
		return nil, err
	}
	err = db.Get(&pcf.SubjectLevelTwo, sqlstr, courseInfo.SubjectId)
	if err != nil {
		return nil, err
	}
	return pcf, err
}

func PublishCourse(courseid string) error {
	db := mysql.GetDb()
	sqlstr := "update edu_course set status = ? where id=?"
	_, err := db.Exec(sqlstr, "Normal", courseid)
	if err != nil {
		return err
	}
	return nil
}

func GetCourseCharpter(courseid string) (*Chapter, error) {
	db := mysql.GetDb()
	sqlstr := "select id,title,course_id,sort from edu_chapter where id=?"
	ct := &Chapter{}
	err := db.Get(ct, sqlstr, courseid)
	if err != nil {
		return nil, err
	}
	return ct, nil
}
func GetCourse() ([]*EduCourse, error) {
	db := mysql.GetDb()
	sqlstr := "select id ,title,status,lesson_num from edu_course"
	listeduc := make([]*EduCourse, 0, 10)
	err := db.Select(&listeduc, sqlstr)
	if err != nil {
		return nil, err
	}
	return listeduc, nil
}
func GetCourseFromTeacher(teacherid string) ([]*EduCourse, error) {
	db := mysql.GetDb()
	sqlstr := "select id ,title,status,lesson_num from edu_course where teacher_id=? "
	listeduc := make([]*EduCourse, 0, 10)
	fmt.Println(sqlstr)
	err := db.Select(&listeduc, sqlstr, teacherid)
	if err != nil {
		return nil, err
	}
	return listeduc, nil
}

func DelCourse(courseId string) error {
	tx, err := mysql.GetDb().Begin() // 开启事务
	if err != nil {
		fmt.Printf("begin trans failed, err:%v\n", err)
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			fmt.Println("rollback")
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
			fmt.Println("commit")
		}
	}()
	sqlstr := "delete from edu_video where course_id=?"
	_, err = tx.Exec(sqlstr, courseId)
	if err != nil {
		return err
	}
	sqlstr = "delete from edu_chapter where course_id=?"
	_, err = tx.Exec(sqlstr, courseId)
	if err != nil {
		return err
	}
	sqlstr = "delete from edu_course_description where id=?"
	_, err = tx.Exec(sqlstr, courseId)
	if err != nil {
		return err
	}
	sqlstr = "delete from edu_course where id=?"
	_, err = tx.Exec(sqlstr, courseId)
	if err != nil {
		return err
	}
	return nil
}

func GetVideoIds(courseId string) ([]string, error) {
	db := mysql.GetDb()
	sqlstr := "select video_source_id from edu_video where course_id= ? "
	res := []string{}
	err := db.Select(&res, sqlstr, courseId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DelChapterById(chapterid string) error {
	db := mysql.GetDb()
	sqlstr := "select count(1)  from edu_video where chapter_id=? "
	nums := 0
	err := db.Get(&nums, sqlstr, chapterid)
	if err != nil {
		return err
	}
	if nums > 0 {
		return errors.New(constants.SqlErrChapterHaveVideo)
	} else {
		sqlstr = "delete from edu_chapter where id=?"
		_, err := db.Exec(sqlstr, chapterid)
		if err != nil {
			return err
		}
	}
	return nil
}

func AddVideo(vf *VideoInfo) error {
	db := mysql.GetDb()
	date := time.Now()
	datestr := date.Format("2006-01-02 15:04:05")

	sqlstr := "insert into edu_video(id,course_id,chapter_id,title,video_source_id,video_original_name,sort,gmt_create,gmt_modified) values(?,?,?,?,?,?,?,?,?)"
	_, err := db.Exec(sqlstr, utils.GenID(), vf.CourseID, vf.ChapterId, vf.Title, vf.VideoSourceId, vf.VideoOriginalName, vf.Sort, datestr, datestr)
	if err != nil {
		return err
	}
	return nil

}

func DelVideo(id string) error {
	db := mysql.GetDb()
	sqlstr := "delete from edu_video where id=? "
	_, err := db.Exec(sqlstr, id)
	if err != nil {
		return err
	}
	return nil
}
