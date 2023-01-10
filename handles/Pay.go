package handles

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goadmin/constants"
	"goadmin/dao"
	"goadmin/utils"
	"net/http"
	"time"
)

func CreatCourseOrder(c *gin.Context) {

	//得到用户id和信息
	courseid := c.Param("courseid")
	resp := constants.ReCode{}
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
	user, err := dao.GetUserById(claims.Id)
	if err != nil {
		fmt.Println("GetUserById ERR", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}

	//得到课程信息
	courseInfo, err := dao.GetCourseFront(courseid)
	if err != nil {
		fmt.Println("GetCourseInfoByid ERR", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}

	//构造记录
	ord := &dao.Order{
		Id:          utils.GenID(),
		OrderNo:     utils.GenID(),
		CourseId:    courseid,
		CourseTitle: courseInfo.Title,
		CourseCover: courseInfo.Cover,
		TeacherName: courseInfo.TeacherName,
		MemberId:    user.Id,
		NickName:    user.Nickname,
		Mobile:      user.Mobile,
		TotalFee:    courseInfo.Price,
		PayType:     1,
		Status:      0,
	}
	err = dao.AddOrd(ord)
	if err != nil {
		fmt.Println("AddOrd ERR", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	resp.Data = gin.H{"orderId": ord.OrderNo}
	c.JSON(http.StatusOK, resp)
}

func GetOrdById(c *gin.Context) {

	ordid := c.Param("ordid")
	resp := constants.ReCode{}
	order, err := dao.GetOrdById(ordid)
	if err != nil {
		fmt.Println("GetOrdById ERR", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	resp.Data = gin.H{"item": order}
	c.JSON(http.StatusOK, resp)
}

func CreatNative(c *gin.Context) {

	orderno := c.Param("orderNo")
	resp := constants.ReCode{}
	order, err := dao.GetOrdById(orderno)
	if err != nil {
		fmt.Println("GetOrdById ERR", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	resp.Data = gin.H{"out_trade_no": orderno, "course_id": order.CourseId, "total_fee": order.TotalFee,
		"result_code": 200, "code_url": "hello world"}
	c.JSON(http.StatusOK, resp)
}

func GetOrdStatus(c *gin.Context) {
	orderno := c.Param("orderNo")
	resp := constants.ReCode{}
	//status, err := dao.GetOrdStatus(orderno)
	//if err != nil {
	//	fmt.Println("GetOrdById ERR", err)
	//	resp.Err()
	//	c.JSON(http.StatusOK, resp)
	//	return
	//}
	err := dao.SetOrdStatus(orderno)
	if err != nil {
		fmt.Println("SetOrdStatus ERR", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	order, err := dao.GetOrdById(orderno)
	if err != nil {
		fmt.Println("GetOrdById ERR", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = dao.CreatPayLog(order)
	if err != nil {
		fmt.Println("CreatPayLog ERR", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	time.Sleep(2 * time.Second)
	resp.Ok()
	resp.Message = "支付成功"
	c.JSON(http.StatusOK, resp)
}
