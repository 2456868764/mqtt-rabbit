package errcode

var (
	Success        = Code(0, "success")
	Unauthorized   = Code(401, "认证失败")
	NothingFound   = Code(404, "数据不存在")
	Canceled       = Code(498, "客户端取消请求")
	ServerError    = Code(500, "服务器错误")
	ServerDeadline = Code(504, "服务调用超时")

	ParamsError = Code(2001, "参数错误")
)

type Status struct {
	Code int
	Msg  string
}

func Code(code int, msg string) *Status {
	return &Status{
		Code: code,
		Msg:  msg,
	}
}
