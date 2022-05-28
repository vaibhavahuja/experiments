package service

import "github.com/experimentsGo/entities"

type ITemplate interface {
	GetRuleFromTemplate() *entities.Rule
}
