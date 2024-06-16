package upload

import (
	"fmt"
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
	"log"
	"path"
)

type UploadController struct {
}

func (UploadController) Unifile(c *gin.Context) {
	file, err := c.FormFile("file") //取的是form Data的数据
	if err == nil {
		fmt.Println(file.Filename)
		//判断文件后缀 类型
		ext := path.Ext(file.Filename)
		allowExtMap := map[string]bool{
			".jpg":  true,
			".jpeg": true,
			".png":  true,
			".gif":  true,
		}
		if _, ok := allowExtMap[ext]; !ok {
			c.String(200, "上传文件不合法")
			return
		}
		//根据日期创建文件夹保存文件
		day := utils.GetDay()
		dst := "./static/" + day + "/" + file.Filename
		c.SaveUploadedFile(file, dst)
		// 构建文件访问URL
		fileURL := fmt.Sprintf("http://%s/static/%s/%s", c.Request.Host, day, file.Filename)
		c.JSON(200, gin.H{
			"message": "Upload",
			"url":     c.Request.RequestURI,
			"file":    fileURL,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Upload失败",
		})
	}
}

func (UploadController) Multifile(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(200, gin.H{
			"message": err.Error(),
		})
	}
	files := form.File["upload_list"] //upload_list[]是参数名称

	fileURLMap := map[string]string{}
	for _, file := range files {
		log.Println(file.Filename)

		// 上传文件至指定目录
		day := utils.GetDay()
		dst := "./static/" + day + "/" + file.Filename
		c.SaveUploadedFile(file, dst)
		// 构建文件访问URL
		fileURL := fmt.Sprintf("http://%s/static/%s/%s", c.Request.Host, day, file.Filename)
		fileURLMap[file.Filename] = fileURL
	}
	c.JSON(200, gin.H{
		"message": "MultifileUpload",
		"file":    fileURLMap,
	})
}
