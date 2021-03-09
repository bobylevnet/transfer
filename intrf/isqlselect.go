package intrf

type Sqlselect interface {
	Select(table string, compare string) Sqlwhere
}
