
angular.module('goquizApp',[])
    .controller('userCtrl', ['$scope', '$http', function( $scope , $http ){
        $scope.data = {
            'name' :'',
            'email' :'',
        };
        $scope.take = function(){
            $http.post('/userinfo', $scope.data )
            .success(function(data) {
                console.log( data )
            })
            .error(function(data) {
                console.log('error in saving user');
            });
        }
    }]);