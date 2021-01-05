# 数据库知识综合
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
&emsp;对于mysql数据库中的join实现算法为`Nested-loop join Algorithms` [2],根据该算法的实现逻辑在执行mysql join操作时实际上是执行循环对比，因此优化的思路和目标即是尽量将循环的次数变小。在优化join操作时典型的优化方法有：  
  - 尽量选择小表作为驱动表
  - 对被驱动表的join字段添加索引
  

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
  
## MySQL相关

## 参考
[1]. https://www.codeproject.com/Articles/33052/Visual-Representation-of-SQL-Joins  
[2]. https://dev.mysql.com/doc/refman/5.7/en/nested-loop-joins.html
[3]. https://juejin.cn/post/6844903573935882247
[4]. https://www.oreilly.com/library/view/high-performance-mysql/9780596101718/ch04.html
[5]. https://medium.com/faun/query-optimization-with-mysql-740460747d89
[6]. https://stackoverflow.blog/2020/10/14/improve-database-performance-with-connection-pooling/
[7]. https://dev.mysql.com/doc/refman/8.0/en/nested-join-optimization.html