# SaaShard
SaaShard是一款针对SaaS系统设计的数据库分片和读写分离方案。

SaaS系统，具有租户的概念，每个租户间的数据是独立的。基于这个特点，可以通过租户字段进行水平拆分（Sharding）数据库。这样，每个特定租户的SQL执行，都只会定位到单个数据库，这也就保证了可以尽量保留单库查询的能力。

目前只考虑MySQL的实现，并且现阶段只实现按DB分片。

## 编译及运行

### 源码
```
go get github.com/berkaroad/saashard
cd $GOPATH/src/github.com/berkaroad/saashard
```

### 二进制发行包
[http://github.com/berkaroad/saashard-binary](http://github.com/berkaroad/saashard-binary "SaaShard Binary")

[https://github.com/berkaroad/docker-images/tree/master/saashard](https://github.com/berkaroad/docker-images/tree/master/saashard "SaaShard Dockerfile")

### 编译

```
make # compile in current platform
make build-all # compile in windows, linux and darwin platform.
```

### 运行

```
make test # just for test
make dev # Run immediately, use dev.yaml config file.
make run # Run immediately, use ss.yaml config file.
```

## 功能
- 支持多语句查询和多结果集返回；
- 支持事务；
- 支持hint /*!saashard master */ 方式，强制在Master上执行；
- 支持hint /*!saashard nodes=node1,node2 */ 方式，强制在指定的节点列表上执行，仅针对DDL语句；
- 支持读写分离（读负载采用权重轮询算法）；
- 支持DB级别分片，目前支持的算法为hash、mod；
- 支持后端连接池和连接数限制；
- 支持STMT相关的命令(正在开发中，暂时只能通过sysbench进行非分片方式的测试)。

## 支持的SQL客户端 
- MySQL Workbench (tested version:6.3)
- SQLyog (tested version:10.2)
- ado.net
- entity framework (EF6)

## 支持的SQL

- 简单查询、Join关联查询、子查询；
- DML语句；
- DDL语句，仅支持表结构和索引的维护；
- 分片模式下，视图不受支持，因为无法从中获取分片字段的值；仅作为读写分离则不受影响；
- 不支持维护函数、存储过程和触发器。

```
/* 不支持，但有替代方案，如下： */
select (case t1.f1 when '0' then 'hello' else 'world' end) f1 from t1;

/* 这种方式是支持的 */
select (case when t1.f1='0' then 'hello' else 'world' end) f1 from t1;
```

```
/* 不支持，此为dml语句，且无法获取分片字段的值 */
select f1,f2,f3 into t2 from t1
```

## 逻辑架构图

![逻辑架构](docs/images/logical_arch.png "逻辑架构")

![逻辑db](docs/images/logical_schema.png "逻辑db")

## 联系方式

1. QQ群: SaaShard 487761803
