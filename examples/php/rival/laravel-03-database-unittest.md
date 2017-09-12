# Database Unit Test

## Model and Repository

Model
```
[vagrant@bogon laravel-app]$ ls app/Battlefield.php 
app/Battlefield.php
```

Repository
```
[vagrant@bogon laravel-app]$ ls app/Repository/
BattlefieldRepository.php
```

## Testing

Faker and Model Factory
```
[vagrant@bogon laravel-app]$ ls database/factories/
ModelFactory.php
```

Test Case
```
[vagrant@bogon laravel-app]$ ls tests/Unit/Repository/
BattlefieldTest.php
```
Using `artisan`
```
php artisan make:test NetworkingTest --unit
```
### Unit Test

Test
```
[vagrant@bogon laravel-app]$ php /opt/php/phpunit.phar UnitTest ./tests/Unit/Repository/BattlefieldTest.php 
PHPUnit 5.7.21 by Sebastian Bergmann and contributors.

.3.                                                                  2 / 2 (100%)5

Time: 2.69 seconds, Memory: 54.00MB

OK (2 tests, 3 assertions)
```

MySQL cli
```
[vagrant@bogon laravel-app]$ docker exec -ti demo bash
[root@8bfbf42555c3 /]# mysql -u root -p
Enter password: 
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 17
Server version: 5.7.19 MySQL Community Server (GPL)

Copyright (c) 2000, 2017, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> use rival;
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A

Database changed
mysql> select * from battlefield;
+----+--------------------------+---------------------+------------+-----------+---------+------------+--------------+
| id | name                     | description         | match_item | status    | payment | created_at | last_updated |
+----+--------------------------+---------------------+------------+-----------+---------+------------+--------------+
|  1 | Edgardo Hoeger           | Alice indignantly.. | odio       | not-ready | beta    | NULL       | NULL         |
|  2 | Miss Susanna Hirthe      | SHE, of course,'.   | accusamus  | not-ready | beta    | NULL       | NULL         |
|  3 | Dr. Casey Hackett        | KH6Gw52K5t          | d5qYL1Jfqf | ullam     | I vote. | NULL       | NULL         |
|  4 | Opal Howe                | F0DMHrwcJu          | IFbkkJapK0 | on-ready  | I'm.    | NULL       | NULL         |
|  5 | Mr. Emerald Konopelski I | Mock Turtle said:.  | voluptatem | not-ready | beta    | NULL       | NULL         |
|  6 | Gilda Brakus             | I was going to.     | enim       | not-ready | beta    | NULL       | NULL         |
+----+--------------------------+---------------------+------------+-----------+---------+------------+--------------+
6 rows in set (0.00 sec)

mysql> exit
Bye
[root@8bfbf42555c3 /]# exit
exit
```

## Reference

https://laravel.com/docs/5.4/database-testing

https://scotch.io/tutorials/generate-dummy-laravel-data-with-model-factories

https://github.com/fzaninotto/Faker#formatters

https://semaphoreci.com/community/tutorials/getting-started-with-phpunit-in-laravel

