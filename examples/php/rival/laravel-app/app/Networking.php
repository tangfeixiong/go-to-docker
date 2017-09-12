// [vagrant@bogon laravel-app]$ php artisan make:model -m Networking
// Model created successfully.
// Created Migration: 2017_09_12_122928_create_networkings_table

<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class Networking extends Model
{
    //

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = ['battlefield_id', 'name', 'gateway', 'subnets', 'bridge', 'dns', 'host_routes', 'infra_type'];
}
