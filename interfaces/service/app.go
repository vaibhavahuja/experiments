package service

type Application struct {
	template ITemplate
}

func GetNewApplication(t ITemplate) *Application {
	return &Application{template: t}
}
