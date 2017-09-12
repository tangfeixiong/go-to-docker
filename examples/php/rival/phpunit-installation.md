Refer to __https://phpunit.de/__ for convenient version

Work dir
```
[vagrant@bogon laravel-app]$ pushd ~/
~ /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/examples/php/rival/laravel-app
```

Bad version
```
[vagrant@bogon ~]$ sudo wget https://phar.phpunit.de/phpunit.phar -O /opt/php/phpunit.phar
--2017-09-12 06:10:54--  https://phar.phpunit.de/phpunit.phar
正在解析主机 phar.phpunit.de (phar.phpunit.de)... 188.94.27.25
正在连接 phar.phpunit.de (phar.phpunit.de)|188.94.27.25|:443... 已连接。
已发出 HTTP 请求，正在等待回应... 302 Moved Temporarily
位置：https://phar.phpunit.de/phpunit-6.3.0.phar [跟随至新的 URL]
--2017-09-12 06:10:56--  https://phar.phpunit.de/phpunit-6.3.0.phar
再次使用存在的到 phar.phpunit.de:443 的连接。
已发出 HTTP 请求，正在等待回应... 200 OK
长度：2724060 (2.6M) [application/octet-stream]
正在保存至: “/opt/php/phpunit.phar”

/opt/php/phpunit.phar                  100%[==========================================================================>]   2.60M   123KB/s    in 22s     

2017-09-12 06:11:20 (123 KB/s) - 已保存 “/opt/php/phpunit.phar” [2724060/2724060])

[vagrant@bogon ~]$ php /opt/php/phpunit.phar --version
PHPUnit 6.3.0 by Sebastian Bergmann and contributors.

This version of PHPUnit is supported on PHP 7.0 and PHP 7.1.
You are using PHP 5.6.29 (/usr/bin/php).
[vagrant@bogon ~]$ php /opt/php/phpunit.phar --help
PHPUnit 6.3.0 by Sebastian Bergmann and contributors.

This version of PHPUnit is supported on PHP 7.0 and PHP 7.1.
You are using PHP 5.6.29 (/usr/bin/php).
```

Correct version
```
[vagrant@bogon ~]$ sudo wget https://phar.phpunit.de/phpunit-5.7.phar -O /opt/php/phpunit.phar
--2017-09-12 06:16:47--  https://phar.phpunit.de/phpunit-5.7.phar
正在解析主机 phar.phpunit.de (phar.phpunit.de)... 188.94.27.25
正在连接 phar.phpunit.de (phar.phpunit.de)|188.94.27.25|:443... 已连接。
已发出 HTTP 请求，正在等待回应... 302 Moved Temporarily
位置：https://phar.phpunit.de/phpunit-5.7.21.phar [跟随至新的 URL]
--2017-09-12 06:16:49--  https://phar.phpunit.de/phpunit-5.7.21.phar
再次使用存在的到 phar.phpunit.de:443 的连接。
已发出 HTTP 请求，正在等待回应... 200 OK
长度：3006816 (2.9M) [application/octet-stream]
正在保存至: “/opt/php/phpunit.phar”

/opt/php/phpunit.phar                  100%[==========================================================================>]   2.87M   499KB/s    in 9.5s    

2017-09-12 06:17:00 (309 KB/s) - 已保存 “/opt/php/phpunit.phar” [3006816/3006816])

[vagrant@bogon ~]$ php /opt/php/phpunit.phar --version
PHPUnit 5.7.21 by Sebastian Bergmann and contributors.

[vagrant@bogon ~]$ php /opt/php/phpunit.phar --help
PHPUnit 5.7.21 by Sebastian Bergmann and contributors.

Usage: phpunit [options] UnitTest [UnitTest.php]
       phpunit [options] <directory>

Code Coverage Options:

  --coverage-clover <file>  Generate code coverage report in Clover XML format.
  --coverage-crap4j <file>  Generate code coverage report in Crap4J XML format.
  --coverage-html <dir>     Generate code coverage report in HTML format.
  --coverage-php <file>     Export PHP_CodeCoverage object to file.
  --coverage-text=<file>    Generate code coverage report in text format.
                            Default: Standard output.
  --coverage-xml <dir>      Generate code coverage report in PHPUnit XML format.
  --whitelist <dir>         Whitelist <dir> for code coverage analysis.
  --disable-coverage-ignore Disable annotations for ignoring code coverage.

Logging Options:

  --log-junit <file>        Log test execution in JUnit XML format to file.
  --log-teamcity <file>     Log test execution in TeamCity format to file.
  --testdox-html <file>     Write agile documentation in HTML format to file.
  --testdox-text <file>     Write agile documentation in Text format to file.
  --testdox-xml <file>      Write agile documentation in XML format to file.
  --reverse-list            Print defects in reverse order

Test Selection Options:

  --filter <pattern>        Filter which tests to run.
  --testsuite <name>        Filter which testsuite to run.
  --group ...               Only runs tests from the specified group(s).
  --exclude-group ...       Exclude tests from the specified group(s).
  --list-groups             List available test groups.
  --list-suites             List available test suites.
  --test-suffix ...         Only search for test in files with specified
                            suffix(es). Default: Test.php,.phpt

Test Execution Options:

  --report-useless-tests    Be strict about tests that do not test anything.
  --strict-coverage         Be strict about @covers annotation usage.
  --strict-global-state     Be strict about changes to global state
  --disallow-test-output    Be strict about output during tests.
  --disallow-resource-usage Be strict about resource usage during small tests.
  --enforce-time-limit      Enforce time limit based on test size.
  --disallow-todo-tests     Disallow @todo-annotated tests.

  --process-isolation       Run each test in a separate PHP process.
  --no-globals-backup       Do not backup and restore $GLOBALS for each test.
  --static-backup           Backup and restore static attributes for each test.

  --colors=<flag>           Use colors in output ("never", "auto" or "always").
  --columns <n>             Number of columns to use for progress output.
  --columns max             Use maximum number of columns for progress output.
  --stderr                  Write to STDERR instead of STDOUT.
  --stop-on-error           Stop execution upon first error.
  --stop-on-failure         Stop execution upon first error or failure.
  --stop-on-warning         Stop execution upon first warning.
  --stop-on-risky           Stop execution upon first risky test.
  --stop-on-skipped         Stop execution upon first skipped test.
  --stop-on-incomplete      Stop execution upon first incomplete test.
  --fail-on-warning         Treat tests with warnings as failures.
  --fail-on-risky           Treat risky tests as failures.
  -v|--verbose              Output more verbose information.
  --debug                   Display debugging information during test execution.

  --loader <loader>         TestSuiteLoader implementation to use.
  --repeat <times>          Runs the test(s) repeatedly.
  --teamcity                Report test execution progress in TeamCity format.
  --testdox                 Report test execution progress in TestDox format.
  --testdox-group           Only include tests from the specified group(s).
  --testdox-exclude-group   Exclude tests from the specified group(s).
  --printer <printer>       TestListener implementation to use.

Configuration Options:

  --bootstrap <file>        A "bootstrap" PHP file that is run before the tests.
  -c|--configuration <file> Read configuration from XML file.
  --no-configuration        Ignore default configuration file (phpunit.xml).
  --no-coverage             Ignore code coverage configuration.
  --no-extensions           Do not load PHPUnit extensions.
  --include-path <path(s)>  Prepend PHP's include_path with given path(s).
  -d key[=value]            Sets a php.ini value.
  --generate-configuration  Generate configuration file with suggested settings.

Miscellaneous Options:

  -h|--help                 Prints this usage information.
  --version                 Prints the version and exits.
  --atleast-version <min>   Checks that version is greater than min and exits.

  --check-version           Check whether PHPUnit is the latest version.
```

Back dir
```
[vagrant@bogon ~]$ popd
/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/examples/php/rival/laravel-app
```
