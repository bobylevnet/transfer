package intrf

type Sqlwhere interface {
	Where(sql string, compare string) Sqlfilter
}
