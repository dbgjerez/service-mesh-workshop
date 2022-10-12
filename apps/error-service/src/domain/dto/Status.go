package dto

type Status struct {
	app App
}

type App struct {
	service string
	version string
}
