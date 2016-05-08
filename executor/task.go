package executor

type Builder struct {
	Name string
	Tag  string
}

type Task struct {
	Name    string
	Builder Builder
	Vars    VarCollection
}

type Var struct {
	Name  string
	Value string
}

func (v Var) String() string {
	return v.Name + "=" + v.Value
}

type VarCollection []Var

func (c VarCollection) String() []string {
	var ret []string
	for _, v := range c {
		ret = append(ret, v.String())
	}
	return ret
}
