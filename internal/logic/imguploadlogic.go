package logic

import (
	"context"
	"errors"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"schoolLifeDemo/internal/config"
	"schoolLifeDemo/internal/utils/stringUtils"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"schoolLifeDemo/internal/svc"
)

type ImgUploadLogic struct {
	logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	request  http.Request
	response http.ResponseWriter
}

func NewImgUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext, request *http.Request, response http.ResponseWriter) *ImgUploadLogic {
	return &ImgUploadLogic{
		Logger:   logx.WithContext(ctx),
		ctx:      ctx,
		svcCtx:   svcCtx,
		request:  *request,
		response: response,
	}
}

func (l *ImgUploadLogic) ImgUpload() (filePath string, err error) {
	// 1、文件最大值判断，5M
	r := l.request
	w := l.response

	r.Body = http.MaxBytesReader(w, r.Body, config.MaxFileSize)
	if err := r.ParseMultipartForm(config.MaxFileSize); err != nil {
		return "", err
	}
	// 2、上传文件类型校验
	//fileType := r.PostFormValue("type")
	file, _, err := r.FormFile("uploadFile")
	if err != nil {
		return "", err
	}
	defer file.Close()
	// 读取文件内容
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	// 判断文件类型
	fileType := http.DetectContentType(fileBytes)
	if stringUtils.IsContain(config.ImgFileType, fileType) {
		// 创建文件路径和名称 用户名称/wallType/随机生成的UUID.jpg
		fileName := uuid.NewV4()
		fileEndings, err := mime.ExtensionsByType(fileType)
		if err != nil {
			return "", err
		}
		newPath := filepath.Join(config.ImgFilePath, fileName.String()+fileEndings[0])
		// 创建文件，并写入
		// 判断目录是否存在，不存在则创建
		if _, err := os.Stat(config.ImgFilePath); err != nil {
			os.MkdirAll(config.ImgFilePath, 0777)
		}
		err = ioutil.WriteFile(newPath, fileBytes, 0666)
		if err != nil {
			return "", err
		}
		return newPath, nil
	} else {
		return "", errors.New("文件类型不再规定范围内，要求文件类型为: " + strings.Join(config.ImgFileType, "|"))
	}
}
