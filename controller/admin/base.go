package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webBlog/helper"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"golang.org/x/net/context"
	"os"
)

type Base struct {
}
// 自定义返回值结构体
type MyPutRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
	Name   string
}
func (b *Base) Upload (c *gin.Context){
	var err error
	file, err := c.FormFile("editormd-image-file")
	if err == nil {
		var key string
		var fileName = "static/uploads/"+file.Filename
		err := c.SaveUploadedFile(file, fileName)
		key, err = uploadFile(fileName)
		config := helper.GetConfig()
		if err == nil {
			//删除文件
			os.Remove(fileName)
			c.JSON(http.StatusOK, gin.H{
				"success": 1,
				"url": config.GetValue("qiniu", "QiniuFileServer") + key ,
				"message":"上传成功",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"success": 0,
		"message": err.Error(),
	})
}
func uploadFile(file string) (string, error) {
	config := helper.GetConfig()
	qiniuAccessKey := config.GetValue("qiniu", "QiniuAccessKey")
	qiniuSecretKey := config.GetValue("qiniu", "QiniuSecretKey")
	qiniuBucket := config.GetValue("qiniu", "QiniuBucket")
	// 创建一个Client
	mac := qbox.NewMac(qiniuAccessKey, qiniuSecretKey)
	// 设置上传的策略
	putPolicy := storage.PutPolicy{
		Scope: qiniuBucket,
		Expires:7200,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	// 生成一个上传token
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := MyPutRet{}
	err := formUploader.PutFileWithoutKey(context.Background(), &ret, upToken, file, nil)
	if err != nil {
		return "", err
	}
	return ret.Key, nil
}
