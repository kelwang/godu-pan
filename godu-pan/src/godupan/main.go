package main

import (
	"godupan/ui"
	"godupan/pan"
	"net/http"
)

func main() {
	go func() {
		client := &http.Client{}
		br := new(pan.BaiduRequest)
		br.Login(client)
	}()
	ui.Start()
}
