package pan

import (
	"fmt"
	"net/http"
)

const (
	BAIDU_URL string = "http://www.baidu.com"
	CHECK_URL string = `https://passport.baidu.com/v2/api/?logincheck`
)

type BaiduRequest struct {
	Id string
}

func (br *BaiduRequest) Login(client *http.Client) {
	fmt.Println("start login ...")
	br.requestId(client)
	fmt.Printf("ID is %s \n", br.Id)

}

func (br *BaiduRequest) requestId(client *http.Client) {
	resp, err := client.Get(BAIDU_URL)
	if err != nil {
		fmt.Println(err)
	}
	cookies := resp.Cookies()
	for _, v := range cookies {
		if v.Name == "BAIDUID" {
			br.Id = v.Value
		}
	}
}
