package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

type FormObj struct {

}


func (fo *FormObj) AddRoutes() {
	http.HandleFunc("/form", fo.HandleForm)
	http.HandleFunc("/upload/file", fo.HandleUploadFile)
}


func (fo *FormObj) HandleForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	for key, val := range r.Form {
		fmt.Printf("%s: %s\n", key, val)
		fmt.Fprintf(w, "%s: %s\n", key, val)
	}
}


func (fo *FormObj) HandleUploadFile(w http.ResponseWriter, r *http.Request) {
	//设置可存放memory的数据量(可选，直接调用FormFile默认为32M)
	if err := r.ParseMultipartForm(MAX_FILE_SIZE); err != nil {
		fmt.Fprintf(w, "parse multipartform failed: %s\n", err)
	}

	// 根据key获取上传的第一个文件
	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, "form file failed: %s\n", err)
	}
	defer file.Close()

	//判断是不是txt文件
	if ext := path.Ext(header.Filename); ext != ".txt" {
		fmt.Fprintf(w, "file ext failed: expected .txt, actual: %s\n", ext)
	}

	// 返回文件meta信息
	fmt.Fprintf(w, "filename: %s, size: %d, header: %v\n", header.Filename, header.Size, header.Header)

	// 读文件
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintf(w, "read file failed:%s\n", err)
	}
	fmt.Fprintf(w, "file size: %d, content: %s\n", len(content), string(content))

	//读form value
	normal := r.FormValue("normal")
	fmt.Fprintf(w, "normal field value : %s\n", normal)
}
