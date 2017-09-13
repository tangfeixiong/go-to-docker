// [vagrant@bogon laravel-app]$ php artisan make:model -m WarCamp
// Model created successfully.
// Created Migration: 2017_09_13_000737_create_war_camps_table

<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class WarCamp extends Model
{
    //

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'warrior_id',
        'networking_id',
        'ipv4',
    ];

    /**
     * Get the war_team that owns the warrior.
     */
    public function warrior()
    {
        return $this->belongsTo('App\Warrior');
    }
}
