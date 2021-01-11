# 数据库知识综合
## 数据库基础
- 数据库范式  
&emsp;范式，即Normal Form，指的是我们在构建数据库所需要遵守的规则和指导方针。
首先要明确的是：满足第三范式，那么就一定满足第二范式、满足第二范式就一定满足第一范式
&emsp;第一范式：字段是最小的的单元不可再分。即列的原子性。
学生信息组成学生信息表，有年龄、性别、学号等信息组成。这些字段都不可再分，所以它是满足第一范式的
&emsp;第二范式：满足第一范式,表中的字段必须完全依赖于全部主键而非部分主键。
比如一个订单表中有主键（订单id，商品id）和单位价格、折扣、数量、产品名称、产品保质期。保质期只依赖于商品id，不符合第二范式。
&emsp;第三范式：满足第二范式，非主键外的所有字段必须互不依赖
换句话说：数据只能存在一个表中，消除互相依赖
比如大学学院表中包含了学院id，院领导，院简介，如果在学生信息表中也包括了院领导，院简介这些字段，这就重复了。
- MySQL事务特性  
  数据库事务transanction正确执行的四个基本要素。ACID,持久性(Durability)、原子性(Atomicity)、一致性(Correspondence)、隔离性(Isolation)。
  - 持久性: 在事务完成以后，该事务所对数据库所作的更改便持久的保存在数据库之中，并不会被回滚。
  - 原子性: 整个事务中的所有操作，要么全部完成，要么全部不完成，不可能停滞在中间某个环节。事务在执行过程中发生错误，会被回滚（Rollback）到事务开始前的状态，就像这个事务从来没有执行过一样。
  - 一致性: 在事务开始之前和事务结束以后，数据库的完整性约束没有被破坏。
  - 隔离性: 隔离状态执行事务，使它们好像是系统在给定时间内执行的唯一操作。如果有两个事务，运行在相同的时间内，执行 相同的功能，事务的隔离性将确保每一事务在系统中认为只有该事务在使用系统。这种属性有时称为串行化，为了防止事务操作间的混淆，必须串行化或序列化请 求，使得在同一时间仅有一个请求用于同一数据。
- 存储过程  
&emsp;存储过程就是一段SQL语句的预编译集合，封装了一组sql语句，实现某些操作，类似于函数的功能。
  - 好处：  
  将代码封装起来，隐藏复杂的商业逻辑  
  预编译，执行效率高  
  可以接受参数，可以回传值
  - 坏处:  
  针对特定的某种数据库，不兼容，难维护  
  - 触发器与存储过程  
&emsp;触发器与存储过程非常相似，触发器也是SQL语句集，两者唯一的区别是触发器不能用EXECUTE语句调用，而是在用户执行Transact-SQL语句时自动触发（激活）执行。触发器是在一个修改了指定表中的数据时执行的存储过程。通常通过创建触发器来强制实现不同表中的逻辑相关数据的引用完整性和一致性。由于用户不能绕过触发器，所以可以用它来强制实施复杂的业务规则，以确保数据的完整性。触发器不同于存储过程，触发器主要是通过事件执行触发而被执行的，而存储过程可以通过存储过程名称名字而直接调用。当对某一表进行诸如UPDATE、INSERT、DELETE这些操作时，SQLSERVER就会自动执行触发器所定义的SQL语句，从而确保对数据的处理必须符合这些SQL语句所定义的规则。  
- 视图
视图是从一个或多个表导出的虚拟的表，具有普通表的结构，但是不实现数据存储。  
  - 作用：  
    直观  
    安全性，暴露出视图，然后把不想让用户看到和修改的内容屏蔽掉  
    独立性，屏蔽了真实表的结构带来的影响。  
  - 缺点：  
    性能差：视图是由一个复杂的多表查询所定义  
    修改限制： 当用户试图修改视图的某些信息时，数据库必须把它转化为对基本表的某些信息的修改
- 事务隔离级别  
  - Read Uncommitted（读取未提交内容）：在该隔离级别，所有事务都可以看到其他未提交事务的执行结果。本隔离级别很少用于实际应用，因为它的性能也不比其他级别好多少。读取未提交的数据，也被称之为脏读（Dirty Read）。
  - Read Committed（读取提交内容）：这是大多数数据库系统的默认隔离级别（但不是MySQL默认的）。它满足了隔离的简单定义：一个事务只能看见已经提交事务所做的改变。这种隔离级别 也支持所谓的不可重复读（Nonrepeatable Read），因为同一事务的其他实例在该实例处理其间可能会有新的commit，所以同一select可能返回不同结果。
  - Repeatable Read（可重读）：这是MySQL的默认事务隔离级别，它确保同一事务的多个实例在并发读取数据时，会看到同样的数据行。不过理论上，这会导致另一个棘手的问题：幻读 （Phantom Read）。简单的说，幻读指当用户读取某一范围的数据行时，另一个事务又在该范围内插入了新行，当用户再读取该范围的数据行时，会发现有新的“幻影” 行。InnoDB和Falcon存储引擎通过多版本并发控制（MVCC，Multiversion Concurrency Control）机制解决了该问题。
  - Serializable（可串行化）：这是最高的隔离级别，它通过强制事务排序，使之不可能相互冲突，从而解决幻读问题。简言之，它是在每个读的数据行上加上共享锁。在这个级别，可能导致大量的超时现象和锁竞争。

  这四种隔离级别采取不同的锁类型来实现，若读取的是同一个数据的话，就容易发生问题。例如：脏读(Drity Read)：某个事务已更新一份数据，另一个事务在此时读取了同一份数据，由于某些原因，前一个RollBack了操作，则后一个事务所读取的数据就会是不正确的。不可重复读(Non-repeatable read):在一个事务的两次查询之中数据不一致，这可能是两次查询过程中间插入了一个事务更新的原有的数据。幻读(Phantom Read):在一个事务的两次查询中数据笔数不一致，例如有一个事务查询了几列(Row)数据，而另一个事务却在此时插入了新的几列数据，先前的事务在接下来的查询中，就会发现有几列数据是它先前所没有的。

|                 |脏读|不可重复读	| 幻读 |    	
|-----            |---|---	       |---	 |
| Read uncommitted| √ | √	         |   	√| 	
| Read committed  | × | √  	       |    √|
| Repeatable read |×  |×           |   √ |
| Serializable 	  | × |  ×         |   × |
|
    - Read committed要求必须读取已提交的数据。  
    - Repeatable read要求读取过程中，其他事务不能修改(图中第一次读取第二次读取中间都算读取过程)，然而读取-写入却不算在内，因此依然有幻读风险。
    - Serializable要求按串行读取，效率低。
- MySQL存储引擎  
最常见的是InnoDB和MyISAM

|Innodb|	myisam|
|-----|---------|
|事务 |	支持|	不支持|
|外键	|支持	|不支持|
|全文本搜索|不支持|	支持
|使用场景	|频繁修改	|查询和插入为主
|

&emsp;另外还有MEMORY，存储在内存中，速度快，安全性不高  
1.InnoDB：默认存储引擎，使用最广泛。  
2.MyISAM：表锁，不支持事务。  
3.Archive：适合日志和数据采集类应用。  
4.Memory：适合访问速度快，数据丢失也没有关系的场景。  
5.CSV：将普通csv保存再MySQL中，主要用于数据交换。    
&emsp;此外还有：Blackhole、Federated、Merge、NDB等存储引擎。

MySQL InooDB和MyISAM存储引擎区别  
   - - 存储结构
   MyISAM：每个MyISAM在磁盘上存储成三个文件。第一个文件的名字以表的名字开始，扩展名指出文件类型。.frm文件存储表定义。数据文件的扩展名为.MYD (MYData)。索引文件的扩展名是.MYI (MYIndex)。  
  InnoDB：所有的表都保存在同一个数据文件中（也可能是多个文件，或者是独立的表空间文件），InnoDB表的大小只受限于操作系统文件的大小，一般为2GB。
      -  存储空间  
      MyISAM：可被压缩，存储空间较小。支持三种不同的存储格式：静态表(默认，但是注意数据末尾不能有空格，会被去掉)、动态表、压缩表。  
      InnoDB：需要更多的内存和存储，它会在主内存中建立其专用的缓冲池用于高速缓冲数据和索引
      - 可移植性、备份及恢复  
      MyISAM：数据是以文件的形式存储，所以在跨平台的数据转移中会很方便。在备份和恢复时可单独针对某个表进行操作。  
      InnoDB：免费的方案可以是拷贝数据文件、备份 binlog，或者用 mysqldump，在数据量达到几十G的时候就相对痛苦了。
      - 事务支持 
      MyISAM：强调的是性能，每次查询具有原子性,其执行数度比InnoDB类型更快，但是不提供事务支持。  
      InnoDB：提供事务支持事务，外部键等高级数据库功能。 具有事务(commit)、回滚(rollback)和崩溃修复能力(crash recovery capabilities)的事务安全(transaction-safe (ACID compliant))型表。
      - AUTO_INCREMENT 
      MyISAM：可以和其他字段一起建立联合索引。引擎的自动增长列必须是索引，如果是组合索引，自动增长可以不是第一列，他可以根据前面几列进行排序后递增。  
      InnoDB：InnoDB中必须包含只有该字段的索引。引擎的自动增长列必须是索引，如果是组合索引也必须是组合索引的第一列。
      - 表锁差异  
      MyISAM：只支持表级锁，用户在操作myisam表时，select，update，delete，insert语句都会给表自动加锁，如果加锁以后的表满足insert并发的情况下，可以在表的尾部插入新的数据。  
      InnoDB：支持事务和行级锁，是innodb的最大特色。行锁大幅度提高了多用户并发操作的新能。但是InnoDB的行锁，只是在WHERE的主键是有效的，非主键的WHERE都会锁全表的
      - 全文索引  
      MyISAM：支持 FULLTEXT类型的全文索引
      InnoDB：5.7之后支持FULLTEXT类型的全文索引
      - 表主键  
      MyISAM：允许没有任何索引和主键的表存在，索引都是保存行的地址。  
      InnoDB：如果没有设定主键或者非空唯一索引，就会自动生成一个6字节的主键(用户不可见)，数据是主索引的一部分，附加索引保存的是主索引的值。
      - 表的具体行数   
      MyISAM：保存有表的总行数，如果select count(*) from table;会直接取出出该值。
      InnoDB：没有保存表的总行数，如果使用select count(*) from table；就会遍历整个表，消耗相当大，但是在加了wehre条件后，myisam和innodb处理的方式都一样。
      - CURD操作  
      MyISAM：如果执行大量的SELECT，MyISAM是更好的选择。  
      InnoDB：如果你的数据执行大量的INSERT或UPDATE，出于性能方面的考虑，应该使用InnoDB表。DELETE 从性能上InnoDB更优，但DELETE FROM table时，InnoDB不会重新建立表，而是一行一行的删除，在innodb上如果要清空保存有大量数据的表，最好使用truncate table这个命令。
      - 外键
      MyISAM：不支持  
      InnoDB：支持
      - 崩溃自动恢复  
      MyISAM：不支持    
      InnoDB：支持  

- MySQL乐观锁和悲观锁  
  悲观锁和乐观锁都是为保证一致性的一种锁。
  - 悲观锁  
  ```
  1.关闭autocommit=0;
  2.在事务中使用 select .. from … where … for update;给行加排他锁  
  3.select命中的行必须有索引，否则会锁表  
  优点：  
    1.保守策略，所以数据安全性高  
  缺点：
    1.有加锁等额外开销，效率低
    2.可能引起死锁
    3.降低并行行，数据被锁住后其他事物必须等待
  ```
  
  - 乐观锁
```
使用：
  1.表中增加版本号或时间戳数据列
  2.读取数据是同时读取版本号
  3.更新数据时添加版本号为条件，同时版本号增加1
  4.如果更更新失败，提示用户
优点：
  1.没有锁，效率高
  2.不会引起死锁
缺点：
1.这里是列表文本遇到两个事务统一时间读取一行数据时，会引起问题
```
- MySQL索引    
&emsp;mysql 索引类型有:
1.主键索引  
2.普通索引  
3.唯一索引  
4.全文索引  
&emsp;索引本身就是一种数据结构，用于加快查找速度，InnoDB和MyISAM都是用的B+树。
经常需要查询，经常使用SELECT和WHERE操作时，可以使用索引
经常做表连接
经常出现在order by、group by、distinct 后面的字段中，可以建立索引  
&emsp;索引的缺点在于：增删改浪费时间，构建索引需要占据空间
构建索引常用的数据结构

| 结构| 区别| 
|----|--- |
|Hash|只存储对应的哈希值，查找速度快，不能排序,不能进行范围查询|
|B+|数据有序,范围查询|  

 聚集索引和非聚集索引的区别
 
|索引|	区别|
|---|----|
|聚集索引|	数据按索引顺序存储，数据行的物理顺序与列值的顺序相同
|非聚集索引|	存储指向真正数据行的指针
|
聚簇索引的叶节点就是数据节点，而非聚簇索引的叶节点仍然是索引节点，并保留一个链接指向对应数据块。  
MyISAM的是非聚簇索引，B+Tree的叶子节点上的data，并不是数据本身，而是数据存放的地址。  
InnoDB使用的是聚簇索引，将主键组织到一棵B+树中，而行数据就储存在叶子节点上  
MySQL InnoDB一定会建立聚簇索引，把实际数据行和相关的键值保存在一块，这也决定了一个表只能有一个聚簇索引  
1.InnoDB通常根据主键值(primary key)进行聚簇  
2.如果没有创建主键，则会用一个唯一且不为空的索引列做为主键，成为此表的聚簇索引  
3.上面二个条件都不满足，InnoDB会自己创建一个虚拟的聚集索引  
聚簇索引的  
优点：就是提高数据访问性能。  
缺点：维护索引很昂贵，特别是插入新行或者主键被更新导至要分页(page split)的时候。


索引底层用B+树而不是用红黑树

|树	|区别|
|---|---|
|红黑树	|增加，删除，红黑树会进行频繁的调整，来保证红黑树的性质，浪费时间|
B树也就是B-树	|B树，查询性能不稳定，查询结果高度不致，每个结点保存指向真实数据的指针，相比B+树每一层每屋存储的元素更多，显得更高一点。|
B+树	|B+树相比较于另外两种树,显得更矮更宽，查询层次更浅|
|
 
| B树和 B+树区别
  第一，B 树一个节点里存的是数据，而 B+树存储的是索引（地址），所以 B 树里一个节点存不了很多个数据，但是 B+树一个节点能存很多索引，B+树叶子节点存所有的数据。  
  第二，B+树的叶子节点是数据阶段用了一个链表串联起来，便于范围查找。

  通过 B 树和 B+树的对比我们看出，B+树节点存储的是索引，在单个节点存储容量有限的情况下，单节点也能存储大量索引，使得整个 B+树高度降低，减少了磁盘 IO。其次，B+树的叶子节点是真正数据存储的地方，叶子节点用了链表连接起来，这个链表本身就是有序的，在数据范围查找时，更具备效率。因此 Mysql 的索引用的就是 B+树，B+树在查找效率、范围查找中都有着非常不错的性能。

索引失效的条件    
 - 在where子句中进行null值判断
 - 避免在where子句中使用or来连接条件,因为如果俩个字段中有一个没有索引的话,引擎会放弃索引而产生全表扫描
 - 避免在where子句中使用like模糊查询
  
- 数据库连接池  
维护一定数量的连接，减少重新创建连接的时间
- MVCC  
使用MVCC时，不会直接用新数据覆盖旧数据，而是将旧数据标记为过时并在别处增加新版本的数据。允许读者读取在他读之前已经存在的数据，即使这些在读的过程中半路被别人修改、删除了，也对先前正在读的用户没有影响。
## SQL
- SQL JOIN 的可视化  
![sql_join](../../asserts/images/md/comprehensive/database/sql_joins.png)  

sql join 的sql语句如下[1]：  

```
SELECT <select_list> FROM TableA A LEFT JOIN TABLEB B ON A.key=B.key
```
```
SELECT <select_list>  FROM Table_A A INNER JOIN Table_B B ON A.Key = B.Key
```
```
SELECT <select_list> FROM Table_A A RIGHT JOIN Table_B B
ON A.Key = B.Key
```
```
SELECT <select_list> FROM Table_A A FULL OUTER JOIN Table_B B
ON A.Key = B.Key

```
```
SELECT <select_list>  FROM Table_A A LEFT JOIN Table_B B ON A.Key = B.Key
WHERE B.Key IS NULL
```
```
SELECT <select_list> FROM Table_A A RIGHT JOIN Table_B B ON A.Key = B.Key
WHERE A.Key IS NULL
```
```
SELECT <select_list> FROM Table_A A FULL OUTER JOIN Table_B B
ON A.Key = B.Key
WHERE A.Key IS NULL OR B.Key IS NULL
```
 - &emsp;对于mysql数据库中的join实现算法为`Nested-loop join Algorithms` [2],根据该算法的实现逻辑在执行mysql join操作时实际上是执行循环对比，因此优化的思路和目标即是尽量将循环的次数变小。在优化join操作时典型的优化方法有：  
    -   尽量选择小表作为驱动表
      - 对被驱动表的join字段添加索引
  
- SQL中高效率获取随机N调数据
```
SELECT *
FROM `TABLE_NAME` AS t1 JOIN (SELECT ROUND(RAND() * (SELECT MAX(id) FROM `TABLE_NAME`)) AS id) AS t2
WHERE t1.id >= t2.id
ORDER BY t1.id ASC LIMIT 4;
或
SELECT * FROM `TABLE_NAME` 
WHERE id >= (SELECT floor(RAND() * (SELECT MAX(id) FROM `TABLE_NAME`)))  and city="city_91" and showSex=1
ORDER BY id LIMIT 4;
或 
SELECT * FROM TABLE_NAME 
WHERE id >= ((SELECT MAX(id) FROM TABLE_NAME)-(SELECT MIN(id) FROM TABLE_NAME)) * RAND() + (SELECT MIN(id) FROM TABLE_NAME)
limit 5;
或
SELECT *
FROM `TABLE_NAME` AS t1 JOIN (SELECT ROUND(RAND() * (
(SELECT MAX(id) FROM `TABLE_NAME` where id<1000 )-(SELECT MIN(id) FROM `TABLE_NAME` where id<1000 ))+(SELECT MIN(id) FROM `TABLE_NAME` where id<1000 )) AS id) AS t2
WHERE t1.id >= t2.id
ORDER BY t1.id LIMIT 5;
```
## SQL 操作过程中的性能优化方法
**sql 查询的优化的主要方向在优化索引的使用和避免全表扫描**
  - join操作时
      - 尽量选择小表作为驱动表
      - 对被驱动表的join字段添加索引
  - select 查询操作
    - 采用limit进行分页查询时结合使用order by id 索引 如  
  ```
  select id from A order by id limit 90000,10;
  或
  select id from A order by id  between 90000 and 90010;
  ```
  - 采用top或limit限制查询结果
    ```
    SELECT id FROM A LIKE 'abc%' limit 1
    ```
  - 查询尽量不要使用select * 进行查询
  ```
  SELECT <select_field> FROM TABLE
  ```
 - like查询时尽量避免%filed%查询，最好单侧执行%
  ```
  SELECT <select_file> FROM TABLE LIKE 'FILED%'
或
  SELECT <select_file> FROM TABLE LIKE '%FILED'

  ```
  - 在where子句中使用or时可以考虑采用union all 或union方式进行替换
```
SELECT id FROM TABLEA WHERE field= A or field = B 修改为
SELECT id FROM TABLEA WHERE  field= A union all SELECT id FROM TABLEA WHERE field=B 
```
 - IS NULL 或 IS NOT NULL 子句中尽量修改为特殊值查询 如将IS NULL 字段置为0 从而避免全表扫描
  ```
  SELECT id FROM A WHERE filed IS NULL 修改为
  SELECT id FROM A WHERE field = 0
  ```
- WHERE 子句中尽量避免使用函数计算，将确定值作为where子句的查询条件
  ```
  SELECT id FROM A WHERE year(addate) <2020 修改为
  SELECT id FROM A where addate<'2020-01-01'
  ```
- 多字段排序问题
  - mysql查询时是只会使用一个索引，如果where子句中已经使用了索引，那么在order by时是不会使用索引的，尽量避免执行多字段的排序操作，不得己时简易对排序字段建立复合索引

- 尽量使用union all替换union
  - union操作会将两个或多个结果集再合并后在进行唯一性过滤，这一过程会设计到排序以及大量的cpu计算。当确认不会出现重复结果集或重复结果集不重要的情况下，使用union all替换union
- inner join与outer join
  - 尽量使用inner join
  - 使用join时，选择小表作为驱动表
- exits与in 比较
  ```
  SELECT * from A WHERE id in ( SELECT id from B )
  ```
  - select in 是在内存中进行遍历比较，exists需要查询数据库，当B表中数据比较大是exits效率优于in；而B表数据比A小的情况下用in比较适合。
## MySQL相关
### MySQL 性能优化
- 选择合适的存储引擎，通常使用InnoDB  
 MyISAM 只缓存索引，而 InnoDB 缓存数据和索引，MyISAM 不支持事务；
- 为每个表分别创建 InnoDB FILE  
这样可以保证 ibdata1 文件不会过大，失去控制。尤其是在执行 mysqlcheck -o –all-databases 的时候,设置方法如下
```
innodb_file_per_table=1
```
- 设置合适的innodb_buffer_pool_size 以尽量保证从内存中读取数据而不是从磁盘中读取数据  
在设置上述参数时也要注意调整服务器本身的内存设置，如`执行 echo 1 > /proc/sys/vm/drop_caches 清除操作系统的文件缓存` 后查看服务器可用内存，  并且需要执行数据预热处理。
- 关闭swap分区，尽量避免数据存储在swap分区中
- 定期优化和重建数据库   
mysqlcheck -o –all-databases 会让 ibdata1 不断增大，真正的优化只有重建数据表结构.
- 减少磁盘写入操作
  - 使用足够大的写入缓存 innodb_log_file_size
  - innodb_flush_log_at_trx_commit
  - 避免双写入缓冲 `innodb_flush_method=O_DIRECT`
  - 提高磁盘读写速度
- 充分利用索引
- 分析查询日志和慢查询日志
- 架构优化
  - 分库分表（垂直，水平）
  - 缓存
  - 读写分离
  - 主从复制
- 配置优化
  - 数据库配置优化
  ```
  Linux系统中MySQl配置文件一般位于/etc/my.cnf
  常用配置参数：
  innodb_buffer_pool_size【用于配置Innodb的缓 冲池,如果数据库中只有Innodb表，则推荐配置量为总内存的75%】
  innodb_buffer_pool_instances【MySQL5.5中新增参数，可以控制缓冲池的个数，默认情况下只有一个缓冲池】
  innodb_flush_log_at_trx_commit【关键参数，对innodb的IO影响很大。默认值为1，可以取0,1,2三个值，0最快，1最安全，2折中。一般建议设为2，但如果数据安全性要求比较高则使用默认值1】
  ```
  - 操作系统配置优化
  ```
  网络方面的配置，要修改/etc/sysctl.conf文件
  增加tcp支持的队列数
  net.ipv4.tcp_max_syn_backlog = 65535
  减少断开连接时 ，资源回收
  net.ipv4.tcp_max_tw_buckets = 8000
  net.ipv4.tcp_tw_reuse = 1
  net.ipv4.tcp_tw_recycle = 1
  net.ipv4.tcp_fin_timeout = 10
   ```
- 优化顺序和优化思路  
1. 发现有问题的SQL   
 读取MYSQL的慢查询日志，记录在MySQL中响应时间超过阈值的语句，可以查询出执行的次数多占用的时间长的SQL。
2. 通过EXPLAIN关键字分析SQL  
使用 EXPLAIN 关键字可以知道MySQL是如何处理你的SQL语句的
3. 进行优化

## 参考
[1]. https://www.codeproject.com/Articles/33052/Visual-Representation-of-SQL-Joins  
[2]. https://dev.mysql.com/doc/refman/5.7/en/nested-loop-joins.html  
[3]. https://juejin.cn/post/6844903573935882247  
[4]. https://www.oreilly.com/library/view/  high-performance-mysql/9780596101718/ch04.html  
[5]. https://medium.com/faun/query-optimization-with-mysql-740460747d89  
[6]. https://stackoverflow.blog/2020/10/14/  improve-database-performance-with-connection-pooling/  
[7]. https://dev.mysql.com/doc/refman/8.0/en/ nested-join-optimization.html  
[8]. https://coolshell.cn/articles/1846.html  
[9]. https://www.jishuchi.com/read/mysql-interview/2809  
[10]. https://github.com/jeremycole/innodb_diagrams  
[11].https://blog.jcole.us/innodb/  
[12].https://dba.stackexchange.com/questions/204561/does-mysql-use-b-tree-btree-or-both  
[13].https://draveness.me/whys-the-design-mysql-b-plus-tree/
[14]. https://www.vertabelo.com/blog/  all-about-indexes-part-2-mysql-index-structure-and-performance/
