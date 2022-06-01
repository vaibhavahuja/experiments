package entities

type Rule struct {
	Name string
}

func (r *Rule) GetListOfNames() (response []string) {
	response = append(response, r.Name)
	return
}
