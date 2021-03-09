package intrf

//Model -  модели поиска
type Sqlfilter interface {
	And(json string, compare string) Sqlfilter
	Or(json string, compare string) Sqlfilter
	Like(json string, compare string) Sqlfilter
}
