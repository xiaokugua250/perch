/**
代理helper方法
*/
package helpers

//todo 代理池
func ProxyPoolsHelper() (string, error) {
	var (
		userAgent string
		err       error
	)
	return userAgent, err
}

//todo 代理切换方法
func ProxySwitcherHelper() (string, error) {
	/*
		if p, err := proxy.RoundRobinProxySwitcher(
			"socks5://127.0.0.1:1337",
			"socks5://127.0.0.1:1338",
			"http://127.0.0.1:8080",
		); err == nil {
			c.SetProxyFunc(p)
		}
	*/
	return "", nil
}
