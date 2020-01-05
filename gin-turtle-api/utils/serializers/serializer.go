package serializers

import "github.com/gin-gonic/gin"

const (
	// CodeCheckLogin 未登录
	CodeCheckLogin = 401
	// CodeNoRightErr 未授权访问
	CodeNoRightErr = 403
	// CodeDBError 数据库操作失败
	CodeDBError = 50001
	// CodeEncryptError 加密失败
	CodeEncryptError = 50002
	//CodeParamErr 各种奇奇怪怪的参数错误
	CodeParamErr = 40001
)

// Response serializer
type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"msg"`
	Error   string      `json:"error,omitempty"`
}

// DataList base list construct
type DataList struct {
	Items interface{} `json:"items"`
	Total int         `json:"total"`
}

// TraceErrorResponse traced error response
type TraceErrorResponse struct {
	Response
	TraceID string `json:"trace_id"`
}

// CheckLogin function for check user login status
func CheckLogin() Response {
	return Response{
		Code:    CodeCheckLogin,
		Message: "未登录",
	}
}

// ErrorHandler general error hanlder
func ErrorHandler(erroCode int, message string, err error) Response {
	result := Response{
		Code:    erroCode,
		Message: message,
	}
	if err != nil && gin.Mode() != gin.ReleaseMode {
		result.Error = err.Error()

	}
	return result
}

// DatabaseError database error handler
func DatabaseError(message string, err error) Response {
	if message == "" {
		message = "handle database error"
	}

	return ErrorHandler(CodeDBError, message, err)
}

// ParameterError paramters error handler
func ParameterError(message string, err error) Response {
	if message == "" {
		message = "parameter error"
	}
	return ErrorHandler(CodeParamErr, message, err)
}

// BuildListResponse list constructor generator
func BuildListResponse(items interface{}, total int) Response {
	return Response{
		Data: DataList{
			Items: items,
			Total: total,
		},
	}
}
