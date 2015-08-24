package pan

import ()

const (
	USER_AGENT string = `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/37.0.2062.120 Safari/537.36`
)

type Session string

type Pan struct {
	Uid     uint
	Uss     string
	SToken  string
	Pcsett  string
	Session Session
}
