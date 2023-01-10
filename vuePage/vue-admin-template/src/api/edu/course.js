import request from '@/utils/request'
export default {
    //1 添加课程信息
    addCourseInfo(courseInfo) {
        return request({
            url: '/eduservice/AddCourseInfo',
            method: 'post',
            data:courseInfo
          })
    },
    //2 查询所有讲师
    getListTeacher() {
        return request({
            url: '/eduservice/teacherlist',
            method: 'get'
          })
    },
    //根据课程id查询课程基本信息
    getCourseInfoId(id) {
        return request({
            url: '/eduservice/GetCourseInfo/'+id,
            method: 'get'
          })
    },
    //修改课程信息
    updateCourseInfo(courseInfo) {
        return request({
            url: '/eduservice/updateCourseInfo',
            method: 'post',
            data: courseInfo
          })
    },
    //课程确认信息显示
    getPublihCourseInfo(id) {
        return request({
            url: '/eduservice/getPublishCourseInfo/'+id,
            method: 'get'
          })
    }, //课程最终发布
    publihCourse(id) {
        return request({
            url: '/eduservice/PublishCourse/'+id,
            method: 'post'
          })
    },
    //TODO 课程列表
    //课程最终发布
    getListCourse() {
        return request({
            url: '/eduservice/Course',
            method: 'get'
          })
    },
    //删除章节
    deleteCourse(courseId) {
        return request({
            url: '/eduservice/delCourse/'+courseId,
            method: 'delete'
          })
    },

}
