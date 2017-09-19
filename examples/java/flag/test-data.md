MySQL test data
```
[vagrant@bogon resources]$ ls mysql-dump.sql 
mysql-dump.sql
[vagrant@bogon resources]$ docker ps -f name=demo
CONTAINER ID        IMAGE                    COMMAND                  CREATED             STATUS              PORTS                                     NAMES
8bfbf42555c3        tangfeixiong/lemp-demo   "/usr/bin/supervisord"   3 weeks ago         Up 12 days          443/tcp, 3306/tcp, 0.0.0.0:8999->80/tcp   demo
[vagrant@bogon resources]$ docker cp mysql-dump.sql 8bfbf42555c3:/tmp
[vagrant@bogon resources]$ docker exec -ti demo bash
[root@8bfbf42555c3 /]# cd /tmp
[root@8bfbf42555c3 tmp]# ls
ks-script-ffclxw  mysql-stderr---supervisor-gv2Eow.log	nginx-stderr---supervisor-D_yaSF.log  php-fpm-stderr---supervisor-5KcJe3.log  yum.log
mysql-dump.sql	  mysql-stdout---supervisor-9lrheQ.log	nginx-stdout---supervisor-WOy0xT.log  php-fpm-stdout---supervisor-xbMesJ.log
[root@8bfbf42555c3 tmp]# mysql -u testuser -p
Enter password: 
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 71
Server version: 5.7.19 MySQL Community Server (GPL)

Copyright (c) 2000, 2017, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> use testdb 
Database changed
mysql> show tables;
Empty set (0.00 sec)

mysql> source mysql-dump.sql

mysql> show tables;
+------------------+
| Tables_in_testdb |
+------------------+
| admin            |
| config           |
| flag             |
| ip_filter        |
| token            |
+------------------+
5 rows in set (0.01 sec)

mysql> select * from admin;
+----+----------+--------+
| id | username | passwd |
+----+----------+--------+
|  1 | admin    | 123456 |
+----+----------+--------+
1 row in set (0.00 sec)

mysql> select * from config;
+----+----------+-------+-------------------+-----------+
| id | common   | count | environment_count | tremcount |
+----+----------+-------+-------------------+-----------+
|  1 | 20170826 |   100 |                 3 |         5 |
+----+----------+-------+-------------------+-----------+
1 row in set (0.00 sec)

mysql> select * from flag;  
+-------+-------+---------+-------+------+----------------------------------+
| id    | round | team_no | token | env  | md5string                        |
+-------+-------+---------+-------+------+----------------------------------+
| 19671 |     1 |       1 | 121   |    1 | 11BF33B2D7F8A37B59F6A128A86693FE |
| 19672 |     2 |       1 | 121   |    1 | C2B4CA9C56973A3EBD65BC91ADB517E3 |
+-------+-------+---------+-------+------+----------------------------------+
2 rows in set (0.00 sec)

mysql> select * from ip_filter;
+----+---------------------------------------------------------------+
| id | value                                                         |
+----+---------------------------------------------------------------+
|  1 | 192.168.0.5,172.17.4.50,172.17.0.9,192.168.1.208,192.168.1.45 |
+----+---------------------------------------------------------------+
1 row in set (0.00 sec)

mysql> select * from token;    
Empty set (0.00 sec)

mysql> INSERT INTO `token` VALUES (231,NULL,''),(232,NULL,''),(233,NULL,''),(234,NULL,''),(235,NULL,''),(236,NULL,''),(237,NULL,'');
Query OK, 7 rows affected (0.00 sec)
Records: 7  Duplicates: 0  Warnings: 0

mysql> select * from token;                                                                                                         
+-----+---------+-------------+
| id  | team_no | token_value |
+-----+---------+-------------+
| 231 |    NULL |             |
| 232 |    NULL |             |
| 233 |    NULL |             |
| 234 |    NULL |             |
| 235 |    NULL |             |
| 236 |    NULL |             |
| 237 |    NULL |             |
+-----+---------+-------------+
7 rows in set (0.00 sec)

mysql> quit
Bye
[root@8bfbf42555c3 tmp]# rm mysql-dump.sql 
rm: remove regular file 'mysql-dump.sql'? y
[root@8bfbf42555c3 tmp]# exit
exit
```
