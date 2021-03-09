package models

type Tbl struct {
	Name string
}

type Vtargetmodel struct {
	UserTargetName string `db:target_name`
	TargetName     string `db:target_name`
}


zsacaxzsc