package utils

type Response struct {
	Data      any    `json:"data"`
	IsSuccess bool   `json:"is_success"`
	Message   string `json:"message"`
}

type ResponseFailureBadRequest struct {
	IsSuccess bool    `json:"is_success" example:"false"`
	Message   string  `json:"message" example:"Có lỗi khi gởi dữ liệu lên server, vui lòng kiểm tra lại!!!"`
	Data      *string `json:"data" example:"null"`
}

type ResponseFailureServerError struct {
	IsSuccess bool    `json:"is_success" example:"false"`
	Message   string  `json:"message" example:"Có lỗi xảy ra tu phia server!!!"`
	Data      *string `json:"data" example:"null"`
}

func SuccessResponseJson(message string, data any) *Response {
	return &Response{
		IsSuccess: true,
		Data:      data,
		Message:   message,
	}
}

func ErrorResponseJson(message string) *Response {
	return &Response{
		IsSuccess: false,
		Data:      nil,
		Message:   message,
	}
}
