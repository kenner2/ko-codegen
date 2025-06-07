package templates

type Template interface {
	SetClassName(string)
	SetClassDesc(string)
	SetParentClass(string)
	AddProperty(def PropertyDefinition)
	AddMethod(def MethodDefinition)
	Generate() (string, error)
	GetFileName() string
}
type PropertyDefinition struct {
	Type        string
	Name        string
	Description string
}

type MethodDefinition struct {
	ReturnType  string
	Name        string
	Params      [][]string
	Body        string
	Description string
}
