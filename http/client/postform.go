package client

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
)

type FormObj struct {

}

func (fo *FormObj) SendForm(urlPath string) (*http.Response, error) {
	form := url.Values{
		"ak": {"av1", "av2", "av3"},
	}
	form.Set("bk", "bv")

	return http.DefaultClient.PostForm(urlPath, form)
}


func (fo *FormObj) UploadFile(filePath, urlPath string) (*http.Response, error) {
	// 打卡要上传的文件
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// 创建构建body的buffer
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	// 填充form的file字段
	fileField, err := writer.CreateFormFile("file", path.Base(filePath))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fileField, f)
	if err != nil {
		return nil, err
	}

	// 填充form的field字段(可选)
	normalField, err := writer.CreateFormField("normal")
	if err != nil {
		return nil, err
	}

	normalField.Write([]byte("normal value"))

	// 关闭multipart writer，这一步会填充ending boundary
	writer.Close()

	// 创建自定义的http.Request
	req, err := http.NewRequest("POST", urlPath, &body)
	if err != nil {
		return nil, err
	}

	// 设置传递类型
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return http.DefaultClient.Do(req)
}
