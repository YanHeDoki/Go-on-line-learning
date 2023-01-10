package handles

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goadmin/constants"
	"goadmin/dao"
	"goadmin/utils"
	"net/http"
)

func SetAuth(c *gin.Context) {
	authHeader := c.Request.Header.Get("x-token")
	if authHeader == "" {
		fmt.Println("GetHead ERR")
		fmt.Println(c.Request.Header)
		c.Abort()
		return
	}
	claims, err := utils.ParseToken(authHeader)
	if err != nil {
		fmt.Println("Token parse ERR", err)
		c.Abort()
		return
	}
	c.Set("token", claims)
}

func IsAdmin(c *gin.Context) {

	//todo 获取token 鉴权
	authHeader := c.Request.Header.Get("x-token")
	if authHeader == "" {
		fmt.Println("GetHead ERR")
		fmt.Println(c.Request.Header)
		c.Abort()
		return
	}
	claims, err := utils.ParseToken(authHeader)
	if err != nil {
		fmt.Println("Token parse ERR", err)
		c.Abort()
		return
	}
	if claims.Auth != "1" {
		c.Abort()
		return
	}
}

func Login(c *gin.Context) {

	resp := constants.ReCode{}
	u := &dao.AdminUser{}
	err := c.ShouldBindJSON(&u)
	if err != nil {
		fmt.Println("ShouldBindJSON err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}

	adminUser, err := dao.AdminUserLogin(u)
	if err != nil {
		fmt.Println("AdminUserLogin err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}

	if adminUser.Password != u.Password {
		fmt.Println("Password err:")
		resp.Err()
		resp.Message = "password or username is err!"
		c.JSON(http.StatusOK, resp)
		return
	}

	token, err := utils.GenToken(adminUser.Id, u.Username, adminUser.Auth)
	if err != nil {
		fmt.Println("GenToken err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	resp.Data = gin.H{"token": token}
	c.JSON(http.StatusOK, resp)
}

func LoginOut(c *gin.Context) {
	resp := constants.ReCode{}
	resp.Ok()
	c.JSON(http.StatusOK, resp)
}

func UserLogin(c *gin.Context) {
	u := &dao.User{}
	err := c.ShouldBindJSON(&u)
	resp := constants.ReCode{}
	if err != nil {
		fmt.Println("ShouldBindJSON err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	if u.Mobile == "" || u.Password == "" {
		fmt.Println("json value err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	user, err := dao.UserLogin(u)
	if err != nil {
		fmt.Println("select loginuser err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	if user.Password != utils.EncryptPassword(u.Password) {
		resp.Err()
		resp.Message = "用户名或者密码不正确"
		c.JSON(http.StatusOK, resp)
		return
	}
	if user.IsDisabled == 1 {
		resp.Err()
		resp.Message = "封禁"
		c.JSON(http.StatusOK, resp)
		return
	}
	token, err := utils.GenToken(user.Id, user.Nickname, user.Id)
	if err != nil {
		fmt.Println("GenToken err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	resp.Data = gin.H{"token": token}
	c.JSON(http.StatusOK, resp)
}

func Info(c *gin.Context) {
	resp := constants.ReCode{}
	token, exists := c.Get("token")
	if !exists {
		fmt.Println("auth Err:")
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	tokenV := token.(*utils.CustomClaims)
	fmt.Println(tokenV)
	resp = constants.ReCode{
		Code:    20000,
		Success: true,
		Message: "成功",
		Data:    gin.H{"roles": "admin管理面板", "name": tokenV.Username, "avatar": "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"},
	}
	c.JSON(http.StatusOK, resp)
}

func Register(c *gin.Context) {
	u := &dao.User{}
	err := c.ShouldBindJSON(&u)
	resp := constants.ReCode{}
	if err != nil {
		fmt.Println("ShouldBindJSON err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	if u.Mobile == "" || u.Password == "" || u.Nickname == "" {
		fmt.Println("json value err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	user, err := dao.GetUserByMobile(u.Mobile)
	if user != nil {
		fmt.Println("用户已存在")
		resp.Err()
		resp.Message = "用户已存在"
		c.JSON(http.StatusOK, resp)
		return
	}
	u.Password = utils.EncryptPassword(u.Password)
	err = dao.Register(u)
	if err != nil {
		fmt.Println("Register err:", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	c.JSON(http.StatusOK, resp)
}

func GetMemberInfo(c *gin.Context) {

	resp := constants.ReCode{}
	authHeader := c.Request.Header.Get("x-token")
	if authHeader == "" {
		fmt.Println("GetHead ERR")
		fmt.Println(c.Request.Header)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	// 按空格分割
	//parts := strings.SplitN(authHeader, " ", 2)
	//if !(len(parts) == 2 && parts[0] == "Bearer") {
	//	fmt.Println("GetHead ERR")
	//	resp.Err()
	//	c.JSON(http.StatusOK, resp)
	//	return
	//}
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
	resp.Ok()
	resp.Data = gin.H{"userInfo": user}
	c.JSON(http.StatusOK, resp)
}

func IsBuy(c *gin.Context) {

	useridstr := c.Param("userid")
	courseidstr := c.Param("courseid")
	//userid, _ := strconv.Atoi(useridstr)
	//courseid, _ := strconv.Atoi(courseidstr)
	resp := constants.ReCode{}
	isBuyCourse, err := dao.UserIsBuyCourse(useridstr, courseidstr)
	if err != nil {
		fmt.Println("UserIsBuyCourse ERR", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	resp.Data = gin.H{"isbuy": isBuyCourse}
	c.JSON(http.StatusOK, resp)
}
