package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//需求
//測試get post 以及post登入 需要用到testify包
//https://kpat.io/2019/06/testing-with-gin/
//最後面有testing的教學
//https://github.com/gin-gonic/gin
//

//testing
//https://studygolang.com/articles/11836

func TestGinGet(t *testing.T) {
	//初始化
	router := setupServer()
	//recorder(response)
	w := httptest.NewRecorder()
	//request
	req, err := http.NewRequest("GET", "/h/99", nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	router.ServeHTTP(w, req)

	fmt.Println("code:", w.Code)
	fmt.Println("pong:", w.Body.String())

	//assert
}

//兩種post 傳form形式 json形式

func TestGinPost(t *testing.T) {
	//初始化
	router := setupServer()
	//recorder(response)
	w := httptest.NewRecorder()
	//body
	body := strings.NewReader(`{"username":"Hyan","password":"123"}`)
	//request
	req, err := http.NewRequest("POST", "/Login", body)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
		return
	}

	router.ServeHTTP(w, req)

	fmt.Println("code:", w.Code)
	fmt.Println("pong:", w.Body.String())
}
