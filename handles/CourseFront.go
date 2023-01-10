package handles

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goadmin/constants"
	"goadmin/dao"
	"net/http"
	"sort"
	"strconv"
	"time"
)

func PageCourse(c *gin.Context) {
	pagestr := c.Param("page")
	limitstr := c.Param("limit")
	page, limit := 1, 10
	page, _ = strconv.Atoi(pagestr)
	limit, _ = strconv.Atoi(limitstr)
	resp := constants.ReCode{}
	if page == 0 {
		page = 1
	}

	CFV := &dao.CourseFrontVo{}
	err := c.ShouldBindJSON(&CFV)
	if err != nil {
		fmt.Println("ShouldBindJSON err", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	courseInfos, total, err := dao.PageFrontCourse(page, limit, CFV)
	if err != nil {
		fmt.Println("PageFrontCourse err", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	if CFV.BuyCountSort != "" {

		sort.Slice(courseInfos, func(i, j int) bool {
			return courseInfos[i].BuyCount > courseInfos[j].BuyCount
		})
	}
	if CFV.PriceSort != "" {
		sort.Slice(courseInfos, func(i, j int) bool {
			return courseInfos[i].Price > courseInfos[j].Price
		})
	}
	if CFV.GmtCreateSort != "" {
		sort.Slice(courseInfos, func(i, j int) bool {
			timei, _ := time.Parse("2006-01-02 15:04:05", courseInfos[i].GmtCreate)
			timej, _ := time.Parse("2006-01-02 15:04:05", courseInfos[j].GmtCreate)
			return timei.Before(timej)
		})
	}
	resp.Ok()
	resp.Data = gin.H{"items": courseInfos, "total": total, "current": page, "pages": total / limit, "size": total/limit + 1, "hasPrevious": page-1 > 0, "hasNext": page*limit < total}
	c.JSON(http.StatusOK, resp)
}
