<!DOCTYPE html>
<html lang="en" ng-app="publicApp">
<head>
	<meta charset="UTF-8">
	<title>Forwarder</title>
	<link rel="stylesheet" href="/dist/styles.css">
	<link href='http://fonts.googleapis.com/css?family=RobotoDraft:300,400,500,700,400italic' rel='stylesheet' type='text/css'>
	<link href="//maxcdn.bootstrapcdn.com/font-awesome/4.2.0/css/font-awesome.min.css" rel="stylesheet">
<body>
	<header>
		<div class="container">
			<div class="logo left">Forwarder</div>
			<div class="menu right">
				<ul>
					<li><a class="button" ng-href="/"><i class="fa fa-home"></i>Home</a></li>
				</ul>
			</div>
		</div>
	</header>

	<section>
		<div class="header">
			<div class="container">
				<h3>Editing URL</h3>
				@if(Session::has('status'))
					<p class="status"><i class="fa fa-thumbs-o-up"></i>{{Session::get('status')}}</p>
				@endif
			</div>
		</div>
	</section>

	<!-- Form Specific -->

	<section class="container url-form">

		{{ Form::model($location, array('route' => array('location.update', $location->id))) }}

			{{ Form::text('url') }}

			{{ Form::submit('Save', array('class'=>'button pos')) }}

			<p>{{$errors->first('url')}}</p>

		{{ Form::close() }}

		<p>This forwarded url can be accessed at <a href="{{$accessUrl}}">{{$accessUrl}}</a></p>
	</section>

	<!-- End Form Specific -->

	<footer>
		<div></div>
	</footer>
</body>
</html>