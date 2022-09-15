package adapter

import "testing"

func TestName(t *testing.T) {
	mac := Mac{}
	mac.InsertIntoLightPort()
	win := &Win{}
	adapter := WinAdapter{
		WindowCache: win,
	}
	adapter.InsertIntoLightPort()
}
