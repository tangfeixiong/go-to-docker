# Laravel installation and creating project

PHP
```
[vagrant@bogon go-to-docker]$ php --version
PHP 5.6.29 (cli) (built: Dec  8 2016 09:19:46) 
Copyright (c) 1997-2016 The PHP Group
Zend Engine v2.6.0, Copyright (c) 1998-2016 Zend Technologies
```

Laravel 5.4
* PHP >= 5.6.4
* OpenSSL PHP Extension
* PDO PHP Extension
* Mbstring PHP Extension
* Tokenizer PHP Extension
* XML PHP Extension

## Table of contents

Create Project named _laravel-app_

完整安装laravel 5.4
* Env文件

Reference

## 创建项目

__issue__

缺少php-mbstring组件
```
[vagrant@bogon rival]$ php /opt/php/composer.phar create-project --prefer-dist laravel/laravel laravel-app "5.4.*"
Installing laravel/laravel (v5.4.30)
  - Installing laravel/laravel (v5.4.30): Downloading (100%)         
Created project in laravel-app
> php -r "file_exists('.env') || copy('.env.example', '.env');"
Loading composer repositories with package information
Updating dependencies (including require-dev)
Your requirements could not be resolved to an installable set of packages.

  Problem 1
    - laravel/framework v5.4.9 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.8 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.7 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.6 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.5 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.4 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.36 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.35 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.34 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.33 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.32 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.31 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.30 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.3 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.29 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.28 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.27 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.26 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.25 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.24 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.23 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.22 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.21 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.20 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.2 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.19 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.18 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.17 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.16 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.15 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.14 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.13 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.12 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.11 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.10 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.1 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - laravel/framework v5.4.0 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
    - Installation request for laravel/framework 5.4.* -> satisfiable by laravel/framework[v5.4.0, v5.4.1, v5.4.10, v5.4.11, v5.4.12, v5.4.13, v5.4.14, v5.4.15, v5.4.16, v5.4.17, v5.4.18, v5.4.19, v5.4.2, v5.4.20, v5.4.21, v5.4.22, v5.4.23, v5.4.24, v5.4.25, v5.4.26, v5.4.27, v5.4.28, v5.4.29, v5.4.3, v5.4.30, v5.4.31, v5.4.32, v5.4.33, v5.4.34, v5.4.35, v5.4.36, v5.4.4, v5.4.5, v5.4.6, v5.4.7, v5.4.8, v5.4.9].

  To enable extensions, verify that they are enabled in your .ini files:
    - /etc/php.ini
    - /etc/php.d/20-bz2.ini
    - /etc/php.d/20-calendar.ini
    - /etc/php.d/20-ctype.ini
    - /etc/php.d/20-curl.ini
    - /etc/php.d/20-exif.ini
    - /etc/php.d/20-fileinfo.ini
    - /etc/php.d/20-ftp.ini
    - /etc/php.d/20-gettext.ini
    - /etc/php.d/20-iconv.ini
    - /etc/php.d/20-phar.ini
    - /etc/php.d/20-posix.ini
    - /etc/php.d/20-shmop.ini
    - /etc/php.d/20-sockets.ini
    - /etc/php.d/20-sysvmsg.ini
    - /etc/php.d/20-sysvsem.ini
    - /etc/php.d/20-sysvshm.ini
    - /etc/php.d/20-tokenizer.ini
    - /etc/php.d/40-json.ini
  You can also run `php --ini` inside terminal to see which files are used by PHP in CLI mode.
```

查找linux package（yum，apt，apk）

__Fedora32__ 
```
[vagrant@bogon rival]$ sudo dnf search php-mbstring
上次元数据过期检查在 1:23:12 前执行于 Mon Sep 11 13:45:25 2017。
================================================================= N/S 匹配：php-mbstring =================================================================
php-mbstring.x86_64 : A module for PHP applications which need multi-byte string handling
```

Installation
```
[vagrant@bogon rival]$ sudo dnf install php-mbstring
上次元数据过期检查在 1:27:29 前执行于 Mon Sep 11 13:45:25 2017。
依赖关系解决。
==========================================================================================================================================================
 Package                                架构                             版本                                     仓库                               大小
==========================================================================================================================================================
安装:
 php-mbstring                           x86_64                           5.6.29-1.fc23                            updates                           584 k

事务概要
==========================================================================================================================================================
安装  1 Package

总下载：584 k
安装大小：2.8 M
确定吗？[y/N]： y
下载软件包：
php-mbstring-5.6.29-1.fc23.x86_64.rpm                                                                                     396 kB/s | 584 kB     00:01    
----------------------------------------------------------------------------------------------------------------------------------------------------------
总计                                                                                                                      164 kB/s | 584 kB     00:03     
运行事务检查
事务检查成功。
运行事务测试
事务测试成功。
运行事务
  安装: php-mbstring-5.6.29-1.fc23.x86_64                                                                                                             1/1 
  验证: php-mbstring-5.6.29-1.fc23.x86_64                                                                                                             1/1 

已安装:
  php-mbstring.x86_64 5.6.29-1.fc23                                                                                                                       

完毕！
```

Statisfy
```
[vagrant@bogon rival]$ php --ini
Configuration File (php.ini) Path: /etc
Loaded Configuration File:         /etc/php.ini
Scan for additional .ini files in: /etc/php.d
Additional .ini files parsed:      /etc/php.d/20-bz2.ini,
/etc/php.d/20-calendar.ini,
/etc/php.d/20-ctype.ini,
/etc/php.d/20-curl.ini,
/etc/php.d/20-exif.ini,
/etc/php.d/20-fileinfo.ini,
/etc/php.d/20-ftp.ini,
/etc/php.d/20-gettext.ini,
/etc/php.d/20-iconv.ini,
/etc/php.d/20-mbstring.ini,
/etc/php.d/20-phar.ini,
/etc/php.d/20-posix.ini,
/etc/php.d/20-shmop.ini,
/etc/php.d/20-sockets.ini,
/etc/php.d/20-sysvmsg.ini,
/etc/php.d/20-sysvsem.ini,
/etc/php.d/20-sysvshm.ini,
/etc/php.d/20-tokenizer.ini,
/etc/php.d/40-json.ini
```

__Issue__

没有继续执行composer update，而是
```
[vagrant@bogon laravel-app]$ php artisan serve
PHP Warning:  require(/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/examples/php/rival/laravel-app/bootstrap/../vendor/autoload.php): failed to open stream: No such file or directory in /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/examples/php/rival/laravel-app/bootstrap/autoload.php on line 17
PHP Fatal error:  require(): Failed opening required '/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/examples/php/rival/laravel-app/bootstrap/../vendor/autoload.php' (include_path='.:/usr/share/pear:/usr/share/php') in /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/examples/php/rival/laravel-app/bootstrap/autoload.php on line 17
```

完成执行composer update
```
[vagrant@bogon laravel-app]$ php /opt/php/composer.phar update
Loading composer repositories with package information
Updating dependencies (including require-dev)
Your requirements could not be resolved to an installable set of packages.

  Problem 1
    - phpunit/phpunit 5.7.9 requires ext-dom * -> the requested PHP extension dom is missing from your system.
    - phpunit/phpunit 5.7.8 requires ext-dom * -> the requested PHP extension dom is missing from your system.
    - phpunit/phpunit 5.7.7 requires ext-dom * -> the requested PHP extension dom is missing from your system.
    - phpunit/phpunit 5.7.6 requires ext-dom * -> the requested PHP extension dom is missing from your system.
    - phpunit/phpunit 5.7.5 requires ext-dom * -> the requested PHP extension dom is missing from your system.
    - phpunit/phpunit 5.7.4 requires ext-dom * -> the requested PHP extension dom is missing from your system.
    - phpunit/phpunit 5.7.3 requires ext-dom * -> the requested PHP extension dom is missing from your system.
    - phpunit/phpunit 5.7.21 requires ext-dom * -> the requested PHP extension dom is missing from your system.
    - phpunit/phpunit 5.7.20 requires ext-dom * -> the requested PHP extension dom is missing from your system.
    - phpunit/phpunit 5.7.2 requires ext-dom * -> the requested PHP extension dom is missing from your system.
    - phpunit/phpunit 5.7.19 requires ext-dom * -> the requested PHP extension dom is missing from your system.
    - phpunit/phpunit 5.7.18 requires ext-dom * -> the requested PHP extension dom is missing from your system.
    - phpunit/phpunit 5.7.17 requires ext-dom * -> the requested PHP extension dom is missing from your system.
    - phpunit/phpunit 5.7.16 requires ext-dom * -> the requested PHP extension dom is missing from your system.
    - phpunit/phpunit 5.7.15 requires ext-dom * -> the requested PHP extension dom is missing from your system.
    - phpunit/phpunit 5.7.14 requires ext-dom * -> the requested PHP extension dom is missing from your system.
    - phpunit/phpunit 5.7.13 requires ext-dom * -> the requested PHP extension dom is missing from your system.
    - phpunit/phpunit 5.7.12 requires ext-dom * -> the requested PHP extension dom is missing from your system.
    - phpunit/phpunit 5.7.11 requires ext-dom * -> the requested PHP extension dom is missing from your system.
    - phpunit/phpunit 5.7.10 requires ext-dom * -> the requested PHP extension dom is missing from your system.
    - phpunit/phpunit 5.7.1 requires ext-dom * -> the requested PHP extension dom is missing from your system.
    - phpunit/phpunit 5.7.0 requires ext-dom * -> the requested PHP extension dom is missing from your system.
    - Installation request for phpunit/phpunit ~5.7 -> satisfiable by phpunit/phpunit[5.7.0, 5.7.1, 5.7.10, 5.7.11, 5.7.12, 5.7.13, 5.7.14, 5.7.15, 5.7.16, 5.7.17, 5.7.18, 5.7.19, 5.7.2, 5.7.20, 5.7.21, 5.7.3, 5.7.4, 5.7.5, 5.7.6, 5.7.7, 5.7.8, 5.7.9].

  To enable extensions, verify that they are enabled in your .ini files:
    - /etc/php.ini
    - /etc/php.d/20-bz2.ini
    - /etc/php.d/20-calendar.ini
    - /etc/php.d/20-ctype.ini
    - /etc/php.d/20-curl.ini
    - /etc/php.d/20-exif.ini
    - /etc/php.d/20-fileinfo.ini
    - /etc/php.d/20-ftp.ini
    - /etc/php.d/20-gettext.ini
    - /etc/php.d/20-iconv.ini
    - /etc/php.d/20-mbstring.ini
    - /etc/php.d/20-phar.ini
    - /etc/php.d/20-posix.ini
    - /etc/php.d/20-shmop.ini
    - /etc/php.d/20-sockets.ini
    - /etc/php.d/20-sysvmsg.ini
    - /etc/php.d/20-sysvsem.ini
    - /etc/php.d/20-sysvshm.ini
    - /etc/php.d/20-tokenizer.ini
    - /etc/php.d/40-json.ini
  You can also run `php --ini` inside terminal to see which files are used by PHP in CLI mode.
```

缺少组件

__Fedora32__
```
[vagrant@bogon laravel-app]$ sudo dnf search php-xml
上次元数据过期检查在 2:24:18 前执行于 Mon Sep 11 13:45:25 2017。
=================================================================== N/S 匹配：php-xml ====================================================================
php-xml.x86_64 : A module for PHP applications which use XML
php-xmlrpc.x86_64 : A module for PHP applications which use the XML-RPC protocol
php-xmlseclibs.noarch : PHP library for XML Security
```

Installation
```
[vagrant@bogon laravel-app]$ sudo dnf search php-xml
上次元数据过期检查在 2:24:18 前执行于 Mon Sep 11 13:45:25 2017。
=================================================================== N/S 匹配：php-xml ====================================================================
php-xml.x86_64 : A module for PHP applications which use XML
php-xmlrpc.x86_64 : A module for PHP applications which use the XML-RPC protocol
php-xmlseclibs.noarch : PHP library for XML Security
[vagrant@bogon laravel-app]$ sudo dnf install -y php-xml
上次元数据过期检查在 2:25:51 前执行于 Mon Sep 11 13:45:25 2017。
依赖关系解决。
==========================================================================================================================================================
 Package                            架构                              版本                                       仓库                                大小
==========================================================================================================================================================
安装:
 php-xml                            x86_64                            5.6.29-1.fc23                              updates                            263 k

事务概要
==========================================================================================================================================================
安装  1 Package

总下载：263 k
安装大小：886 k
下载软件包：
php-xml-5.6.29-1.fc23.x86_64.rpm                                                                                           30 kB/s | 263 kB     00:08    
----------------------------------------------------------------------------------------------------------------------------------------------------------
总计                                                                                                                       25 kB/s | 263 kB     00:10     
运行事务检查
事务检查成功。
运行事务测试
事务测试成功。
运行事务
  安装: php-xml-5.6.29-1.fc23.x86_64                                                                                                                  1/1 
  验证: php-xml-5.6.29-1.fc23.x86_64                                                                                                                  1/1 

已安装:
  php-xml.x86_64 5.6.29-1.fc23                                                                                                                            

完毕！
```

## 完整安装laravel 5.4

Satisfy
```
[vagrant@bogon laravel-app]$ php /opt/php/composer.phar update
Loading composer repositories with package information
Updating dependencies (including require-dev)
Package operations: 59 installs, 0 updates, 0 removals
  - Installing symfony/css-selector (v3.3.8): Downloading (100%)         
  - Installing tijsverkoyen/css-to-inline-styles (2.2.0): Downloading (100%)         
  - Installing doctrine/inflector (v1.1.0): Downloading (100%)         
  - Installing symfony/polyfill-mbstring (v1.5.0): Downloading (100%)         
  - Installing symfony/var-dumper (v3.3.8): Downloading (100%)         
  - Installing jakub-onderka/php-console-color (0.1): Downloading (100%)         
  - Installing jakub-onderka/php-console-highlighter (v0.3.2): Downloading (100%)         
  - Installing dnoegel/php-xdg-base-dir (0.1): Downloading (100%)         
  - Installing nikic/php-parser (v3.1.1): Downloading (100%)         
  - Installing psr/log (1.0.2): Downloading (100%)         
  - Installing symfony/debug (v3.3.8): Downloading (100%)         
  - Installing symfony/console (v3.3.8): Downloading (100%)         
  - Installing psy/psysh (v0.8.11): Downloading (100%)         
  - Installing vlucas/phpdotenv (v2.4.0): Downloading (100%)         
  - Installing symfony/routing (v3.3.8): Downloading (100%)         
  - Installing symfony/process (v3.3.8): Downloading (100%)         
  - Installing symfony/http-foundation (v3.3.8): Downloading (100%)         
  - Installing symfony/event-dispatcher (v3.3.8): Downloading (100%)         
  - Installing symfony/http-kernel (v3.3.8): Downloading (100%)         
  - Installing symfony/finder (v3.3.8): Downloading (100%)         
  - Installing swiftmailer/swiftmailer (v5.4.8): Downloading (100%)         
  - Installing paragonie/random_compat (v2.0.10): Downloading (100%)         
  - Installing ramsey/uuid (3.7.0): Downloading (100%)         
  - Installing symfony/translation (v3.3.8): Downloading (100%)         
  - Installing nesbot/carbon (1.22.1): Downloading (100%)         
  - Installing mtdowling/cron-expression (v1.2.0): Downloading (100%)         
  - Installing monolog/monolog (1.23.0): Downloading (100%)         
  - Installing league/flysystem (1.0.41): Downloading (100%)         
  - Installing erusev/parsedown (1.6.3): Downloading (100%)         
  - Installing laravel/framework (v5.4.36): Downloading (100%)         
  - Installing laravel/tinker (v1.0.2): Downloading (100%)         
  - Installing fzaninotto/faker (v1.7.1): Downloading (100%)         
  - Installing hamcrest/hamcrest-php (v1.2.2): Downloading (100%)         
  - Installing mockery/mockery (0.9.9): Downloading (100%)         
  - Installing symfony/yaml (v3.3.8): Downloading (100%)         
  - Installing sebastian/version (2.0.1): Downloading (100%)         
  - Installing sebastian/resource-operations (1.0.0): Downloading (100%)         
  - Installing sebastian/recursion-context (2.0.0): Downloading (100%)         
  - Installing sebastian/object-enumerator (2.0.1): Downloading (100%)         
  - Installing sebastian/global-state (1.1.1): Downloading (100%)         
  - Installing sebastian/exporter (2.0.0): Downloading (100%)         
  - Installing sebastian/environment (2.0.0): Downloading (100%)         
  - Installing sebastian/diff (1.4.3): Downloading (100%)         
  - Installing sebastian/comparator (1.2.4): Downloading (100%)         
  - Installing doctrine/instantiator (1.0.5): Downloading (100%)         
  - Installing phpunit/php-text-template (1.2.1): Downloading (100%)         
  - Installing phpunit/phpunit-mock-objects (3.4.4): Downloading (100%)         
  - Installing phpunit/php-timer (1.0.9): Downloading (100%)         
  - Installing phpunit/php-file-iterator (1.4.2): Downloading (100%)         
  - Installing sebastian/code-unit-reverse-lookup (1.0.1): Downloading (100%)         
  - Installing phpunit/php-token-stream (1.4.11): Downloading (100%)         
  - Installing phpunit/php-code-coverage (4.0.8): Downloading (100%)         
  - Installing webmozart/assert (1.2.0): Downloading (100%)         
  - Installing phpdocumentor/reflection-common (1.0): Downloading (100%)         
  - Installing phpdocumentor/type-resolver (0.3.0): Downloading (100%)         
  - Installing phpdocumentor/reflection-docblock (3.2.2): Downloading (100%)         
  - Installing phpspec/prophecy (v1.7.2): Downloading (100%)         
  - Installing myclabs/deep-copy (1.6.1): Downloading (100%)         
  - Installing phpunit/phpunit (5.7.21): Downloading (100%)         
symfony/var-dumper suggests installing ext-symfony_debug ()
symfony/console suggests installing symfony/filesystem ()
psy/psysh suggests installing ext-pdo-sqlite (The doc command requires SQLite to work.)
psy/psysh suggests installing hoa/console (A pure PHP readline implementation. You'll want this if your PHP install doesn't already support readline or libedit.)
symfony/routing suggests installing doctrine/annotations (For using the annotation loader)
symfony/routing suggests installing symfony/config (For using the all-in-one router or any loader)
symfony/routing suggests installing symfony/dependency-injection (For loading routes from a service)
symfony/routing suggests installing symfony/expression-language (For using expression matching)
symfony/event-dispatcher suggests installing symfony/dependency-injection ()
symfony/http-kernel suggests installing symfony/browser-kit ()
symfony/http-kernel suggests installing symfony/class-loader ()
symfony/http-kernel suggests installing symfony/config ()
symfony/http-kernel suggests installing symfony/dependency-injection ()
paragonie/random_compat suggests installing ext-libsodium (Provides a modern crypto API that can be used to generate random bytes.)
ramsey/uuid suggests installing ircmaxell/random-lib (Provides RandomLib for use with the RandomLibAdapter)
ramsey/uuid suggests installing ext-libsodium (Provides the PECL libsodium extension for use with the SodiumRandomGenerator)
ramsey/uuid suggests installing ext-uuid (Provides the PECL UUID extension for use with the PeclUuidTimeGenerator and PeclUuidRandomGenerator)
ramsey/uuid suggests installing moontoast/math (Provides support for converting UUID to 128-bit integer (in string form).)
ramsey/uuid suggests installing ramsey/uuid-doctrine (Allows the use of Ramsey\Uuid\Uuid as Doctrine field type.)
ramsey/uuid suggests installing ramsey/uuid-console (A console application for generating UUIDs with ramsey/uuid)
symfony/translation suggests installing symfony/config ()
monolog/monolog suggests installing aws/aws-sdk-php (Allow sending log messages to AWS services like DynamoDB)
monolog/monolog suggests installing doctrine/couchdb (Allow sending log messages to a CouchDB server)
monolog/monolog suggests installing ext-amqp (Allow sending log messages to an AMQP server (1.0+ required))
monolog/monolog suggests installing ext-mongo (Allow sending log messages to a MongoDB server)
monolog/monolog suggests installing graylog2/gelf-php (Allow sending log messages to a GrayLog2 server)
monolog/monolog suggests installing mongodb/mongodb (Allow sending log messages to a MongoDB server via PHP Driver)
monolog/monolog suggests installing php-amqplib/php-amqplib (Allow sending log messages to an AMQP server using php-amqplib)
monolog/monolog suggests installing php-console/php-console (Allow sending log messages to Google Chrome)
monolog/monolog suggests installing rollbar/rollbar (Allow sending log messages to Rollbar)
monolog/monolog suggests installing ruflin/elastica (Allow sending log messages to an Elastic Search server)
monolog/monolog suggests installing sentry/sentry (Allow sending log messages to a Sentry server)
league/flysystem suggests installing league/flysystem-aws-s3-v2 (Allows you to use S3 storage with AWS SDK v2)
league/flysystem suggests installing league/flysystem-aws-s3-v3 (Allows you to use S3 storage with AWS SDK v3)
league/flysystem suggests installing league/flysystem-azure (Allows you to use Windows Azure Blob storage)
league/flysystem suggests installing league/flysystem-cached-adapter (Flysystem adapter decorator for metadata caching)
league/flysystem suggests installing league/flysystem-eventable-filesystem (Allows you to use EventableFilesystem)
league/flysystem suggests installing league/flysystem-rackspace (Allows you to use Rackspace Cloud Files)
league/flysystem suggests installing league/flysystem-sftp (Allows you to use SFTP server storage via phpseclib)
league/flysystem suggests installing league/flysystem-webdav (Allows you to use WebDAV storage)
league/flysystem suggests installing league/flysystem-ziparchive (Allows you to use ZipArchive adapter)
league/flysystem suggests installing spatie/flysystem-dropbox (Allows you to use Dropbox storage)
league/flysystem suggests installing srmklive/flysystem-dropbox-v2 (Allows you to use Dropbox storage for PHP 5 applications)
laravel/framework suggests installing aws/aws-sdk-php (Required to use the SQS queue driver and SES mail driver (~3.0).)
laravel/framework suggests installing doctrine/dbal (Required to rename columns and drop SQLite columns (~2.5).)
laravel/framework suggests installing guzzlehttp/guzzle (Required to use the Mailgun and Mandrill mail drivers and the ping methods on schedules (~6.0).)
laravel/framework suggests installing league/flysystem-aws-s3-v3 (Required to use the Flysystem S3 driver (~1.0).)
laravel/framework suggests installing league/flysystem-rackspace (Required to use the Flysystem Rackspace driver (~1.0).)
laravel/framework suggests installing nexmo/client (Required to use the Nexmo transport (~1.0).)
laravel/framework suggests installing pda/pheanstalk (Required to use the beanstalk queue driver (~3.0).)
laravel/framework suggests installing predis/predis (Required to use the redis cache and queue drivers (~1.0).)
laravel/framework suggests installing pusher/pusher-php-server (Required to use the Pusher broadcast driver (~2.0).)
laravel/framework suggests installing symfony/dom-crawler (Required to use most of the crawler integration testing tools (~3.2).)
laravel/framework suggests installing symfony/psr-http-message-bridge (Required to psr7 bridging features (0.2.*).)
sebastian/global-state suggests installing ext-uopz (*)
phpunit/phpunit-mock-objects suggests installing ext-soap (*)
phpunit/php-code-coverage suggests installing ext-xdebug (^2.5.1)
phpunit/phpunit suggests installing phpunit/php-invoker (~1.1)
phpunit/phpunit suggests installing ext-xdebug (*)
Writing lock file
Generating optimized autoload files
> Illuminate\Foundation\ComposerScripts::postUpdate
> php artisan optimize
Generating optimized class loader
The compiled services file has been removed.
```

### Env

Project dir
```
[vagrant@bogon laravel-app]$ ls -a
.   app      bootstrap      config    .env          .gitattributes  package.json  public     resources  server.php  tests
..  artisan  composer.json  database  .env.example  .gitignore      phpunit.xml   readme.md  routes     storage     webpack.mix.js
```


## Reference

[Laravel 5.4 documentation](https://laravel.com/docs/5.4)

