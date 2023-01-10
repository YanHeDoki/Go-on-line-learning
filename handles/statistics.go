package handles

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goadmin/constants"
	"goadmin/dao"
	"math/rand"
	"net/http"
)

func CountRegister(c *gin.Context) {

	day := c.Param("day")
	resp := constants.ReCode{}
	count, err := dao.CountRegister(day)
	if err != nil {
		fmt.Println("CountRegister err", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	sd := &dao.StatisticsDaily{
		DateCalculated: day,
		RegisterNum:    count,
		LoginNum:       rand.Intn(200),
		VideoViewNum:   rand.Intn(200),
		CourseNum:      rand.Intn(200),
	}
	err = dao.InsertStatistics(day, sd)
	if err != nil {
		fmt.Println("InsertStatistics err", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	resp.Data = gin.H{"countRegister": count}
	c.JSON(http.StatusOK, resp)
}

func ShowData(c *gin.Context) {
	datetype := c.Param("type")
	start := c.Param("start")
	end := c.Param("end")
	resp := constants.ReCode{}
	showlist, err := dao.GetShowData(datetype, start, end)
	if err != nil {
		fmt.Println("GetShowData err", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	datelist := make([]string, 0, 10)
	numlist := make([]int, 0, 10)
	for _, v := range showlist {
		datelist = append(datelist, v.DateCalculated)

		switch datetype {
		case "login_num":
			numlist = append(numlist, v.LoginNum)
		case "register_num":
			numlist = append(numlist, v.RegisterNum)
		case "video_view_num":
			numlist = append(numlist, v.VideoViewNum)
		case "course_num":
			numlist = append(numlist, v.CourseNum)
		}
	}
	resp.Ok()
	resp.Data = gin.H{"date_calculatedList": datelist, "numDataList": numlist}
	c.JSON(http.StatusOK, resp)
}
