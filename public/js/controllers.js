
var controllers = angular.module('controllers', []);

controllers.controller('SetCtrl', ['$scope', '$http', '$routeParams',
	'$location' ,'$timeout', '$interval', setController]);

function setController($scope, $http, $routeParams, $location, $timeout, $interval) {
	// Initialize Location
	$scope.location = {
		uid: ''
	};
	$scope.error = {};
	$scope.host = window.location.host;
	$scope.uidInput = '';

	if ($routeParams.uid) {
		$scope.location.uid = $routeParams.uid;
		$scope.uidSet = true;
	} else {
		$scope.uidSet = false;
	};

	$interval(function() {
		var loc = $scope.location;
		if (!loc.life) return;
		var timeNow = Math.floor((new Date().getTime()) / 1000);

		// Expire if required
		if (timeNow >= loc.expiresAt) {
			loc.expired = true;
			loc.lifePercent = '100%';
			return;
		}

		loc.lifePercent = Math.floor(100-((loc.expiresAt - timeNow) / loc.life)*100) + '%';
	}, 2000);

	$scope.getLocation = function(uid) {
		$http.get('/api/get/' + uid)
		.success(function(data) {
			$scope.location = data;
		}).error(function(data) {
			console.log(data);
		});
	}

	if ($scope.uidSet) {
		$scope.getLocation($scope.location.uid);
	}

	$scope.save = function(location) {
		$scope.statusClass = '';
		$http.post('/api/set', location)
		.success(function(data) {
			$scope.status = 'Link Updated';
			$scope.statusClass = true;
			$scope.error = {};
			$scope.location = data;
			$timeout(function() {
				$scope.statusClass = false;
			}, 1500);
		}).error(function(data) {
			console.log(data);
			if (data.url) {
				$scope.error.url = data.url[0];
			} else if (data.uid) {
				$scope.error.url = data.uid[0];
			};
		});
	}

	$scope.setUid = function(uid) {
		$location.path('/set/' + uid);
	}

}

controllers.controller('RecentCtrl', ['$scope', '$http', '$routeParams', '$interval',
	function($scope, $http, $routeParams, $interval) {

		$scope.getLocations = function() {
			$http.get('/api/recents')
			.success(function(data) {
				$scope.locations = data;
			}).error(function(data) {
				console.log(data);
			});
		}
		$scope.getLocations();


		$interval(function() {
			var timeNow = Math.floor((new Date().getTime()) / 1000);

			$scope.locations.forEach(function(loc) {
				// Expire if required
				if (timeNow >= loc.expiresAt) {
					loc.expired = true;
					loc.lifePercent = '100%';
					return;
				}

				loc.lifePercent = Math.floor(100-((loc.expiresAt - timeNow) / loc.life)*100) + '%';
			});

		}, 2500);


	}]);
