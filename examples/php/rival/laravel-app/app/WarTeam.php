// [vagrant@bogon laravel-app]$ php artisan make:model -m WarTeam
// Model created successfully.
// Created Migration: 2017_09_12_134134_create_war_teams_table

<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class WarTeam extends Model
{
    //

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'name',
        'logo',
        'background',
        'teamname',
        'organization',
        'strengths',
        'teamlogo',
        'introduction',
        'status',
        'create_time',
        'token',
        'ip_part',
        'province',
    ];

    /**
     * Get the warrior for the WarTeam.
     */
    public function warrior()
    {
        return $this->hasMany('App\Warrior');
    }
}
