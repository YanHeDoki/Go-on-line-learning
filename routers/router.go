package routers

import (
	"github.com/gin-gonic/gin"
	"goadmin/handles"
)

func Router() *gin.Engine {

	g := gin.Default()
	g.Use(handles.Cors())

	//用户操作
	g.POST("/user/login", handles.Login)
	g.POST("/user/logout", handles.LoginOut)

	//前台
	g.POST("/user/front/login", handles.UserLogin)
	g.POST("/user/front/register", handles.Register)
	g.GET("/user/front/getMemberInfo", handles.GetMemberInfo)

	//前台 无需要权限
	g.GET("/educms/front/getBanner", handles.GetAllBanner)
	g.GET("/educms/front/indexfront", handles.IndexFront)
	g.POST("/educms/front/pageTeacher/:page/:limit", handles.PageTeacher)
	g.GET("/eduservice/front/teacherinfo/:id", handles.GetTeacherInfo)

	g.POST("/educms/front/pageCourse/:page/:limit", handles.PageCourse)
	g.GET("/educms/front/GetCourseFront/:courseid", handles.GetCourseFront)

	g.GET("/eduservice/listSubject", handles.CourseList)
	g.Use(handles.SetAuth)
	g.GET("/user/info", handles.Info)

	//教师管理
	g.GET("/eduservice/teacherlist", handles.IsAdmin, handles.TeacherList)
	g.POST("/eduservice/teacher/pageTeacherCondition/:page/:limit", handles.IsAdmin, handles.PageTeacherCondition)
	g.DELETE("/eduservice/delteacherid/:id", handles.IsAdmin, handles.DelTeacherId)
	g.POST("/eduservice/addteacher", handles.IsAdmin, handles.AddTeacher)
	g.GET("/eduservice/getteacher/:id", handles.IsAdmin, handles.GetTeacher)
	g.POST("/eduservice/updateteacher", handles.IsAdmin, handles.UpdateTeacher)
	g.POST("/eduservice/uploadavatar", handles.IsAdmin, handles.UploadAvatar)

	//课程管理
	g.POST("/eduservice/addSubject", handles.Addcourse)
	//g.GET("/eduservice/listSubject", handles.CourseList)
	g.POST("/eduservice/AddCourseInfo", handles.AddCourseInfo)
	g.GET("/eduservice/GetChapterVideo/:courseid", handles.GetChapterVideo)
	g.GET("/eduservice/GetCourseInfo/:courseid", handles.GetCourseInfo)
	g.POST("/eduservice/updateCourseInfo", handles.UpdateCourseInfo)
	g.GET("/eduservice/getPublishCourseInfo/:courseid", handles.GetPublishCourseInfo)
	g.POST("/eduservice/PublishCourse/:courseid", handles.PublishCourseInfo)

	g.POST("/eduservice/addCourseChapter", handles.AddCourseChapter)
	g.GET("/eduservice/CourseChapter/:chapterId", handles.GetCourseChapter)
	g.GET("/eduservice/Course", handles.GetCourse)
	g.POST("/eduservice/updateCourseChapter", handles.UpdateCourseChapter)
	g.DELETE("/eduservice/delCourseChapter/:chapterid", handles.DelCourseChapter)
	g.DELETE("/eduservice/delCourse/:courseid", handles.DelCourse)

	g.POST("/eduservice/video/addvideo", handles.AddVideo)
	g.DELETE("/eduservice/video/delvideo/:id", handles.DelVideo)
	g.POST("/eduservice/video/Upvideo", handles.UpVideo)
	g.DELETE("/eduservice/video/DelAlivideo/:id", handles.DelAliVideo)
	g.DELETE("/eduservice/video/DelAlivideoList", handles.DelAliVideoList)
	g.GET("/eduservice/video/GetVideoAuth/:id", handles.GetVideoAuth)

	g.GET("/educms/admin/pageBanner/:page/:limit", handles.PageGetBanner)
	g.GET("/educms/admin/getBanner/:id", handles.GetBanner)

	g.POST("/educms/admin/addBanner", handles.AddBanner)
	g.PUT("/educms/admin/updateBanner", handles.UpdateBanner)
	g.DELETE("/educms/admin/DelBanner/:id", handles.DelBanner)

	//pay
	g.POST("/pay/CourseOrder/:courseid", handles.CreatCourseOrder)
	g.GET("/pay/GetOrder/:ordid", handles.GetOrdById)
	g.GET("/pay/CreatNative/:orderNo", handles.CreatNative)
	g.GET("/pay/GetOrdStatus/:orderNo", handles.GetOrdStatus)
	g.GET("/isBuy/:userid/:courseid", handles.IsBuy)

	//统计
	g.POST("/countRegister/:day", handles.IsAdmin, handles.CountRegister)
	g.GET("/showData/:type/:start/:end", handles.IsAdmin, handles.ShowData)

	//权限
	g.GET("/admin/acl/permission", handles.GetAllMenu)
	g.DELETE("/admin/acl/remove/:id", handles.DelAllMenu)
	g.POST("/admin/acl/doAssign", handles.DoAssign)
	return g
}
