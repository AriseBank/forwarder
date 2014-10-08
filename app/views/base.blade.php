<!DOCTYPE html>
<html lang="en" ng-app="forwarder">
<head>
	<meta charset="UTF-8">
	<title>Forwarder</title>
	<meta name="viewport" content="width=device-width, user-scalable=no">
	<link rel="stylesheet" href="/dist/styles.css">
	<link href='http://fonts.googleapis.com/css?family=RobotoDraft:300,400,500,700,400italic' rel='stylesheet' type='text/css'>
	<link href="//maxcdn.bootstrapcdn.com/font-awesome/4.2.0/css/font-awesome.min.css" rel="stylesheet">
	<script src="//ajax.googleapis.com/ajax/libs/angularjs/1.2.26/angular.min.js"></script>
	<script src="//ajax.googleapis.com/ajax/libs/angularjs/1.2.23/angular-route.min.js"></script>
	<script src="https://ajax.googleapis.com/ajax/libs/angularjs/1.2.23/angular-resource.min.js"></script>
	<script src="/js/app.js"></script>
	<script src="/js/controllers.js"></script>

<body>
	<header>
		<div class="container">
			<div class="logo left">Forwarder</div>
			<div class="menu right">
				<ul>
					<li><a class="button" ng-href="/recent"><i class="fa fa-clock-o"></i>Recent</a></li>
				</ul>
			</div>
		</div>
	</header>



	@yield('content')

	<section ng-view></section>

	<footer>
		<div></div>
	</footer>
</body>
</html>