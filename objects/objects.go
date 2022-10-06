package objects

type DataObject struct {
	ID          string
	Length      int
	Value       interface{}
	TemplateMap map[string]*DataObject
}
