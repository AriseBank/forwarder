
var forwarder = angular.module('forwarder', ['ngRoute', 'controllers']);

forwarder.config(['$routeProvider', '$locationProvider',
	function($routeProvider, $locationProvider) {
		
		$routeProvider.
			when('/set/:uid', {
				templateUrl: '/parts/set.html',
				controller: 'SetCtrl'
			}).when('/set', {
				templateUrl: '/parts/set.html',
				controller: 'SetCtrl'
			}).when('/recent', {
				templateUrl: '/parts/recent.html',
				controller: 'RecentCtrl'
			});

		$locationProvider.html5Mode(true);

	}]);