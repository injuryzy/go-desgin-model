package proxy

type NginxServer struct {
	application       *App
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func (n *NginxServer) HandleRequest(url, method string) (int, string) {
	if n.checkRateLimiting(url) {
		return 403, "Not Allowed"
	}
	return n.application.HandleRequest(url, method)
}

func NewNginx() *NginxServer {
	return &NginxServer{
		application:       &App{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (n *NginxServer) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
		return true
	}
	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}
	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}
