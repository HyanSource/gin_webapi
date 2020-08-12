package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

//需求
//測試get post 以及post登入 需要用到testify包
//https://kpat.io/2019/06/testing-with-gin/
//最後面有testing的教學
//https://github.com/gin-gonic/gin

func TestGinGet(t *testing.T) {
	testServer := httptest.NewServer(setupServer())

	defer testServer.Close()

	r, err := http.Get(fmt.Sprintf("%s/h/99", testServer.URL))
	if err != nil {
		t.Fatalf("get err:%v", err)
	}
	if r.StatusCode != 200 {
		t.Fatalf("code err:%v", err)
	}

	fmt.Println("get value:", r.Body.)
}

func TestGinPost(t *testing.T) {

}
