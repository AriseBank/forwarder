<?php

class LocationController extends BaseController {

	protected $location;

	public function __construct(Location $location)
	{
		$this->location = $location;
	}


	public function index($id = 0)
	{
		$this->location = $this->getLocation($id);
		return Redirect::to($this->location->url);
	}

	public function getLocation($id)
	{
		$this->location = $this->location->find($id);

		if (is_null($this->location)) {
			$location = new Location();
			$location->id = $id;
			$location->url = 'http://www.google.com';
			return $location;
		}

		return $this->location;
	}

	public function setForm($id = 0)
	{
		$this->location = $this->getLocation($id);
		$url = route('location.get', $id);
		return View::make('set', ['location' => $this->location, 'accessUrl' => $url]);
	}


	public function update($id = 0)
	{
		$this->location = $this->getLocation($id);
		$this->location->url = Input::get('url');

		if ($this->location->validate()->hasErrors()) {
			return Redirect::back()->withInput()->withErrors($this->location->errors);
		}

		$this->location->save();
		return Redirect::route('location.set', $id)->withStatus('URL Saved');
	}


}
