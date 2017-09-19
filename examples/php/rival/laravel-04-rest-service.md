# ReSTful service

__Routes__

```
[vagrant@bogon laravel-app]$ ls routes/api.php 
routes/api.php
```

__Controller__

```
[vagrant@bogon laravel-app]$ php artisan make:controller BattlefiledRestController
Controller created successfully.
```

move
```
[vagrant@bogon laravel-app]$ ls app/Http/Controllers/ReSTful/BattlefiledRestController.php 
app/Http/Controllers/ReSTful/BattlefiledRestController.php
```

others
```
[vagrant@bogon laravel-app]$ php artisan make:controller WarTeamRepoController --resource
Controller created successfully.
[vagrant@bogon laravel-app]$ php artisan make:controller WarriorRepoController --resource --model=Warrior
Controller created successfully.

```





## Reference

https://laravel.com/docs/5.4/routing
https://laravel.com/docs/5.4/controllers
https://laravel.com/docs/5.4/requests

https://www.toptal.com/laravel/restful-laravel-api-tutorial


