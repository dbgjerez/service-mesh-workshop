package dto

type Info struct {
	App App `json:"app"`
}

type App struct {
	Service   string `json:"name"`
	Version   string `json:"version"`
	BuildTime string `json:"buildTime"`
}
