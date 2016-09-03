package sqlparser

import "testing"

func TestParseSelect(t *testing.T) {
	var sql string
	sql = "SET GLOBAL TRANSACTION ISOLATION LEVEL READ COMMITTED"
	sql = "set charset utf8"
	sql = "set CHARACTER SET utf32"

	sql = "SHOW cHarset"
	sql = "SHOW CHARACTER SET"
	sql = "SHOW CHARACTER SET where Charset = 'utf8'"
	sql = "SHOW SESSION VARIABLES LIKE 'lower_case_table_names'"
	sql = "show session status"
	sql = "SHOW Databases"
	sql = "show tables from `t1`"
	sql = "show full tables from `t1`.`t2`"
	sql = "show columns from t1"
	sql = "show full columns from t1"
	sql = "show procedure status"
	sql = "show function status"

	// sql = "select version()"
	// sql = "select version() limit 0,100"
	// sql = "select 1 as `f2`, 'hello' `f3`, `t1`.`f1`, 'world', 7 from `table1` as `t1` left join table2 on table2.f1=t1.f3 where 1>0"
	// sql = "select f1,f2,f3 into t2 from t1" // Unsupported
	// sql = "insert into t2 select f1,f2,f3 from t1 where t1.a3=8"
	// sql = "select (case t1.f1 when '0' then 'hello' else 'world' end) f1 from t1"
	// sql = "select (case when t1.f1='0' then 'hello' else 'world' end) f1 from t1"

	t.Log(sql)

	if stmt, err := Parse(sql); err != nil {
		t.Error(err)
	} else {
		t.Log(String(stmt))
		t.Fail()
	}

}
