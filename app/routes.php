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

Route::get('/', array(
	'uses' => 'LocationController@index',
	'as'   => 'location.index'
	));

Route::get('/set', 'LocationController@setForm');
Route::get('/set/{id}', array(
	'uses' => 'LocationController@setForm',
	'as' => 'location.set'
	));

Route::post('/set/{id}', array('as'=>'location.update', 'uses'=>'LocationController@update'));


Route::get('{id}', array(
	'uses' => 'LocationController@index',
	'as'   => 'location.get'
	));
