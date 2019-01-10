package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"mime/multipart"
	"os"
	"webBlog/helper"
)

type Base struct {
}
// 获取文件大小的接口
type Size interface {
	Size() int64
}

// 获取文件信息的接口
type Stat interface {
	Stat() (os.FileInfo, error)
}

// 构造返回值字段
type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}
func (b *Base) Upload (c *gin.Context){
	var err error
	file, _, err := c.Request.FormFile("file")
	if err == nil {
		var key string
		key, err = uploadFile(file)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"succeed": true,
				"url":     system.GetConfiguration().QiniuFileServer + key,
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"succeed": false,
		"message": err.Error(),
	})
}
func uploadFile(file multipart.File) (string, error) {
	config := helper.GetConfig()
	qiniuAccessKey := config.GetValue("qiniu", "QiniuAccessKey")
	qiniuSecretKey := config.GetValue("qiniu", "QiniuSecretKey")
	qiniuBucket := config.GetValue("qiniu", "QiniuBucket")
	// 创建一个Client
	c := kodo.New(0, nil)
	// 设置上传的策略
	policy := &kodo.PutPolicy{
		Scope: qiniuBucket,
		//设置Token过期时间
		Expires: 3600,
	}
	// 生成一个上传token
	token := c.MakeUptoken(policy)
	// 构建一个uploader
	zone := 0
	uploader := kodocli.NewUploader(zone, nil)

	var size int64
	if statInterface, ok := file.(Stat); ok {
		fileInfo, _ := statInterface.Stat()
		size = fileInfo.Size()
	}
	if sizeInterface, ok := file.(Size); ok {
		size = sizeInterface.Size()
	}

	var ret PutRet
	err := uploader.PutWithoutKey(nil, &ret, token, file, size, nil)
	if err != nil {
		return "", err
	}
	return ret.Key, nil
}
