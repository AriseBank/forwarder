<?php

/*
|--------------------------------------------------------------------------
| Application Routes
|--------------------------------------------------------------------------
|
| Here is where you can register all of the routes for an application.
| It's a breeze. Simply tell Laravel the URIs it should respond to
| and give it the Closure to execute when that URI is requested.
|
*/

// API Routes
Route::get('/api/get/{uid}', 'LocationController@getLocation');
Route::get('/api/get', 'LocationController@getLocation');

Route::post('/api/set/{uid}', array('as'=>'location.update', 'uses'=>'LocationController@update'));


// Catch all for angular
Route::get('/set', function() {
	return View::make('base');
});
Route::get('/set/{path}', function() {
	return View::make('base');
})->where('path', '.*');


Route::get('{uid}', array(
	'uses' => 'LocationController@index',
	'as'   => 'location.get'
	));

