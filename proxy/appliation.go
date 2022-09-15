package proxy

type Request interface {
	HandleRequest(url, method string) (int, string)
}

type App struct {
}

func (a App) HandleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "Ok"
	}

	if url == "/create/user" && method == "POST" {
		return 201, "User Created"
	}
	return 404, "Not Ok"
}
