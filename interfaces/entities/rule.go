package entities

import log "github.com/sirupsen/logrus"

type Rule struct {
	Uuid string
	Name string
}

func (r *Rule) SendInfo() {
	log.Infof("My Rule is %v and has been generated from", r)
}

func (r *Rule) GetRuleFromTemplate() *Rule {
	rule := &Rule{
		Uuid: "My UUID ",
		Name: "GoodMorning",
	}
	log.Infof("Reacher here mate")
	return rule
}
