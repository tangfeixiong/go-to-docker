// [vagrant@bogon laravel-app]$ php artisan make:model -m WarCamp
// Model created successfully.
// Created Migration: 2017_09_13_000737_create_war_camps_table

<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateWarCampsTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('war_camp', function (Blueprint $table) {
            // $table->increments('id');
            // $table->timestamps();

            $table->bigInteger('warrior_id')->comment('id of warrior table');
            $table->bigInteger('war_team_id')->comment('id of war_team table');
            $table->ipAddress('ipv4');
            $table->macAddress('mac')->nullable()->comment('MAC address');
            
            $table->timestampsTz();
            // $table->softDeletes();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('war_camp');
    }
}
