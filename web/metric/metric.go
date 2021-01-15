package metric

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"perch/web/model"

	"strconv"
	"time"
)

/**
框架处理函数
@param
@param
@param
@param
@param
@param
*/
func ProcessMetricFunc(w http.ResponseWriter, r *http.Request, bean interface{}, middlePlugin MiddlewarePlugins, f func(ctx context.Context, bean interface{}, respone *model.ResultResponse) error) {
	var (
		response model.ResultResponse
		ctx      context.Context
		err      error
	)
	defer func() {
		if err == nil {
			response.Code = http.StatusOK
		} else {
			if response.Spec == nil {
				response.Spec = err.Error()
			}
			response.Total = 1
		}
		if w != nil {
			err = json.NewEncoder(w).Encode(response)
			if err != nil {
				log.Println(err)
			}
		}

		log.Printf("request url %s with method %s,remote addr is %s\n", r.URL, r.Method, GetRemoteIP(r))
	}()
	now := time.Now()
	timeoutStr := r.Header.Get("Time_Out")
	if timeoutStr == "" {
		deadlineStr := r.Header.Get("Dead_Line")
		if deadlineStr == "" {
			// 无需设置截止时间
			ctx = context.Background()
		} else {
			deadline, err := time.Parse(time.RFC3339, deadlineStr)
			if err != nil {
				err = errors.New("超时时间设置出错")
				return
			}
			fmt.Println("deadline -1 is:", deadline.Unix())
			if deadline.Before(now) {
				err = errors.New("超时时间设置错误")
				return
			}
			ctx, _ = context.WithDeadline(context.Background(), deadline)
		}
	} else {
		timeout, err := strconv.Atoi(timeoutStr)
		if err != nil {
			err = errors.New("超时时间设置错误")
			return
		}
		deadline := now.Add(time.Duration(timeout) * time.Second)
		if deadline.Before(now) {
			err = errors.New("程序处理超时")
			return
		}
		ctx, _ = context.WithDeadline(context.Background(), deadline)
	}
	errChan := make(chan error, 1)
	// 调用回调函数进行业务逻辑的处理
	go func() {
		errChan <- f(ctx, bean, &response)
	}()
	// 等待超时或者回调函数处理完成
	select {
	case <-ctx.Done():
		err = errors.New("函数处理超时")
		break
	case err = <-errChan:
		//fmt.Printf("error is %+v\n",err)
		break
	}
}

/**
中间件插件
*/
type MiddlewarePlugins struct {
	AuthPlugin `json:"auth_plugin"`
}

type AuthPlugin struct {
	AuthToken bool `json:"auth_token"` //token 进行认证
}

func GetRemoteIP(r *http.Request) string {
	clientIp := r.Header.Get("X-Real-Ip")
	if clientIp == "" {
		clientIp = r.Header.Get("X-Forwarded-For")
	}
	if clientIp == "" {
		clientIp = r.RemoteAddr
	}
	return clientIp
}

// 处理请求
