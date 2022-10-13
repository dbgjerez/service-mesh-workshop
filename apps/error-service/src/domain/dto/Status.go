package dto

type Status struct {
	App App `json:"app"`
}

type App struct {
	Service string `json:"name"`
	Version string `json:"version"`
}
