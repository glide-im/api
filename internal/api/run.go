package api

// RunHttpServer 启动 http 服务器, 以 HTTP 服务方式访问 api
func RunHttpServer(addr string, port int) error {
	return run(addr, port)
}
