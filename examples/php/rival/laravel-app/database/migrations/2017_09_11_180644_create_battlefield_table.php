<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateBattlefieldTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('battlefield', function (Blueprint $table) {
            // $table->increments('id');
            $table->bigIncrements('id');
            
            $table->string('name', 100);
            $table->text('description');
            $table->string('match_item', 100)->comment('比赛项目: web, 数据库, 等等');
            $table->string('status', 20)->comment('房间或赛台状态: 就绪, 未就绪, 已开始, 已结束, 等等');
            $table->string('payment', 100)->nullable()->comment('使用方式: 开放, 不对外, 公测, 收费, 等等');
            
            // $table->timestamps();
            // $table->date('created_at');
            // $table->dateTime('created_at');
            $table->dateTimeTz('created_at')->nullable();
            $table->dateTimeTz('last_updated')->nullable();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('battlefield');
    }
}
