package metric

import (
	"context"
	"encoding/json"
	"errors"
	"expvar"
	"fmt"
	"log"
	"net/http"
	"perch/web/model"
	"strconv"
	"time"
)

var (
	ServerMertics = expvar.NewInt("serverMetrics")
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
func ProcessMetricFunc(w http.ResponseWriter, r *http.Request, bean interface{}, f func(ctx context.Context, bean interface{}, respone *model.ResultReponse) error) {
	var (
		response model.ResultReponse
		ctx      context.Context
		err      error
	)
	defer func() {
		if w != nil {
			err = json.NewEncoder(w).Encode(response)
			if err != nil {
				log.Println(err)
			}
		}
	}()
	now := time.Now()
	timeoutStr := r.Header.Get("TIMEOUT")
	if timeoutStr == "" {
		deadlineStr := r.Header.Get("DEADLINE")
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
		break
	}
}
func MerticFunc(ctx context.Context, result model.ResultReponse, err error) http.HandlerFunc {

	//log.Println(r.RequestURI)
	// Call the next handler, which can be another middleware in the chain, or the final handler.

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		start := time.Now()
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			r.Response,
			time.Since(start),
		)

	})
}

// 处理请求

/**
通过expvar 发送服务端监控情况
*/
func MerticsServerFunc(w http.ResponseWriter, r *http.Request) {

	ServerMertics.Add(1)

}
