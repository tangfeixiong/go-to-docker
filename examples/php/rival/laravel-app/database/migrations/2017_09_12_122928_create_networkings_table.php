// [vagrant@bogon laravel-app]$ php artisan make:model -m Networking
// Model created successfully.
// Created Migration: 2017_09_12_122928_create_networkings_table

<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateNetworkingsTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('networkings', function (Blueprint $table) {
            // $table->increments('id');
            // $table->timestamps();

            $table->bigIncrements('id');
            
            $table->bigInteger('battlefield_id')->index()->comment('id of battlefield table');
            $table->string('name', 100)->index()->comment('room number');
            $table->string('bridge', 100)->nullable()->comment('docker bridge name');
            $table->string('subnets', 50)->nullable()->comment('CIDR');
            $table->string('gateway', 50)->nullable()->comment('bridge ipv4 address');
            $table->string('dns', 50)->nullable()->comment('domain name server, for example embeded dns: 127.0.0.11');
            $table->string('host_routes', 100)->nullable()->comment('next hop other than defaul');
            $table->string('infra_type')->nullable()->comment('docker bridge, kvm bridge, physical, or others');
            
            $table->timestampsTz();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('networkings');
    }
}
