import request from '@/utils/request'


export default{
    //讲师列表（条件分页查询）
    //current 分页的当前页 limti每页条数，teacherQuery查询对象
    getTeacherList(current,limit,teacherQuery){
        return request({
            url: `/eduservice/teacher/pageTeacherCondition/${current}/${limit}`,
            method: 'post',
            //teacherQuery 条件对象
            data:teacherQuery
            })
    },
    deleteTeacherId(id){
        return request({
            url: `/eduservice/delteacherid/${id}`,
            method: 'delete',
            })
    },
    addTeacher(teacher){
        return request({
            url: `/eduservice/addteacher`,
            method: 'post',
            data:teacher
            })
    },
    getTeacherInfo(id){
        return request({
            url: `/eduservice/getteacher/${id}`,
            method: 'get',
            })
    },
    //修改讲师
    updateTeacherInfo(teacher){
        return request({
            url: `/eduservice/updateteacher`,
            method: 'post',
            data:teacher
            })
    }
}
