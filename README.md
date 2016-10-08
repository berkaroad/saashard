# SaaShard
[中文版](README.zh-CN.md "中文版")

SaaShard is a MySQL sharding solution, that based on SaaS application.

SaaS application, use multi-tenancy technologyis.
Each tenancy's data is isolated from db logically.For this feature, we can horizontal split data by the field 'tenant_id'.
Each sql statement that related with special tenancy, is only located at single db. So, we can remain the query capability as soon as possible.

## Compile and Run

### Source
```
go get github.com/berkaroad/saashard
cd $GOPATH/src/github.com/berkaroad/saashard
```

### Binary Release
[http://github.com/berkaroad/saashard-binary](http://github.com/berkaroad/saashard-binary "SaaShard Binary")

[https://github.com/berkaroad/docker-images/tree/master/saashard](https://github.com/berkaroad/docker-images/tree/master/saashard "SaaShard Dockerfile")

### Compile

```
make # compile in current platform
make build-all # compile in windows, linux and darwin platform.
```

### Run

```
make test # just for test
make dev # Run immediately, use dev.yaml config file.
make run # Run immediately, use ss.yaml config file.
```

## Features
- Support multi-query and multi-result.
- Support transaction.
- Support hint /*!saashard master */ to force execute on master.
- Support hint /*!saashard nodes=node1,node2 */ to force specify node list in DDL statement.
- Support split read and write. (Read balance use polling algorithm.)
- Support database sharding, supported algorithm is 'hash', 'mod'.
- Support backend connection pool.
- Support Stmt related command.(developing)

## SQL Client Support 
- MySQL Workbench (tested version:6.3)
- SQLyog (tested version:10.2)
- ado.net
- entity framework (EF6)

## SQL Support

- simple query, join query, sub query is supported.
- DML statement
- DDL that only support table struct and index.
- VIEW is not supported, because it couldn't get shard key's value from those.
- maintain FUNCTION, PROCEDURE or TRIGGER is not supported.

```
/* Not supported, but you can replace it to above */
select (case t1.f1 when '0' then 'hello' else 'world' end) f1 from t1;

/* This is supported */
select (case when t1.f1='0' then 'hello' else 'world' end) f1 from t1;
```

```
/* Not supported, because this is dml statement, and couldn't get shard key's value. */
select f1,f2,f3 into t2 from t1
```

## Logical Architecture

![logical architecture](docs/images/logical_arch.png "logical architecture")

![logical schema](docs/images/logical_schema.png "logical schema")

## Contact Info

1. QQ Group: SaaShard 487761803
