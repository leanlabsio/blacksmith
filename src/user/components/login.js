(function() {
    'use strict';

    angular.module('bs.user').directive('login', function() {
        return {
            restrict: 'EA',
            controller: ['$auth', UserController],
            templateUrl: '/html/user/components/login.html',
            controllerAs: 'lc'
        };
    });

    function UserController($auth) {
        this.authenticate = function(provider) {
            $auth.authenticate(provider)
                .then(function(res) {
                    console.log(res);
                });
        };
    }
}(window.angular));
