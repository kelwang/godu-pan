package main

import (
	"fmt"
	"net/http"
	"godupan/pan"
)


func main() {
	fmt.Println("ok")

	client := &http.Client{}
	br := new(pan.BaiduRequest)
	br.Login(client)
}