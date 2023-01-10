import request from '@/utils/request'
export default {
  getPlayAuth(vid) {
    return request({
      url: `/eduservice/video/GetVideoAuth/${vid}`,
      method: 'get'
    })
  }

}