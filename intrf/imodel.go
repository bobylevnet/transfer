package intrf

import (
	"transfer/st"
)

//Model -  модели поиска
type Model interface {
	Update(modelUpdate *st.Abstracttbl, result *string) bool
	Delete() bool
	Insert(modelInsert *st.Abstracttbl) int
	Getone(modelFind *st.Abstracttbl, result *string)
	Find(modelFind *st.Abstracttbl, result *string)
	Getall(modelFind *st.Abstracttbl, result *string)
	Gettable() string
}
