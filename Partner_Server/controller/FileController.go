package controller

import (
	"Partner_Web/Partner_Server/common"
	_ "bytes"
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// 头像存储目录
const avatarDir = "./avatars"

type FileController struct{}

func NewFileController() *FileController {
	return &FileController{}
}

// UploadAvatar 处理用户头像上传
func (f *FileController) UploadAvatar(c *gin.Context) {
	// 获取用户ID（从请求中获取）
	userID := c.PostForm("userID")
	if userID == "" {
		common.Fail(c, http.StatusBadRequest, nil, "用户ID不能为空")
		return
	}

	// 获取上传的文件
	file, err := c.FormFile("avatar")
	if err != nil {
		common.Fail(c, http.StatusBadRequest, nil, "请上传图片文件")
		return
	}

	// 检查文件类型
	contentType := file.Header.Get("Content-Type")
	if !isImageFile(contentType) {
		common.Fail(c, http.StatusBadRequest, nil, "只允许上传图片文件")
		return
	}

	// 生成文件路径（统一保存为 .jpg 格式）
	filePath := filepath.Join(avatarDir, userID+".jpg")

	// 检查并创建 avatars 文件夹
	if _, err := os.Stat(avatarDir); os.IsNotExist(err) {
		err := os.Mkdir(avatarDir, os.ModePerm)
		if err != nil {
			common.Fail(c, http.StatusInternalServerError, nil, "无法创建 avatars 文件夹")
			return
		}
	}

	// 创建目标文件
	out, err := os.Create(filePath)
	if err != nil {
		common.Fail(c, http.StatusInternalServerError, nil, "无法创建文件")
		return
	}
	defer out.Close()

	// 打开上传的文件
	in, err := file.Open()
	if err != nil {
		common.Fail(c, http.StatusInternalServerError, nil, "无法打开上传的文件")
		return
	}
	defer in.Close()

	// 将上传的文件内容转换为 .jpg 格式并保存
	err = convertToJPEG(in, out)
	if err != nil {
		common.Fail(c, http.StatusInternalServerError, nil, err.Error())
		return
	}

	// 返回成功响应
	common.Success(c, nil, "头像上传成功")
}

// 检查是否为图片文件类型
func isImageFile(contentType string) bool {
	allowedTypes := []string{"image/jpeg", "image/png", "image/gif", "image/webp"}
	for _, t := range allowedTypes {
		if contentType == t {
			return true
		}
	}
	return false
}

func convertToJPEG(in io.Reader, out io.Writer) error {
	// 注册PNG格式
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	// 检查图片格式
	config, format, err := image.DecodeConfig(in)
	if err != nil {
		return fmt.Errorf("无法解码图片配置: %v", err)
	}
	fmt.Printf("图片格式: %s, 尺寸: %dx%d\n", format, config.Width, config.Height)

	// 重新定位到文件开头
	_, err = in.(io.Seeker).Seek(0, io.SeekStart)
	if err != nil {
		return fmt.Errorf("无法重新定位文件: %v", err)
	}

	// 解码图片
	img, _, err := image.Decode(in)
	if err != nil {
		return fmt.Errorf("无法解码图片: %v", err)
	}

	// 编码为 JPEG 格式
	err = jpeg.Encode(out, img, &jpeg.Options{Quality: 90})
	if err != nil {
		return fmt.Errorf("无法编码为JPEG: %v", err)
	}

	return nil
}
