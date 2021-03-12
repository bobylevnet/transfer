package intrf

//Model -  модели поиска
type Models interface {
	ConvertJS(interface{}) string
}
