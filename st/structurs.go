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

/*type Formfiles struct {
	Value map[string][]string
	File  map[string][]*FileHeader
}*/

//Upload - структура для загрузки
type Upload struct {
	NameFile   string `db:"name_file"`
	DateCreate string `db:"date_create"`
	DateDelete string `db:"date_delete"`
	PathFile   string `db:"path_file"`
	IDUser     int    `db:"id_user"`
	Idtarget   int    `db:"id_target"`
}

//Auhthb - структура возврата авторизации bitrix
type Auhthb struct {
	AccessToken string
	UserID      int
	ClientID    int
	Expires     int
	Scope       string
	UserName    string
}
