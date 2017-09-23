package sqlparser

import "testing"

func TestParseSelect(t *testing.T) {
	var sql string
	sql = "SET GLOBAL TRANSACTION ISOLATION LEVEL READ COMMITTED"
	// sql = "set charset utf8"
	// sql = "set CHARACTER SET utf32"

	// sql = "SHOW cHarset"
	// sql = "SHOW CHARACTER SET"
	// sql = "SHOW CHARACTER SET where Charset = 'utf8'"
	// sql = "SHOW SESSION VARIABLES LIKE 'lower_case_table_names'"
	// sql = "show session status"
	// sql = "SHOW Databases"
	// sql = "show tables from `t1`"
	// sql = "show full tables from `t1`.`t2`"
	// sql = "show columns from t1"
	// sql = "show full columns from t1"
	// sql = "show procedure status"
	// sql = "show function status"
	// sql = "SHOW TRIGGERS FROM `kgw2` WHERE `Table` = 'buy_list'"

	// sql = "select version()"
	// sql = "select version() limit 0,100"
	// sql = "select 1 as `f2`, 'hello' `f3`, `t1`.`f1`, 'world', 7 from `table1` as `t1` left join table2 on table2.f1=t1.f3 where 1>0"
	// sql = "select f1,f2,f3 into t2 from t1" // Unsupported
	// sql = "insert into t2 select f1,f2,f3 from t1 where t1.a3=8"
	// sql = "select (case t1.f1 when '0' then 'hello' else 'world' end) f1 from t1"
	// sql = "select (case when t1.f1='0' then 'hello' else 'world' end) f1 from t1"

	// sql = "UPDATE `buy_list` SET `status_id`=5 WHERE `buy_id`=1"

	// sql = "explain select * from `kgw2`.`buy_list`"
	sql = "select version()"
	sql = "select database()"
	sql = "select * from info_goods where POSITION('1000001' in value_text)=1"
	sql = "select * from info_goods where locate('1000001', value_text)=1"
	// sql = "show slave status"

	sql = `CREATE TABLE abc.sys_admin_user_permission (
  user_permission_id int(11) NOT NULL AUTO_INCREMENT,
  client_id int(11) NOT NULL COMMENT '客户id',
  user_id int(11) NOT NULL COMMENT '用户id',
  permission_id int(11) NOT NULL COMMENT '对应sys_permission_page的permission_id',
  permission_value int(11) NOT NULL COMMENT '权限值',
  operator_id int(11) NOT NULL DEFAULT '0' COMMENT '创建人',
  modifier_id int(11) NOT NULL DEFAULT '0' COMMENT '修改人',
  data_created datetime NOT NULL COMMENT '创建时间',
  data_lasted datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (user_permission_id),
  KEY client_id (client_id),
  KEY user_id (user_id),
  CONSTRAINT fk_sys_admin_user_permission_client_id FOREIGN KEY (client_id) REFERENCES sys_admin_client (client_id),
  CONSTRAINT fk_sys_admin_user_permission_user_id FOREIGN KEY (user_id) REFERENCES sys_admin_user (user_id)
) ENGINE=InnoDB AUTO_INCREMENT=32683 DEFAULT CHARSET=utf8 COMMENT='用户权限值表'`

	// 	sql = "create index XXX   USING btree on     table1(col1(5) ASC,col2 DESC)"

	// 	sql = `ALTER TABLE sub_category
	// ADD COLUMN sub_categorycolSS VARCHAR(45) NULL AFTER description`
	sql = "select count(*) from input_list"
	sql = "SELECT c from sbtest where id=?"

	sql = `CREATE TABLE abpzero.new_table (                                      
  idnew_table INT NOT NULL COMMENT '',                                                                                    
  PRIMARY KEY (idnew_table)  COMMENT '')`

	sql = "SELECT state, ROUND(SUM(duration),5) AS `duration (summed) in sec` FROM information_schema.profiling WHERE query_id = 0 GROUP BY state ORDER BY `duration (summed) in sec` DESC"
	sql = "select concat('    -          name : ' , `TABLE_NAME`,' ') as `config` from `information_schema`.`tables` where table_schema='productcatalog' and Table_Type='BASE TABLE' LIMIT 0, 1000"

	t.Log(sql)
	if stmt, err := Parse(sql); err != nil {
		t.Error(err)
	} else {
		t.Log(String(stmt))
		switch v := stmt.(type) {
		case *Update:
			t.Log("Update statement", String(v.Where.Expr))
		}
	}

}
