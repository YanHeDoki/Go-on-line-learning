# Go-on-line-learning



## 这是一个go实现的在线学习网页的后端

##### goadmin  文件夹为go的后端程序

后端部分主要组件使用：

- Gin 作为web框架 
- 数据库组件为sqlx
- Oss模块为阿里云oss
- 视频点播也为阿里的视频点播接口
- 鉴权使用Token 为  jwt-go



##### vuePage为前端页面，使用Vue 构建来分为前后台，使用的公开的资料如下

前端 和数据库来源：尚硅谷_谷粒学苑-微服务+全栈在线教育实战项目

https://www.bilibili.com/video/BV1dQ4y1A75e?p=1&vd_source=b29c4817c351d0a9330506cdd3f8f007



## 启动

##### 步骤

- mysql 8.0 执行数据库脚本中的文件

- 在constants模块中的 RC.go 里配置Oss配置 如下

- ```go
  Endpoint        = "your oss link "
  AccessKeyId     = "your AccessKeyId"
  AccessKeySecret = "your AccessKeySecret"
  ```

- VuePage 下的两个Vue模具分别 安装依赖我使用的node版本为node-v10.14.1-x64：

- ```bash
  npm install
  npm run dev
  ```

- Run goadmin中的main.go





