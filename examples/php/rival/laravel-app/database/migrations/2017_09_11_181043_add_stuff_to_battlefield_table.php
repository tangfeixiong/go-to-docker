// [vagrant@bogon laravel-app]$ php artisan make:migration add_stuff_to_battlefield_table --table=battlefield
// Created Migration: 2017_09_11_181043_add_stuff_to_battlefield_table

<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class AddStuffToBattlefieldTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::table('battlefield', function (Blueprint $table) {
            //
            $table->text('theme')->nullable();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::table('battlefield', function (Blueprint $table) {
            //
        });
    }
}
