# SaaShard
SaaShard是一款针对SaaS系统设计的数据库分片和读写分离方案。

SaaS系统，具有租户的概念，每个租户间的数据是独立的。基于这个特点，可以通过租户字段进行水平拆分（Sharding）数据库。这样，每个特定租户的SQL执行，都只会定位到单个数据库，这也就保证了可以尽量保留单库查询的能力。

目前只考虑MySQL的实现，并且现阶段只实现按DB分片。

## Unsupported SQL

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
|
```