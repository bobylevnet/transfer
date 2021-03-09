package intrf

import (
	"transfer/st"
)

//Model -  модели поиска
type Modeln interface {
	Update(modelUpdate st.Abstracttblnew, result *string) bool
	Delete() bool
	Insert(modelInsert st.Abstracttblnew) bool
	Findone(modelFind st.Abstracttblnew) Sqlselect
	Find(modelFind st.Abstracttblnew) Sqlselect
	Findall(modelFind st.Abstracttblnew) Sqlselect
	Upload(model st.Abstracttblnew) Sqlselect
	FindWhere(model st.Abstracttblnew) Sqlselect
}
