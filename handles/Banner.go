package handles

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goadmin/constants"
	"goadmin/dao"
	"net/http"
	"sort"
	"strconv"
)

func PageGetBanner(c *gin.Context) {

	page := c.Param("page")
	limit := c.Param("limit")
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	banners, total, err := dao.GetPageBanner(pageInt, limitInt)
	resp := constants.ReCode{}
	if err != nil {
		fmt.Println("GetPageBanner err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	resp.Data = gin.H{"total": total, "items": banners}
	c.JSON(http.StatusOK, resp)
}

func GetBanner(c *gin.Context) {
	id := c.Param("id")
	resp := constants.ReCode{}
	b, err := dao.GetBannerById(id)
	if err != nil {
		fmt.Println("DelBanner err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	resp.Data = gin.H{"banner": b}
	c.JSON(http.StatusOK, resp)
}

func AddBanner(c *gin.Context) {
	addbanner := &dao.Banner{}
	err := c.ShouldBindJSON(&addbanner)
	resp := constants.ReCode{}
	if err != nil {
		fmt.Println("ShouldBindJSON err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = dao.AddBanner(addbanner)
	if err != nil {
		fmt.Println("AddBanner err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	c.JSON(http.StatusOK, resp)
}

func UpdateBanner(c *gin.Context) {
	updatebanner := &dao.Banner{}
	err := c.ShouldBindJSON(&updatebanner)
	resp := constants.ReCode{}
	if err != nil {
		fmt.Println("ShouldBindJSON err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = dao.UpdateBanner(updatebanner)
	if err != nil {
		fmt.Println("UpdateBanner err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	c.JSON(http.StatusOK, resp)
}

func DelBanner(c *gin.Context) {
	id := c.Param("id")
	resp := constants.ReCode{}
	err := dao.DelBanner(id)
	if err != nil {
		fmt.Println("DelBanner err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	c.JSON(http.StatusOK, resp)
}

func GetAllBanner(c *gin.Context) {
	resp := constants.ReCode{}
	banners, err := dao.GetAllBanner()
	if err != nil {
		fmt.Println("GetBannerAll err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	resp.Data = gin.H{"list": banners}
	c.JSON(http.StatusOK, resp)
}

func IndexFront(c *gin.Context) {

	courselist, err := dao.GetCourse()
	resp := constants.ReCode{}
	if err != nil {
		fmt.Println("get course err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	sort.Slice(courselist, func(i, j int) bool {
		return courselist[i].Id > courselist[j].Id
	})
	if len(courselist) > 8 {
		courselist = courselist[:8]
	} else {
		courselist = courselist[:len(courselist)]
	}
	teacherList, err := dao.TeacherList()
	if err != nil {
		fmt.Println("get TeacherList err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	sort.Slice(teacherList, func(i, j int) bool {
		return teacherList[i].Id > teacherList[j].Id
	})
	if len(teacherList) > 4 {
		teacherList = teacherList[:4]
	} else {
		teacherList = teacherList[:len(teacherList)]
	}
	resp.Ok()
	resp.Data = gin.H{"eduList": courselist, "teacherList": teacherList}
	c.JSON(http.StatusOK, resp)
}
