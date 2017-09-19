<?php

use Illuminate\Http\Request;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::middleware('auth:api')->get('/user', function (Request $request) {
    return $request->user();
});

// Route::post('/v1/echo', function (Request $request) {
//     return $request->all();
// });

Route::post('/v1/battlefields', 'ReSTful\BattlefieldRepoController@store');
Route::get('/v1/battlefields', 'ReSTful\BattlefieldRepoController@index');
Route::get('/v1/battlefields/{article}', 'ReSTful\BattlefieldRepoController@show');
Route::put('/v1/battlefields/{article}', 'ReSTful\BattlefieldRepoController@update');
Route::delete('/v1/battlefields/{article}', 'ReSTful\BattlefieldRepoController@delete');

