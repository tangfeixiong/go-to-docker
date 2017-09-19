<?php

namespace App\Http\Controllers;

use App\Warrior;
use Illuminate\Http\Request;

class WarriorRepoController extends Controller
{
    // [vagrant@bogon laravel-app]$ php artisan make:controller WarriorRepoController --resource --model=Warrior
    // Controller created successfully.


    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        //
    }

    /**
     * Show the form for creating a new resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function create()
    {
        //
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function store(Request $request)
    {
        //
    }

    /**
     * Display the specified resource.
     *
     * @param  \App\Warrior  $warrior
     * @return \Illuminate\Http\Response
     */
    public function show(Warrior $warrior)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  \App\Warrior  $warrior
     * @return \Illuminate\Http\Response
     */
    public function edit(Warrior $warrior)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \App\Warrior  $warrior
     * @return \Illuminate\Http\Response
     */
    public function update(Request $request, Warrior $warrior)
    {
        //
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  \App\Warrior  $warrior
     * @return \Illuminate\Http\Response
     */
    public function destroy(Warrior $warrior)
    {
        //
    }
}
