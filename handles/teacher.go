package handles

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goadmin/constants"
	"goadmin/dao"
	ossHandle "goadmin/handles/osshandle"
	"net/http"
	"sort"
	"strconv"
)

func TeacherList(c *gin.Context) {

	teacherList, err := dao.TeacherList()
	resp := constants.ReCode{}
	if err != nil {
		fmt.Println("get teacher list err", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = gin.H{"items": teacherList}
	resp.Ok()
	c.JSON(http.StatusOK, resp)
}

func PageTeacherCondition(c *gin.Context) {

	pagestr := c.Param("page")
	limitstr := c.Param("limit")

	Qteacher := &dao.Teacher{}
	err := c.ShouldBindJSON(&Qteacher)
	if err != nil {
		resp := constants.ReCode{
			Code:    20000,
			Success: false,
			Message: "查询失败",
			Data:    gin.H{"rows": nil, "total": 0},
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	page, limit := 1, 10
	page, _ = strconv.Atoi(pagestr)
	limit, _ = strconv.Atoi(limitstr)
	if page == 0 {
		page = 1
	}
	teachers, total, err := dao.GetPageTeacherCondition(page, limit, Qteacher)
	if err != nil {
		resp := constants.ReCode{
			Code:    20000,
			Success: false,
			Message: "查询失败",
			Data:    gin.H{"rows": nil, "total": 0},
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	sort.Slice(teachers, func(i, j int) bool {
		return teachers[i].Sort > teachers[j].Sort
	})
	resp := constants.ReCode{
		Code:    20000,
		Success: true,
		Message: "查询成功",
		Data:    gin.H{"rows": teachers, "total": total},
	}
	c.JSON(http.StatusOK, resp)
}

func DelTeacherId(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	err := dao.DelTeacherId(id)
	resp := constants.ReCode{}
	if err != nil {
		fmt.Println("del teacher err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	c.JSON(http.StatusOK, resp)
}

func AddTeacher(c *gin.Context) {

	t := &dao.Teacher{}
	err := c.ShouldBindJSON(&t)
	resp := constants.ReCode{}
	if err != nil {
		fmt.Println(err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = dao.AddTeacher(t)
	if err != nil {
		fmt.Println(err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp.Ok()
	c.JSON(http.StatusOK, resp)
}

func GetTeacher(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	teacher, err := dao.GetTeacherById(id)
	resp := constants.ReCode{}
	if err != nil {
		fmt.Println("getteacher err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	resp.Data = gin.H{"teacher": teacher}
	c.JSON(http.StatusOK, resp)
}

func UpdateTeacher(c *gin.Context) {

	t := &dao.Teacher{}
	err := c.ShouldBindJSON(&t)
	resp := constants.ReCode{}
	if err != nil {
		fmt.Println(err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
	}
	err = dao.UpdateTeacher(t)
	if err != nil {
		fmt.Println("update teacher err", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
	}
	resp.Ok()
	c.JSON(http.StatusOK, resp)
}

func UploadAvatar(c *gin.Context) {

	f, err := c.FormFile("avatar")
	resp := constants.ReCode{}
	if err != nil {
		fmt.Println("gin get file err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	url, err := ossHandle.OssUpload(f)
	if err != nil {
		fmt.Println("ossHandle.OssUpload err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp.Ok()
	resp.Data = gin.H{"url": url}
	c.JSON(http.StatusOK, resp)
}

func PageTeacher(c *gin.Context) {
	pagestr := c.Param("page")
	limitstr := c.Param("limit")
	page, limit := 1, 10
	page, _ = strconv.Atoi(pagestr)
	limit, _ = strconv.Atoi(limitstr)
	if page == 0 {
		page = 1
	}
	teachers, total, err := dao.PageTeacher(page, limit)
	if err != nil {
		resp := constants.ReCode{
			Code:    20001,
			Success: false,
			Message: "查询失败",
			Data:    gin.H{"rows": nil, "total": 0},
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	resp := constants.ReCode{
		Code:    20000,
		Success: true,
		Message: "查询成功",
		Data:    gin.H{"items": teachers, "total": total, "current": page, "pages": total / limit, "size": total/limit + 1, "hasPrevious": page-1 > 0, "hasNext": page*limit < total},
	}
	c.JSON(http.StatusOK, resp)
}

func GetTeacherInfo(c *gin.Context) {
	teacherid := c.Param("id")
	tid, _ := strconv.Atoi(teacherid)
	resp := constants.ReCode{}
	teacher, err := dao.GetTeacherById(tid)

	if err != nil {
		fmt.Println("GetTeacher err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	courseInfos, err := dao.GetCoursesByTeacherId(teacherid)
	if err != nil {
		fmt.Println("GetCoursesByTeacherId err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	resp.Data = gin.H{"teacher": teacher, "courseList": courseInfos}
	c.JSON(http.StatusOK, resp)

}
