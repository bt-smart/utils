package result

type Response struct {
	Code int    `json:"code"` // 错误码
	Msg  string `json:"msg"`  // 错误描述
	Data any    `json:"data"` // 返回数据
}

func Ok() *Response {
	return &Response{Code: 0, Msg: "", Data: nil}
}

func Fail() *Response {
	return &Response{Code: 1, Msg: "", Data: nil}
}

func FailWithMsg(msg string) *Response {
	return &Response{Code: 1, Msg: msg, Data: nil}
}

func FailWithCodeAndMsg(code int, msg string) *Response {
	return &Response{Code: code, Msg: msg, Data: nil}
}

func Data(data any) *Response {
	return &Response{Code: 0, Msg: "", Data: data}
}
