# SaaShard
SaaShard是一款针对SaaS系统设计的数据库分片和读写分离方案。

SaaS系统，具有租户的概念，每个租户间的数据是独立的。基于这个特点，可以通过租户字段进行水平拆分（Sharding）数据库。这样，每个特定租户的SQL执行，都只会定位到单个数据库，这也就保证了可以尽量保留单库查询的能力。

目前只考虑MySQL的实现，并且现阶段只实现按DB分片。

## Compile and Run

### Source
```
go get github.com/berkaroad/saashard
cd $GOPATH/src/github.com/berkaroad/saashard
```

### Binary release
[http://github.com/berkaroad/saashard-binary](http://github.com/berkaroad/saashard-binary)

### Compile

```
make # compile in current platform
make build-all  # compile in windows, linux  and darwin platform.
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
- Support maintain table struct, index.

### Unsupported SQL

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

```
/* Not supported, because it couldn't get shard key's value from those. */
VIEW, FUNCTION, PROCEDURE or TRIGGER not supported
```

## Logical Architecture

![logical architecture](docs/images/logical_arch.png "logical architecture")

![logical schema](docs/images/logical_schema.png "logical schema")

## Contact Info

1. QQ Group: SaaShard 487761803