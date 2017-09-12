
### Predis

__Test__

Subscriber
```
[vagrant@localhost rival]$ php pubsubredis_consumer.php 
Subscribed to control_channel
Subscribed to notifications
```

Publisher
```
[vagrant@bogon rival]$ docker exec -ti gitlab_redis_1 bash
root@fe0ae27989c4:/# redis-cli PUBLISH notifications "this is a test"    
(integer) 1
root@fe0ae27989c4:/# redis-cli PUBLISH notifications "this is a test too"
(integer) 1
root@fe0ae27989c4:/# redis-cli PUBLISH control_channel quit_loop
(integer) 1
```

Subscription
```
[vagrant@localhost rival]$ php pubsubredis_consumer.php 
Subscribed to control_channel
Subscribed to notifications
Received the following message from notifications:
  this is a test

Received the following message from notifications:
  this is a test too

Aborting pubsub loop...
Goodbye from Redis 2.8.4!
```

### Redis

https://redis.io/topics/rediscli#showing-help-about-redis-commands

* help @<category> shows all the commands about a given category. The categories are: @generic, @list, @set, @sorted_set, @hash, @pubsub, @transactions, @connection, @server, @scripting, @hyperloglog.
* help <commandname> shows specific help for the command given as argument.

pubsub
```
root@fe0ae27989c4:/# redis-cli help @pubsub  

  PSUBSCRIBE pattern [pattern ...]
  summary: Listen for messages published to channels matching the given patterns
  since: 2.0.0

  PUBLISH channel message
  summary: Post a message to a channel
  since: 2.0.0

  PUNSUBSCRIBE [pattern [pattern ...]]
  summary: Stop listening for messages posted to channels matching the given patterns
  since: 2.0.0

  SUBSCRIBE channel [channel ...]
  summary: Listen for messages published to the given channels
  since: 2.0.0

  UNSUBSCRIBE [channel [channel ...]]
  summary: Stop listening for messages posted to the given channels
  since: 2.0.0
```

string
```
  APPEND key value
  summary: Append a value to a key
  since: 2.0.0

  BITCOUNT key [start] [end]
  summary: Count set bits in a string
  since: 2.6.0

  BITOP operation destkey key [key ...]
  summary: Perform bitwise operations between strings
  since: 2.6.0

  DECR key
  summary: Decrement the integer value of a key by one
  since: 1.0.0

  DECRBY key decrement
  summary: Decrement the integer value of a key by the given number
  since: 1.0.0

  GET key
  summary: Get the value of a key
  since: 1.0.0

  GETBIT key offset
  summary: Returns the bit value at offset in the string value stored at key
  since: 2.2.0

  GETRANGE key start end
  summary: Get a substring of the string stored at a key
  since: 2.4.0

  GETSET key value
  summary: Set the string value of a key and return its old value
  since: 1.0.0

  INCR key
  summary: Increment the integer value of a key by one
  since: 1.0.0

  INCRBY key increment
  summary: Increment the integer value of a key by the given amount
  since: 1.0.0

  INCRBYFLOAT key increment
  summary: Increment the float value of a key by the given amount
  since: 2.6.0

  MGET key [key ...]
  summary: Get the values of all the given keys
  since: 1.0.0

  MSET key value [key value ...]
  summary: Set multiple keys to multiple values
  since: 1.0.1

  MSETNX key value [key value ...]
  summary: Set multiple keys to multiple values, only if none of the keys exist
  since: 1.0.1

  PSETEX key milliseconds value
  summary: Set the value and expiration in milliseconds of a key
  since: 2.6.0

  SET key value [EX seconds] [PX milliseconds] [NX|XX]
  summary: Set the string value of a key
  since: 1.0.0

  SETBIT key offset value
  summary: Sets or clears the bit at offset in the string value stored at key
  since: 2.2.0

  SETEX key seconds value
  summary: Set the value and expiration of a key
  since: 2.0.0

  SETNX key value
  summary: Set the value of a key, only if the key does not exist
  since: 1.0.0

  SETRANGE key offset value
  summary: Overwrite part of a string at key starting at the specified offset
  since: 2.2.0

  STRLEN key
  summary: Get the length of the value stored in a key
  since: 2.2.0
```

server
```
root@fe0ae27989c4:/# redis-cli help @server

  BGREWRITEAOF -
  summary: Asynchronously rewrite the append-only file
  since: 1.0.0

  BGSAVE -
  summary: Asynchronously save the dataset to disk
  since: 1.0.0

  CLIENT GETNAME -
  summary: Get the current connection name
  since: 2.6.9

  CLIENT KILL ip:port
  summary: Kill the connection of a client
  since: 2.4.0

  CLIENT LIST -
  summary: Get the list of client connections
  since: 2.4.0

  CLIENT SETNAME connection-name
  summary: Set the current connection name
  since: 2.6.9

  CONFIG GET parameter
  summary: Get the value of a configuration parameter
  since: 2.0.0

  CONFIG RESETSTAT -
  summary: Reset the stats returned by INFO
  since: 2.0.0

  CONFIG SET parameter value
  summary: Set a configuration parameter to the given value
  since: 2.0.0

  DBSIZE -
  summary: Return the number of keys in the selected database
  since: 1.0.0

  DEBUG OBJECT key
  summary: Get debugging information about a key
  since: 1.0.0

  DEBUG SEGFAULT -
  summary: Make the server crash
  since: 1.0.0

  FLUSHALL -
  summary: Remove all keys from all databases
  since: 1.0.0

  FLUSHDB -
  summary: Remove all keys from the current database
  since: 1.0.0

  INFO [section]
  summary: Get information and statistics about the server
  since: 1.0.0

  LASTSAVE -
  summary: Get the UNIX time stamp of the last successful save to disk
  since: 1.0.0

  MONITOR -
  summary: Listen for all requests received by the server in real time
  since: 1.0.0

  SAVE -
  summary: Synchronously save the dataset to disk
  since: 1.0.0

  SHUTDOWN [NOSAVE] [SAVE]
  summary: Synchronously save the dataset to disk and then shut down the server
  since: 1.0.0

  SLAVEOF host port
  summary: Make the server a slave of another instance, or promote it as master
  since: 1.0.0

  SLOWLOG subcommand [argument]
  summary: Manages the Redis slow queries log
  since: 2.2.12

  SYNC -
  summary: Internal command used for replication
  since: 1.0.0

  TIME -
  summary: Return the current server time
  since: 2.6.0
```

### workerman with thinkphp

__Issue__

Workerman\posix_getpid()
```
PHP Fatal error:  Call to undefined function Workerman\posix_getpid() in /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/examples/php/rival/think/vendor/workerman/workerman/Worker.php on line 1380

Fatal error: Call to undefined function Workerman\posix_getpid() in /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/examples/php/rival/think/vendor/workerman/workerman/Worker.php on line 1380


                                                       
  [think\exception\ErrorException]                     
  Call to undefined function Workerman\posix_getpid()  
```

Fedora 23
```
[vagrant@bogon think]$ sudo dnf list | grep  ^php-pr
php-process.x86_64                       5.6.29-1.fc23                   updates

[vagrant@bogon think]$ sudo dnf install php-process
上次元数据过期检查在 0:28:52 前执行于 Fri Sep  8 03:29:59 2017。
依赖关系解决。
==========================================================================================================================================================
 Package                               架构                             版本                                      仓库                               大小
==========================================================================================================================================================
安装:
 php-process                           x86_64                           5.6.29-1.fc23                             updates                            88 k

事务概要
==========================================================================================================================================================
安装  1 Package

总下载：88 k
安装大小：187 k
确定吗？[y/N]： y
下载软件包：
php-process-5.6.29-1.fc23.x86_64.rpm                                                                                      100 kB/s |  88 kB     00:00    
----------------------------------------------------------------------------------------------------------------------------------------------------------
总计                                                                                                                       35 kB/s |  88 kB     00:02     
运行事务检查
事务检查成功。
运行事务测试
事务测试成功。
运行事务
  安装: php-process-5.6.29-1.fc23.x86_64                                                                                                              1/1 
  验证: php-process-5.6.29-1.fc23.x86_64                                                                                                              1/1 

已安装:
  php-process.x86_64 5.6.29-1.fc23                                                                                                                        

完毕！
```

__Test__

Server
```
[vagrant@bogon think]$ php server.php start
Workerman[server.php] start in DEBUG mode
----------------------- WORKERMAN -----------------------------
Workerman version:3.4.5          PHP version:5.6.29
------------------------ WORKERS -------------------------------
user          worker        listen               processes status
vagrant       none          http://0.0.0.0:2346   4         [OK] 
----------------------------------------------------------------
Press Ctrl-C to quit. Start success.
```

Client
```
{"get":[],"post":[],"cookie":[],"server":{"QUERY_STRING":"","REQUEST_METHOD":"GET","REQUEST_URI":"\/","SERVER_PROTOCOL":"HTTP\/1.1","SERVER_SOFTWARE":"workerman\/3.4.5","SERVER_NAME":"localhost","HTTP_HOST":"localhost:2346","HTTP_USER_AGENT":"curl\/7.43.0","HTTP_ACCEPT":"*\/*","HTTP_ACCEPT_LANGUAGE":"","HTTP_ACCEPT_ENCODING":"","HTTP_COOKIE":"","HTTP_CONNECTION":"","REMOTE_ADDR":"127.0.0.1","REMOTE_PORT":46848,"REQUEST_TIME":1504843269,"SERVER_PORT":"2346"},"files":[]}
```