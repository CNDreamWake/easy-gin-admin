package vo

import (
	"cmp"
	"fmt"
)

type Err struct {
	code int
	msg  string
}

func NewErr() *Err {
	return &Err{}
}

func (e *Err) Error() string {
	return fmt.Sprintf("ErrCode:%d ErrMsg:%s", e.code, e.msg)
}

func (e *Err) Msg(msg string, code ...int) *Err {
	if len(code) > 0 {
		e.code = cmp.Or(code[0], ERROR)
	}
	e.msg = msg
	return e
}
