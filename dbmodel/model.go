package dbmodel

import "transfer/intrf"

type Model struct {
	sql string
}

type Filter struct {
}

type Wheres struct {
}

var sqltxt string

func (m Model) Select(table string) intrf.Sqlwhere {

	sqltxt = "SELECT * FROM " + table

	return Wheres{}

}

func (m Wheres) Where(s string, compare string) intrf.Sqlfilter {

	sqltxt = sqltxt + " WHERE "
	return Filter{}
}

func (f Filter) And(s string, compare string) intrf.Sqlfilter {

	sqltxt = sqltxt + " AND " + gensql(s, compare)
	return f
}

func (f Filter) Or(s string, compare string) intrf.Sqlfilter {

	sqltxt = sqltxt + " OR " + gensql(s, compare)
	return f
}

func (f Filter) Like(s string, compare string) intrf.Sqlfilter {

	//for key, value := range s {
	//	sqltxt = sqltxt + key + " LIKE " + value
	//}

	return f
}

//тут дожен быть json
func gensql(sql string, compare string) string {

	/* 	var txt string
	   	if compare == "" {
	   		compare = "="
	   	}
	   	for key, value := range sql {
	   		txt = key + compare + value
	   	}
	   	return txt */

}
