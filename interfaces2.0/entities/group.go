package entities

type Group struct {
	RuleGroupList []*RuleGroup
	Age           int
	Description   string
}

func (g *Group) GetListOfNames() (response []string) {
	for i := 0; i < len(g.RuleGroupList); i++ {
		response = append(response, g.RuleGroupList[i].Name)
	}
	return
}
