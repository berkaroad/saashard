package sqlparser

import "testing"

func TestParseSelect(t *testing.T) {
	sql := "select 1 as `f2`, 'hello' `f3`, `t1`.`f1`, 'world', 7 from `table1` as `t1` left join table2 on table2.f1=t1.f3 where 1>0"
	sql = "select f1,f2,f3 into t2 from t1" // Unsupported
	sql = "insert into t2 select f1,f2,f3 from t1 where t1.a3=8"
	sql = "select (case t1.f1 when '0' then 'hello' else 'world' end) f1 from t1"
	sql = "select (case when t1.f1='0' then 'hello' else 'world' end) f1 from t1"
	if stmt, err := Parse(sql); err != nil {
		t.Error(err)
	} else {
		t.Log(String(stmt))
		switch v := stmt.(type) {
		case *Select:

		case *Insert:

		case *Update:
			t.Log("Update", v)
		case *Delete:
			t.Log("Delete", v)
		case *Replace:
			t.Log("Replace", v)
		case *Set:
			t.Log("Set", v)
		case *Begin:
			t.Log("Begin", v)
		case *Commit:
			t.Log("Commit", v)
		case *Rollback:
			t.Log("Rollback", v)
		case *Admin:
			t.Log("Admin", v)
		case *AdminHelp:
			t.Log("AdminHelp", v)
		case *UseDB:
			t.Log("UseDB", v)
		case *SimpleSelect:
			t.Log("SimpleSelect", v)
		default:
			t.Log(" not support now")
		}
		t.Fail()
	}

}
