package web

type Result struct {
	Code    int32       `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Success(data interface{}) *Result {
	return &Result{
		Code:    0,
		Data:    data,
		Message: "success",
	}
}

func Error(code int32, message string) *Result {
	return &Result{
		Code:    code,
		Data:    nil,
		Message: message,
	}
}
