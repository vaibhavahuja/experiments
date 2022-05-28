package service

import (
	"github.com/experimentsGo/entities"
	log "github.com/sirupsen/logrus"
)

func (app *Application) ConvertToRule() *entities.Rule {
	log.Infof("Converting the given entity %v to a rule", app.template)
	rule := app.template.GetRuleFromTemplate()
	log.Infof("Successfully converted. Here are the results %v", rule)

	return rule
}
