<?php 

class Location extends BaseModel {


	protected $table = 'locations';

	protected $fillable = 
		[
			'uid',
			'url'
		];

	protected static $rules = 
		[
			'uid' => 'required|alpha_dash',
			'url'	=>	'required|url'
		];

}
