### 基于SaaS应用的数据分片
每个租户的操作是相互独立的，一个业务db下每张表都会存在一个区分租户的租户字段，因此可以为每张表使用相同的db级别的分片策略，即按租户id进行分片。此做法有两个好处：
* 增、删、改、查，都是仅针对一个租户的，很好的进行了隔离；
* 保留了单个db内表之间的关联查询能力。

### 支持读写分离
* 在主库（Master）上执行:DML（update、insert、replace、delete）、DDL（create、drop、alter）、set语句
* 在从库（Slave）上执行：show、select

读操作，如果有多台读库（Slave），则会以权重轮询方式，进行负载均衡。

### 支持Hint
支持select语句和DDL；set和show语句，暂时不支持。
* 针对select语句，可以通过`/*!saashard master*/`强制在master上执行；
* 针对DDL，可以通过`/*!saashard nodes=node1,node2*/`强制在指定的数据节点列表上执行（不设置默认在所有节点上执行）。

### 后端连接池管理
* 主库（Master）和从库（Slave）的连接池管理
* 连接池最大连接数限制，超出则等待1秒再取，尝试3次仍然失败后则关闭客户端连接