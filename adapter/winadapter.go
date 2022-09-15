package adapter

type WinAdapter struct {
	WindowCache *Win
}

func (w WinAdapter) InsertIntoLightPort() {
	w.WindowCache.InsertIntoUsb()
}
