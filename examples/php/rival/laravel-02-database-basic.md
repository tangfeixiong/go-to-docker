# Laravel Database and ORM

Docker
```
[vagrant@bogon laravel-app]$ docker ps -f name=demo
CONTAINER ID        IMAGE                    COMMAND                  CREATED             STATUS              PORTS                                     NAMES
8bfbf42555c3        tangfeixiong/lemp-demo   "/usr/bin/supervisord"   3 weeks ago         Up 9 days           443/tcp, 3306/tcp, 0.0.0.0:8999->80/tcp   demo
```

MySQL
```
[vagrant@bogon laravel-app]$ docker inspect -f {{.NetworkSettings.IPAddress}} demo
172.17.0.8
```

## Table of contents

Schema

Connection configuration

Model

Reference

## Schema

__Battlefield__
```
[vagrant@bogon laravel-app]$ php artisan make:migration create_battlefield_table --create=battlefield
Created Migration: 2017_09_11_180644_create_battlefield_table
[vagrant@bogon laravel-app]$ ls database/migrations/2017_09_11_180644_create_battlefield_table.php 
database/migrations/2017_09_11_180644_create_battlefield_table.php
```

more
```
[vagrant@bogon laravel-app]$ php artisan make:migration add_stuff_to_battlefield_table --table=battlefield
Created Migration: 2017_09_11_181043_add_stuff_to_battlefield_table
```

__Issue__

缺少PDO组件
```
[vagrant@bogon laravel-app]$ php artisan migrate
PHP Fatal error:  Class 'PDO' not found in /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/examples/php/rival/laravel-app/vendor/laravel/framework/src/Illuminate/Database/Connection.php on line 1200

                                                           
  [Symfony\Component\Debug\Exception\FatalErrorException]  
  Class 'PDO' not found                                    
                                                           
```

Resolve
```
[vagrant@bogon laravel-app]$ sudo dnf search php-pdo
上次元数据过期检查在 2:09:05 前执行于 Mon Sep 11 16:48:24 2017。
=================================================================== N/S 匹配：php-pdo ====================================================================
php-pdo.i686 : A database access abstraction module for PHP applications
php-pdo.x86_64 : A database access abstraction module for PHP applications
[vagrant@bogon laravel-app]$ sudo dnf install -y php-pdo
上次元数据过期检查在 2:10:51 前执行于 Mon Sep 11 16:48:24 2017。
依赖关系解决。
==========================================================================================================================================================
 Package                            架构                              版本                                       仓库                                大小
==========================================================================================================================================================
安装:
 php-pdo                            x86_64                            5.6.29-1.fc23                              updates                            153 k

事务概要
==========================================================================================================================================================
安装  1 Package

总下载：153 k
安装大小：401 k
下载软件包：
php-pdo-5.6.29-1.fc23.x86_64.rpm                                                                                          179 kB/s | 153 kB     00:00    
----------------------------------------------------------------------------------------------------------------------------------------------------------
总计                                                                                                                       67 kB/s | 153 kB     00:02     
运行事务检查
事务检查成功。
运行事务测试
事务测试成功。
运行事务
  安装: php-pdo-5.6.29-1.fc23.x86_64                                                                                                                  1/1 
  验证: php-pdo-5.6.29-1.fc23.x86_64                                                                                                                  1/1 

已安装:
  php-pdo.x86_64 5.6.29-1.fc23                                                                                                                            

完毕！
```

__Issue__

缺少POD for MySQL组件
```
[vagrant@bogon laravel-app]$ php artisan migrate

                                                                                                                                   
  [Illuminate\Database\QueryException]                                                                                             
  could not find driver (SQL: select * from information_schema.tables where table_schema = homestead and table_name = migrations)  
                                                                                                                                   

                         
  [PDOException]         
  could not find driver  
                         
```

Resolve
```
[vagrant@bogon laravel-app]$ sudo dnf search php-mysql
上次元数据过期检查在 2:20:08 前执行于 Mon Sep 11 16:48:24 2017。
================================================================== N/S 匹配：php-mysql ===================================================================
php-mysqlnd.i686 : A module for PHP applications that use MySQL databases
php-mysqlnd.x86_64 : A module for PHP applications that use MySQL databases
[vagrant@bogon laravel-app]$ sudo dnf install -yq php-mysqlnd
```

__Files__
```
[vagrant@bogon laravel-app]$ ls /usr/lib64/php/modules/{mysql*,pdo*}
/usr/lib64/php/modules/mysqli.so   /usr/lib64/php/modules/mysql.so      /usr/lib64/php/modules/pdo.so
/usr/lib64/php/modules/mysqlnd.so  /usr/lib64/php/modules/pdo_mysql.so  /usr/lib64/php/modules/pdo_sqlite.so
[vagrant@bogon laravel-app]$ php --ini | egrep '(pdo|mysql)'
/etc/php.d/20-mysqlnd.ini,
/etc/php.d/20-pdo.ini,
/etc/php.d/30-mysql.ini,
/etc/php.d/30-mysqli.ini,
/etc/php.d/30-pdo_mysql.ini,
/etc/php.d/30-pdo_sqlite.ini,
```

__Issue__

数据库连接失败
```
[vagrant@bogon laravel-app]$ php artisan migrate
                                                                                                                                                       
  [Illuminate\Database\QueryException]                                                                                                                 
  SQLSTATE[HY000] [2002] Connection refused (SQL: select * from information_schema.tables where table_schema = homestead and table_name = migrations)  
                                                                                                                                                       

                                             
  [PDOException]                             
  SQLSTATE[HY000] [2002] Connection refused  
                                             
```

Or 
```
                                                                                                                                                         
  [Illuminate\Database\QueryException]                                                                                                                   
  SQLSTATE[HY000] [1049] Unknown database 'homestead' (SQL: select * from information_schema.tables where table_schema = homestead and table_name = mig  
  rations)                                                                                                                                               
                                                                                                                                                         

                                                       
  [PDOException]                                       
  SQLSTATE[HY000] [1049] Unknown database 'homestead'  
                                                                                                                                                         
```                                                       

or
```
  [Illuminate\Database\QueryException]                                                                                                                   
  SQLSTATE[HY000] [1045] Access denied for user 'homestead'@'gateway' (using password: YES) (SQL: select * from information_schema.tables where table_s  
  chema = homestead and table_name = migrations)                                                                                                         
                                                                                                                                                         

                                                                                             
  [PDOException]                                                                             
  SQLSTATE[HY000] [1045] Access denied for user 'homestead'@'gateway' (using password: YES)  
                                                                                             
```

修改.env文件
```
[vagrant@bogon laravel-app]$ cat .env | grep DB_
DB_CONNECTION=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=homestead
DB_USERNAME=homestead
DB_PASSWORD=secret
[vagrant@bogon laravel-app]$ sed -i 's/\(DB_HOST=127.0.0.1\)/# \1\nDB_HOST=172.17.0.8/' .env
[vagrant@bogon laravel-app]$ sed -i 's/\(DB_DATABASE=homestead\)/# \1\nDB_DATABASE=rival/' .env
[vagrant@bogon laravel-app]$ sed -i 's/\(DB_USERNAME=homestead\)/# \1\nDB_USERNAME=root/;s/\(DB_PASSWORD=\).*/\1password/' .env
[vagrant@bogon laravel-app]$ cat .env | grep DB_
DB_CONNECTION=mysql
# DB_HOST=127.0.0.1
DB_HOST=172.17.0.8
DB_PORT=3306
# DB_DATABASE=homestead
DB_DATABASE=rival
# DB_USERNAME=homestead
DB_USERNAME=root
DB_PASSWORD=password
```

Create DB
```
[vagrant@bogon laravel-app]$ docker exec -ti demo bash
[root@8bfbf42555c3 /]# mysql -u root -p
Enter password: 
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 3
Server version: 5.7.19 MySQL Community Server (GPL)

Copyright (c) 2000, 2017, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> show database;
ERROR 1064 (42000): You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near 'database' at line 1
mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sys                |
| testdb             |
+--------------------+
5 rows in set (0.03 sec)

mysql> create DATABASE rival;
Query OK, 1 row affected (0.00 sec)

mysql> quit
Bye
[root@8bfbf42555c3 /]# exit
exit
```

__Successful__
```
[vagrant@bogon laravel-app]$ php artisan migrate
Migration table created successfully.
Migrating: 2017_09_11_180644_create_battlefield_table
Migrated:  2017_09_11_180644_create_battlefield_table
Migrating: 2017_09_11_181043_add_stuff_to_battlefield_table
Migrated:  2017_09_11_181043_add_stuff_to_battlefield_table
```

## Connection configuration

config/database.php
```
        'mysql_dev' => [
            'read' => [
                'host' => '172.17.0.8',
            ],
            'write' => [
                'host' => '172.17.0.8'
            ],
            'port' => '3306',
            'driver'    => 'mysql',
            'database'  => 'database',
            'username'  => 'root',
            'password'  => 'password',
            'charset' => 'utf8mb4',
            'collation' => 'utf8mb4_unicode_ci',
            'prefix'    => '',
        ],
```

### Model

__Issue__

如果发生以下错误，是因为laravel没有完整安装，确认`composer update`返回成功
Reason：refer to [create project](./laravel-01-create-project.md)
```
[vagrant@bogon laravel-app]$ php artisan make:model Battlefield
PHP Warning:  require(/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/examples/php/rival/laravel-app/bootstrap/../vendor/autoload.php): failed to open stream: No such file or directory in /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/examples/php/rival/laravel-app/bootstrap/autoload.php on line 17
PHP Fatal error:  require(): Failed opening required '/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/examples/php/rival/laravel-app/bootstrap/../vendor/autoload.php' (include_path='.:/usr/share/pear:/usr/share/php') in /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/examples/php/rival/laravel-app/bootstrap/autoload.php on line 17
```

Model文件
```
[vagrant@bogon laravel-app]$ php artisan make:model Battlefield
Model created successfully.
[vagrant@bogon laravel-app]$ ls app/Battlefield.php 
app/Battlefield.php
```

## Reference

https://laravel.com/docs/5.4/database

https://laravel.com/docs/5.4/migrations

https://laravel.com/docs/5.4/eloquent

https://scotch.io/tutorials/a-guide-to-using-eloquent-orm-in-laravel

https://tutorials.kode-blog.com/laravel-5-eloquent-orm


