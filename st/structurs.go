package st

//Event - структура событий
type Event struct {
	Model  string
	Action string
}

// Abstracttbl - струкутра данных табблицы
type Abstracttbl struct {
	Idpk int
	Dt   map[string]string
}

//Metadatatable метаданные таблицы
type Metadatatable struct {
	Columnname string `db:"column_name"`
	Typecolumn string `db:"data_type"`
}
