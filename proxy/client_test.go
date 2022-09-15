package proxy

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	nginx := NewNginx()
	fmt.Println(nginx.HandleRequest("/app/status", "GET"))
	fmt.Println(nginx.HandleRequest("/app/status", "GET"))
	fmt.Println(nginx.HandleRequest("/app/status", "GET"))
	fmt.Println(nginx.HandleRequest("/app/status", "GET"))
	fmt.Println(nginx.HandleRequest("/app/status", "GET"))
}
