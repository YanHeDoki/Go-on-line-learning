package ossHandle

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vod"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"goadmin/constants"
	"io"
	"os"
	"strings"
)

func InitVodClient() (client *vod.Client, err error) {

	// 点播服务接入地域
	regionId := "cn-shanghai"

	// 创建授权对象
	credential := &credentials.AccessKeyCredential{
		constants.AccessKeyId,
		constants.AccessKeySecret,
	}

	// 自定义config
	config := sdk.NewConfig()
	config.AutoRetry = true     // 失败是否自动重试
	config.MaxRetryTime = 3     // 最大重试次数
	config.Timeout = 3000000000 // 连接超时，单位：纳秒；默认为3秒

	// 创建vodClient实例
	return vod.NewClientWithOptions(regionId, config, credential)
}

func GetPlayInfo(videoId string) {
	client, err := InitVodClient()
	if err != nil {
		// 异常处理
		panic(err)
	}
	request := vod.CreateGetPlayInfoRequest()
	request.VideoId = videoId
	request.AcceptFormat = "JSON"

	// 发起请求并处理异常，调用client.${apiName}(request)
	getPlayInfo, err := client.GetPlayInfo(request)
	if err != nil {
		// 异常处理
		panic(err)
	}

	playList := getPlayInfo.PlayInfoList.PlayInfo
	for _, playInfo := range playList {
		// 打印清晰度和对应的播放地址
		fmt.Printf("%s: %s\n", playInfo.Definition, playInfo.PlayURL)
	}
}

func GetPlayAuth(videoId string) (string, error) {
	client, err := InitVodClient()
	if err != nil {
		return "", err
	}
	request := vod.CreateGetVideoPlayAuthRequest()
	request.VideoId = videoId
	request.AcceptFormat = "JSON"
	response, err := client.GetVideoPlayAuth(request)
	if err != nil {
		return "", err
	}
	return response.PlayAuth, nil
}

type UploadAuthDTO struct {
	AccessKeyId     string
	AccessKeySecret string
	SecurityToken   string
}
type UploadAddressDTO struct {
	Endpoint string
	Bucket   string
	FileName string
}

func InitOssClient(uploadAuthDTO UploadAuthDTO, uploadAddressDTO UploadAddressDTO) (*oss.Client, error) {
	client, err := oss.New(uploadAddressDTO.Endpoint,
		uploadAuthDTO.AccessKeyId,
		uploadAuthDTO.AccessKeySecret,
		oss.SecurityToken(uploadAuthDTO.SecurityToken),
		oss.Timeout(86400*7, 86400*7))
	return client, err
}

func MyCreateUploadVideo(Title, Description, FileName, CoverURL string, client *vod.Client) (response *vod.CreateUploadVideoResponse, err error) {
	request := vod.CreateCreateUploadVideoRequest()
	request.Title = Title
	request.Description = Description
	request.FileName = FileName
	//request.CateId = "-1"
	//Cover URL示例：http://example.alicdn.com/tps/TB1qnJ1PVXXXXXCXXXXXXXXXXXX-700-****.png
	request.CoverURL = CoverURL
	request.Tags = "test"
	request.AcceptFormat = "JSON"
	return client.CreateUploadVideo(request)
}

func UploadLocalFile(client *oss.Client, uploadAddressDTO UploadAddressDTO, fileStream io.Reader) {
	// 获取存储空间。
	bucket, err := client.Bucket(uploadAddressDTO.Bucket)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 上传本地文件。
	err = bucket.PutObject(uploadAddressDTO.FileName, fileStream)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

func VideoFileUp(Title, Description, FileName, CoverURL string, fileStream io.Reader) (string, error) {

	// 初始化VOD客户端并获取上传地址和凭证
	vodClient, err := InitVodClient()
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	// 获取上传地址和凭证
	response, err := MyCreateUploadVideo(Title, Description, FileName, CoverURL, vodClient)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	// 执行成功会返回VideoId、UploadAddress和UploadAuth
	var videoId = response.VideoId
	var uploadAuthDTO UploadAuthDTO
	var uploadAddressDTO UploadAddressDTO
	var uploadAuthDecode, _ = base64.StdEncoding.DecodeString(response.UploadAuth)
	var uploadAddressDecode, _ = base64.StdEncoding.DecodeString(response.UploadAddress)
	json.Unmarshal(uploadAuthDecode, &uploadAuthDTO)
	json.Unmarshal(uploadAddressDecode, &uploadAddressDTO)
	// 使用UploadAuth和UploadAddress初始化OSS客户端
	var ossClient, _ = InitOssClient(uploadAuthDTO, uploadAddressDTO)
	// 上传文件，注意是同步上传会阻塞等待，耗时与文件大小和网络上行带宽有关
	UploadLocalFile(ossClient, uploadAddressDTO, fileStream)
	//MultipartUploadFile(ossClient, uploadAddressDTO, localFile)
	return videoId, nil
}

func DelAliVideo(id string) error {
	client, err := InitVodClient()
	if err != nil {
		return err
	}
	request := vod.CreateDeleteVideoRequest()
	// 支持批量删除视频，多个用逗号分隔
	request.VideoIds = id
	request.AcceptFormat = "JSON"
	_, err = client.DeleteVideo(request)
	if err != nil {
		return err
	}

	return nil
}

func DelAliVideoList(ids []string) error {
	client, err := InitVodClient()
	if err != nil {
		return err
	}
	request := vod.CreateDeleteVideoRequest()
	// 支持批量删除视频，多个用逗号分隔
	idsstr := strings.Join(ids, ",")
	request.VideoIds = idsstr
	request.AcceptFormat = "JSON"
	_, err = client.DeleteVideo(request)
	if err != nil {
		return err
	}

	return nil
}
