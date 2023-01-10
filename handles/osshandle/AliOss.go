package ossHandle

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"goadmin/constants"
	"goadmin/utils"
	"mime/multipart"
	"time"
)

func OssUpload(file *multipart.FileHeader) (string, error) {

	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	client, err := oss.New(constants.Endpoint, constants.AccessKeyId, constants.AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket("your Bucket Name ")
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	// 填写本地文件的完整路径，例如D:\\localpath\\examplefile.txt。
	fd, err := file.Open()
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	defer fd.Close()

	//名称处理操作
	//随机字符串
	getRandstring := utils.GetRandstring()

	//时间戳
	datepath := time.Now().Format("2006/01")

	upname := datepath + "/" + getRandstring + file.Filename
	// 将文件流上传至exampledir目录下的exampleobject.txt文件。
	err = bucket.PutObject(upname, fd)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	//获取文件的url返回
	url := "your oss link" + "/" + upname
	return url, nil
}
