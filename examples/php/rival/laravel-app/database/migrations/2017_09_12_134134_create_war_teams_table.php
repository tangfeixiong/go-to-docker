<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateWarTeamsTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('war_team', function (Blueprint $table) {
            // $table->increments('id');
            // $table->timestamps();
            
            $table->bigIncrements('id');
            $table->string('name', 50)->index();
            $table->blob('logo')->nullable()->comment('image bytes');
            $table->text('background')->nullable()->comment('introduction, history, and honor');
            
            $table->string('teamname', 20)->comment('compatial with os_team');
            $table->string('organization', 255)->nullable()->comment('compatial with os_team');
            $table->string('strengths', 20)->nullable()->comment('compatial with os_team');
            $table->string('teamlogo', 255)->nullable()->comment('compatial with os_team');
            $table->text('introduction')->nullable()->comment('compatial with os_team');
            $table->integer('status')->nullable()->comment('compatial with os_team');
            $table->dateTime('create_time')->nullable()->comment('compatial with os_team');
            $table->string('token', 32)->nullable()->comment('compatial with os_team');
            $table->string('ip_part', 100)->nullable()->comment('compatial with os_team');
            $table->string('province', 255)->nullable()->comment('compatial with os_team');
            
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
        Schema::dropIfExists('war_team');
    }
}

