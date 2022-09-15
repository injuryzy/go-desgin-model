package adapter

import "fmt"

type Computer interface {
	InsertIntoLightPort()
}
type Mac struct {
}

func (m *Mac) InsertIntoLightPort() {
	fmt.Println("mac")
}

type Win struct {
}

func (w *Win) InsertIntoUsb() {
	fmt.Println("win")
}
