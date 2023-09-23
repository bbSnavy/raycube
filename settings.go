package main

type Settings struct {
	GameWindowWidth  int32
	GameWindowHeight int32
	GameWindowName   string
}

func (settings *Settings) Init() *Settings {
	settings.GameWindowWidth = 1280
	settings.GameWindowHeight = 720
	settings.GameWindowName = "raycube"

	return settings
}
