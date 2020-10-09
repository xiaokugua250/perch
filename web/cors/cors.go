// package cors
// 跨域支持中间件
// 开发者：天河星光开发小组(钟康游、李江、郭贵鑫、杜量、赵帅帅、曹鹏)
// 邮箱：kangyou.zhong@nscc-gz.cn、jiang.li@nscc-gz.cn、guixin.guo@nscc-gz.cn、liang.du@nscc-gz.cn、shuaishuai.zhao@nscc-gz.cn、peng.cao@nscc-gz.cn
package cors

import "net/http"

// CORS 跨域支持数据结构
type CORS struct {
	h http.Handler
}

// ServeHTTP 自动添加跨域支持信息到返回的头部里面
func (cors *CORS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Headers", "X-Requested-With, Authorization, Content-Type, Cache-Control,x-token, ETag, TIMEOUT, DEADLINE, content-range")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "HEAD, OPTIONS, GET, PUT, POST, PATCH, DELETE")
	// 对于Ajax的OPTIONS请求无需执行真正的处理，直接返回即可
	if r.Method == "OPTIONS" {
		return
	}
	cors.h.ServeHTTP(w, r)
}

// NewHandler 产生新的跨域支持数据结构
func NewHandler(handler http.Handler) *CORS {
	return &CORS{h: handler}
}
