
var controllers = angular.module('controllers', []);

controllers.controller('SetCtrl', ['$scope', '$http', '$routeParams',
	function($scope, $http, $routeParams) {
		// Initialize Location
		$scope.location = {};
		$scope.error = {};
		$scope.host = window.location.host;

		if ($routeParams.uid) {
			$scope.location.uid = $routeParams.uid;
		} else {
			$scope.location.uid = 0;
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
				};
			});
		}

	}]);