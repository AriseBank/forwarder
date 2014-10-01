<?php 

class Location extends BaseModel {


	protected $table = 'locations';

	protected $fillable = 
		[
			'id',
			'url'
		];

	protected static $rules = 
		[
			'id' => 'required|integer',
			'url'	=>	'required|url'
		];

}
