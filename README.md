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
- Support hint /*! saashard master */ to force execute on master.
- Support split read and write. (Read balance feature is in 'TODO list'.)
- Database sharding feature is in 'TODO list'.

### Unsupported SQL

```
select (case t1.f1 when '0' then 'hello' else 'world' end) f1 from t1 
select f1,f2,f3 into t2 from t1
```

## Code Structure

```
|admin |
|
|backend |
|--------|mysql |
|
|cmd |
|----|saashard |
|
|config |
|
|errors |
|
|net |
|----|mysql |
|
|proxy |
|
|route |
|
|server |
|
|sqlparser |
|----------|sqltypes |
|
|statistic |
|
|utils |
|------|golog |
```