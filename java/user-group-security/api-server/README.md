# Development

## Test

### Run

MariaDB server, for Docker example
```bash
[vagrant@kubedev-172-17-4-59 user-group-security]$ docker inspect mariadb -f {{.NetworkSettings.Networks.user_organization_security.IPAddress}}
172.19.0.2
```

Maven
```
[vagrant@kubedev-172-17-4-59 user-group-security]$ mvn package spring-boot:run -Dmaven.test.skip=true -Dspring.profiles.active=dev -Drun.arguments=--spring.datasource.url=jdbc:mysql://172.19.0.2:3306/example,--spring.datasource.username=dbuser,--spring.datasource.password=dbpass --projects api-server
```

Curl
```
fanhonglingdeMacBook-Pro:go-to-docker fanhongling$ TEST_HOST=172.17.4.59:8090 ./java/user-group-security/api-server/runtests_curl.sh addone
* Hostname was NOT found in DNS cache
*   Trying 172.17.4.59...
* Connected to 172.17.4.59 (172.17.4.59) port 8090 (#0)
> POST //v1/default/users HTTP/1.1
> User-Agent: curl/7.37.1
> Host: 172.17.4.59:8090
> Accept: */*
> Content-Type: application/json
> Content-Length: 49
> 
* upload completely sent off: 49 out of 49 bytes
< HTTP/1.1 200 
HTTP/1.1 200 
< Content-Type: application/json;charset=UTF-8
Content-Type: application/json;charset=UTF-8
< Transfer-Encoding: chunked
Transfer-Encoding: chunked
< Date: Tue, 19 Jun 2018 05:38:33 GMT
Date: Tue, 19 Jun 2018 05:38:33 GMT

< 
* Connection #0 to host 172.17.4.59 left intact
100
```

Mysql
```
[vagrant@kubedev-172-17-4-59 user-group-security]$ docker exec -ti mariadb mysql -u dbuser --password=dbpass --database=example --execute="select * from user"
+-----+----------+----------+-------+--------+--------+
| id  | username | password | email | mobile | status |
+-----+----------+----------+-------+--------+--------+
| 100 | admin    | 123456   |       |        |      0 |
+-----+----------+----------+-------+--------+--------+
```


```bash
fanhonglingdeMacBook-Pro:user-group-security fanhongling$ mvn package spring-boot:run -Dmaven.test.skip=true -Dspring.profiles.active=dev -Dspring.datasource.url=jdbc:mysql://172.17.4.59:3306/example -Dspring.datasource.username=dbuser -Dspring.datasource.password=dbpass --projects api-server
```

```
[vagrant@kubedev-172-17-4-59 user-group-security]$ TEST_HOST=192.168.0.106:8090 ./api-server/runtests_curl.sh withdrawuser
```


For example, skip all tests (http://maven.apache.org/plugins-archives/maven-surefire-plugin-2.12.4/examples/skipping-test.html)
```
[vagrant@kubedev-172-17-4-59 user-group-security]$ mvn package spring-boot:repackage --Dmaven.test.skip=true --projects api-server
```


```
[vagrant@kubedev-172-17-4-59 user-group-security]$ api-server/build-ship-run.sh 
```

Core
```
[vagrant@kubedev-172-17-4-59 user-group-security]$ mvn install --no-snapshot-updates --projects core
```

