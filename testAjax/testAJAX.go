package main

import (
	"bytes"
	"errors"
	"fmt"
	_ "github.com/astaxie/beego"
	"html/template"
	"io"
	"net/http"
	"runtime"
)

// 端口
const (
	HTTP_PORT  string = "8006"
	HTTPS_PORT string = "443"
)

// 目录
const (
	CSS_CLIENT_PATH   = "/css/"
	DART_CLIENT_PATH  = "/js/"
	IMAGE_CLIENT_PATH = "/image/"

	CSS_SVR_PATH   = "testAjax/web"
	DART_SVR_PATH  = "testAjax/web"
	IMAGE_SVR_PATH = "web"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// 先把css和脚本服务上去
	//http.Handle("/a",http.FileServer(http.Dir("/a")))
	http.Handle(CSS_CLIENT_PATH, http.FileServer(http.Dir(CSS_SVR_PATH)))
	http.Handle(DART_CLIENT_PATH, http.FileServer(http.Dir(DART_SVR_PATH)))

	// 网址与处理逻辑对应起来
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/ajax", OnAjax)

	// 开始服务
	err := http.ListenAndServe(":"+HTTP_PORT, nil)
	if err != nil {
		fmt.Println("服务失败 /// ", err)
	}
}

func WriteTemplateToHttpResponse(res http.ResponseWriter, t *template.Template) error {
	if t == nil || res == nil {
		return errors.New("WriteTemplateToHttpResponse: t must not be nil.")
	}
	var buf bytes.Buffer
	err := t.Execute(&buf, nil)
	if err != nil {
		return err
	}
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err = res.Write(buf.Bytes())
	return err
}

func HomePage(res http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("testAjax/web/loli.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = WriteTemplateToHttpResponse(res, t)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func OnAjax(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "这是从后台发送的数据")
	t,err := template.ParseFiles("testAjax/web/loli.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = t.Execute(res,nil)

}