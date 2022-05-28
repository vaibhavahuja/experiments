package entities

type Template struct {
	uuid string
	name string
}

func GetNewTemplate(uuid, name string) *Template {
	return &Template{
		uuid: uuid,
		name: name,
	}
}

func (t *Template) GetRuleFromTemplate() *Rule {
	rule := &Rule{
		Uuid: "My UUID " + t.uuid,
		Name: "GoodMorning " + t.name,
	}
	return rule
}
