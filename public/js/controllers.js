
var controllers = angular.module('controllers', []);

controllers.controller('SetCtrl', ['$scope', '$http', '$routeParams', '$location',
	function($scope, $http, $routeParams, $location) {
		// Initialize Location
		$scope.location = {};
		$scope.error = {};
		$scope.host = window.location.host;

		if ($routeParams.uid) {
			$scope.location.uid = $routeParams.uid;
			$scope.uidSet = true;
		} else {
			$scope.uidSet = false;
			//$scope.location.uid = 0;
		};

		$scope.getLocation = function(uid) {
			$http.get('/api/get/' + uid)
			.success(function(data) {
				$scope.location = data;
			}).error(function(data) {
				console.log(data);
			});
		}
		$scope.getLocation($scope.location.uid);

		$scope.save = function(location) {
			$scope.statusClass = '';
			$http.post('/api/set/' + location.uid, location)
			.success(function(data) {
				$scope.status = data.message;
				$scope.statusClass = 'fadeOut';
				$scope.error = {};
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

	}]);

controllers.controller('RecentCtrl', ['$scope', '$http', '$routeParams',
	function($scope, $http, $routeParams) {

		$scope.getLocations = function() {
			$http.get('/api/recents')
			.success(function(data) {
				$scope.locations = data;
			}).error(function(data) {
				console.log(data);
			});
		}
		$scope.getLocations();


	}]);
