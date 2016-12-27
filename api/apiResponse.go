package api

type ApiResponse struct {
	Id    string                 `json:"id,omitempty"`
	State State                  `json:"state"`
	Page  *Page                  `json:"page,omitempty"` // default nil for pointer
	User  *User                  `json:"user,omitempty"`
	Data  map[string]interface{} `json:"data,omitempty"`
}

type State struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	Successful = State{0, "请求成功"}
	Error      = State{100001, "服务端错误"}
	JsonError  = State{100002, "请求数据错误"}
)
