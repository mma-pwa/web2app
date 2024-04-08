package upload

import (
	"errors"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
	"web2app/global"
	"web2app/utils"
)

type Uploads struct{}

func (*Uploads) UploadFile(file *multipart.FileHeader, filePath string) (string, string, error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	if ext != ".png" {
		//return "", "", errors.New("文件后缀名必须是png")
	}
	//// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V([]byte(name))
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	// 尝试创建此路径
	mkdirErr := os.MkdirAll(filePath, os.ModePerm)
	if mkdirErr != nil {
		global.GVA_LOG.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名

	filepath := filePath + "/" + filename
	f, openError := file.Open() // 读取文件
	if openError != nil {
		global.GVA_LOG.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(filepath)
	if createErr != nil {
		global.GVA_LOG.Error("function os.Create() Filed", zap.Any("err", createErr.Error()))

		return "", "", errors.New("function os.Create() Filed, err:" + createErr.Error())
	}
	defer out.Close()

	_, copyErr := io.Copy(out, f)
	if copyErr != nil {
		global.GVA_LOG.Error("function io.Copy() Filed", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return filepath, filename, nil
}

func (*Uploads) DeleteFile(key string) error {
	p := global.GVA_CONFIG.Local.UploadPath + "/" + key
	if strings.Contains(p, global.GVA_CONFIG.Local.UploadPath) {
		if err := os.Remove(p); err != nil {
			return errors.New("本地文件删除失败, err:" + err.Error())
		}
	}
	return nil
}
