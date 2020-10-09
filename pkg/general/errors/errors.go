package errors

/**
包含其他信息的Error 结构
*/
type ComplexError struct {
	ErrMap   map[string]interface{} `json:"err_map"`
	ErrMsg   string                 `json:"err_msg"`
	ErrorRaw error                  `json:"error_raw"`
}
