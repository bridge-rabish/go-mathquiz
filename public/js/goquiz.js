angular.module('goquizApp', [])
    .controller('userCtrl', ['$scope', '$http', 'quizService', function( $scope , $http , quizService  ){
        $scope.data = {
            'name' :'',
            'email' :'',
            'show' : true
        };
        
        $scope.take = function(){
            $scope.data.show = false;
            quizService.post( $scope.data )
        }
    }])
    .controller('quesCtrl', ['$scope', '$http', 'quizService', function( $scope , $http , quizService ){
        var quiz = {
            questions : []
        };
        $scope.quiz = quiz  ;
        $scope.data ={
            show : false
        };
        quizService.set( $scope );
    }])
    .factory('quizService', [ '$http', function($http){
        var $scope ;
        return {
            set : function ( scope ){
                $scope = scope ;
            },
            post : function ( data ){
                $http.post('/userinfo', data )
                .success(function(data) {
                    $scope.quiz = data;
                    $scope.data.show = true  ;
                })
                .error(function(data) {
                    console.log('error in saving user');
                });
            },
        };
    }]);