package handles

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goadmin/constants"
	"goadmin/dao"
	ossHandle "goadmin/handles/osshandle"
	"goadmin/utils"
	"net/http"
)

func Addcourse(c *gin.Context) {

	excelfile, err := c.FormFile("courseexcel")
	resp := constants.ReCode{}
	if err != nil {
		fmt.Println("Get FormFile err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	courseMap, err := utils.ReadExcel(excelfile)
	if err != nil {
		fmt.Println("Excel err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = dao.AddCourseForExcel(courseMap)
	if err != nil {
		fmt.Println("insert db err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	c.JSON(http.StatusOK, resp)
}

func CourseList(c *gin.Context) {

	resp := constants.ReCode{}

	courseList, err := dao.GetCourseList()
	if err != nil {
		fmt.Println("get course err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp.Ok()
	resp.Data = gin.H{"list": courseList}
	c.JSON(http.StatusOK, resp)
}

func AddCourseInfo(c *gin.Context) {

	cif := &dao.CourseInfo{}
	err := c.ShouldBindJSON(&cif)
	resp := &constants.ReCode{}
	if err != nil {
		fmt.Println("ShouldBindJSON err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	cid, err := dao.AddCourseInfo(cif)
	if err != nil {
		fmt.Println("add err:", err)
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	resp.Data = gin.H{"courseId": cid}
	c.JSON(http.StatusOK, resp)
}

func GetChapterVideo(c *gin.Context) {
	courseid := c.Param("courseid")
	resp := constants.ReCode{}
	chapterVideo, err := dao.GetChapterVideo(courseid)
	if err != nil {
		fmt.Println("select err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	resp.Data = gin.H{"allChapterVideo": chapterVideo}
	c.JSON(http.StatusOK, resp)
}

func GetCourseInfo(c *gin.Context) {

	courseid := c.Param("courseid")
	resp := constants.ReCode{}
	courseInfo, err := dao.GetCourseInfoByid(courseid)
	if err != nil {
		fmt.Println(err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	resp.Data = gin.H{"courseInfoVo": courseInfo}
	c.JSON(http.StatusOK, resp)
}

func GetCourseFront(c *gin.Context) {
	courseid := c.Param("courseid")
	resp := constants.ReCode{}
	courseInfo, err := dao.GetCourseFront(courseid)
	if err != nil {
		fmt.Println(err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	videolist, err := dao.GetChapterVideo(courseid)
	if err != nil {
		fmt.Println(err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}

	authHeader := c.Request.Header.Get("x-token")
	if authHeader == "" {
		fmt.Println("GetHead ERR")
		fmt.Println(c.Request.Header)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	claims, err := utils.ParseToken(authHeader)
	if err != nil {
		fmt.Println("Token parse ERR", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	isBuyCourse, err := dao.UserIsBuyCourse(claims.Id, courseid)
	if err != nil {
		fmt.Println("UserIsBuyCourse ERR", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp.Ok()
	resp.Data = gin.H{"courseWebVo": courseInfo, "chapterVideoList": videolist, "isBuy": isBuyCourse}
	c.JSON(http.StatusOK, resp)
}

func UpdateCourseInfo(c *gin.Context) {
	updatecif := &dao.CourseInfo{}
	err := c.ShouldBindJSON(&updatecif)
	resp := constants.ReCode{}
	if err != nil {
		fmt.Println("ShouldBindJSON err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = dao.UpdateCourseInfo(updatecif)
	if err != nil {
		fmt.Println(err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	c.JSON(http.StatusOK, resp)
}

func GetPublishCourseInfo(c *gin.Context) {
	courseid := c.Param("courseid")
	resp := constants.ReCode{}
	publicCourseInfo, err := dao.GetPublishCourseInfo(courseid)
	if err != nil {
		fmt.Println("get publicCourseInfo err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	resp.Data = gin.H{"publishCourse": publicCourseInfo}

	c.JSON(http.StatusOK, resp)
}

func PublishCourseInfo(c *gin.Context) {

	course_id := c.Param("courseid")
	err := dao.PublishCourse(course_id)
	resp := constants.ReCode{}
	if err != nil {
		fmt.Println("Publish err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
	}
	resp.Ok()
	c.JSON(http.StatusOK, resp)
}

func AddCourseChapter(c *gin.Context) {
	Ct := &dao.Chapter{}
	err := c.ShouldBindJSON(&Ct)
	resp := constants.ReCode{}
	if err != nil {
		fmt.Println("ShouldBindJSON err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = dao.AddCourseChapter(Ct)
	if err != nil {
		fmt.Println("add course chapter err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	c.JSON(http.StatusOK, resp)
}

func UpdateCourseChapter(c *gin.Context) {
	Ct := &dao.Chapter{}
	err := c.ShouldBindJSON(&Ct)
	resp := constants.ReCode{}
	if err != nil {
		fmt.Println("ShouldBindJSON err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = dao.UpdateCourseChapter(Ct)
	if err != nil {
		fmt.Println("update CourseChapter err", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	c.JSON(http.StatusOK, resp)
}

func GetCourseChapter(c *gin.Context) {
	chapterId := c.Param("chapterId")
	resp := constants.ReCode{}
	getCourseCharpter, err := dao.GetCourseCharpter(chapterId)
	if err != nil {
		fmt.Println("GetCourseChapter Err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	resp.Data = gin.H{"chapter": getCourseCharpter}
	c.JSON(http.StatusOK, resp)
}

func GetCourse(c *gin.Context) {

	resp := constants.ReCode{}
	token, exists := c.Get("token")
	if !exists {
		fmt.Println("auth Err:")
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	tokenV := token.(*utils.CustomClaims)
	eduCourses := make([]*dao.EduCourse, 0, 10)
	var err error
	if tokenV.Auth != "1" {
		eduCourses, err = dao.GetCourseFromTeacher(tokenV.Id)
		if err != nil {
			fmt.Println("GetCourse Err:", err)
			resp.Err()
			c.JSON(http.StatusOK, resp)
			return
		}
	} else {
		eduCourses, err = dao.GetCourse()
		if err != nil {
			fmt.Println("GetCourse Err:", err)
			resp.Err()
			c.JSON(http.StatusOK, resp)
			return
		}
	}
	resp.Ok()
	resp.Data = gin.H{"list": eduCourses}
	c.JSON(http.StatusOK, resp)
}
func DelCourse(c *gin.Context) {
	courseid := c.Param("courseid")
	resp := constants.ReCode{}
	ids, err := dao.GetVideoIds(courseid)
	if err != nil {
		fmt.Println("DelCourse err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = ossHandle.DelAliVideoList(ids)
	if err != nil {
		fmt.Println("DelCourse err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = dao.DelCourse(courseid)
	if err != nil {
		fmt.Println("DelCourse err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	c.JSON(http.StatusOK, resp)
}

func DelCourseChapter(c *gin.Context) {
	chapterid := c.Param("chapterid")
	resp := constants.ReCode{}
	err := dao.DelChapterById(chapterid)
	if err != nil {
		resp.Err()
		if err.Error() == constants.SqlErrChapterHaveVideo {
			resp.Message = "不能删除"
			c.JSON(http.StatusOK, resp)
			return
		}
		fmt.Println("Del chapter id err:", err)
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	c.JSON(http.StatusOK, resp)

}

func AddVideo(c *gin.Context) {

	Vinfo := &dao.VideoInfo{}
	resp := constants.ReCode{}
	err := c.ShouldBindJSON(&Vinfo)
	if err != nil {
		fmt.Println("ShouldBindJSON err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = dao.AddVideo(Vinfo)
	if err != nil {
		fmt.Println("Add Video err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	c.JSON(http.StatusOK, resp)
}

func UpVideo(c *gin.Context) {
	file, err := c.FormFile("video")
	resp := constants.ReCode{}
	if err != nil {
		fmt.Println("Feile Get Err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	videofile, err := file.Open()
	if err != nil {
		fmt.Println("Feile open Err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	videoid, err := ossHandle.VideoFileUp(file.Filename, file.Filename, file.Filename, "", videofile)
	if err != nil {
		fmt.Println("VideoFileUp  Err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	resp.Data = gin.H{"videoId": videoid}
	c.JSON(http.StatusOK, resp)
}

func DelVideo(c *gin.Context) {
	id := c.Param("id")
	resp := constants.ReCode{}
	err := dao.DelVideo(id)
	if err != nil {
		fmt.Println("del video err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	c.JSON(http.StatusOK, resp)
}

func DelAliVideo(c *gin.Context) {
	id := c.Param("id")
	resp := constants.ReCode{}
	err := ossHandle.DelAliVideo(id)
	if err != nil {
		fmt.Println("Del AliVideo err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	c.JSON(http.StatusOK, resp)
}

func DelAliVideoList(c *gin.Context) {
	ids := []string{}
	resp := constants.ReCode{}
	err := ossHandle.DelAliVideoList(ids)
	if err != nil {
		fmt.Println("Del AliVideo err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	c.JSON(http.StatusOK, resp)
}

func GetVideoAuth(c *gin.Context) {
	id := c.Param("id")
	playAuth, err := ossHandle.GetPlayAuth(id)
	resp := constants.ReCode{}
	if err != nil {
		fmt.Println("found video auth err", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp.Data = gin.H{"playAuth": playAuth}
	c.JSON(http.StatusOK, resp)

}
