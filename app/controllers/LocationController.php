<?php

class LocationController extends BaseController {

	protected $location;

	public function __construct(Location $location)
	{
		$this->location = $location;
	}


	public function index($uid = 0)
	{
		$this->location = $this->getLocation($uid);
		return Redirect::to($this->location->url);
	}

	public function getLocation($uid = 0)
	{
		$this->location = $this->location->where('uid', '=', $uid)->first();

		if (is_null($this->location)) {
			$location = new Location();
			$location->uid = $uid;
			$location->url = 'http://www.google.com';
			return $location;
		}

		return $this->location;
	}


	public function update($uid = 0)
	{
		$this->location = $this->getLocation($uid);
		$this->location->url = Input::get('url');

		if ($this->location->validate()->hasErrors()) {
			return Response::json($this->location->errors, 400);
		}

		$this->location->save();
		return array('message' => 'URL Saved');
	}

	public function getRecent()
	{
		return DB::table('locations')->orderBy('updated_at', 'desc')->take(10)->get();
	}


}
