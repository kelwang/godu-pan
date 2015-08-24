package pan

import(
	"net/http"
)

type BaiduAccount struct {
	username string
	password string
	cookies  []http.Cookie	
}