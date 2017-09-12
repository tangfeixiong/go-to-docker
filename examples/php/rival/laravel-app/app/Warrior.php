// [vagrant@bogon laravel-app]$ php artisan make:model -m Warrior
// Model created successfully.
// Created Migration: 2017_09_12_141429_create_warriors_table

<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class Warrior extends Model
{
    //

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'war_team_id',
        'name',
        'avatar',
        'title',
        'position',
        'rank',
        'background',
    ];

    /**
     * Get the war_team that owns the warrior.
     */
    public function warTeam()
    {
        return $this->belongsTo('App\WarTeam');
    }

}
