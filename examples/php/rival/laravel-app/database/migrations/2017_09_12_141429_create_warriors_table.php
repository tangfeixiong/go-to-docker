<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateWarriorsTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('warrior', function (Blueprint $table) {
            // $table->increments('id');
            // $table->timestamps();

            $table->bigIncrements('id');
            $table->bigInteger('war_team_id')->comment('id of war_team table');
            $table->string('name', 50)->index();
            $table->blob('avator')->nullable()->comment('image bytes');
            $table->string('title', 20)->nullable()->comment('role');
            $table->string('position', 20)->nullable()->comment('skill');
            $table->integer('rank')->nullable()->comment('summary');
            $table->text('background')->nullable()->comment('introduction, history, and honor');
            
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
        Schema::dropIfExists('warrior');
    }
}
