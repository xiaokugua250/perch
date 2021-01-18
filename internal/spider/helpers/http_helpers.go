package helpers

//todo httpheader 设置
func HttpHeadersHelper() (string, error) {
	var (
		userAgent string
		err       error
	)
	return userAgent, err
}

//todo http transport helper 设置
func HttpTransPortHelper() (string, error) {
	/**
	  c.WithTransport(&http.Transport{
	      Proxy: http.ProxyFromEnvironment,
	      DialContext: (&net.Dialer{
	          Timeout:   30 * time.Second,          // 超时时间
	          KeepAlive: 30 * time.Second,          // keepAlive 超时时间
	          DualStack: true,
	      }).DialContext,
	      MaxIdleConns:          100,               // 最大空闲连接数
	      IdleConnTimeout:       90 * time.Second,  // 空闲连接超时
	      TLSHandshakeTimeout:   10 * time.Second,  // TLS 握手超时
	      ExpectContinueTimeout: 1 * time.Second,
	  }
	*/
	return "", nil
}
