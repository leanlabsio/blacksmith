(function() {
    'use strict';

    angular.module('bs.user').directive('login', function() {
        return {
            restrict: 'EA',
            controller: ['$auth', '$router', UserController],
            templateUrl: '/html/components/user/login.html',
            controllerAs: 'lc'
        };
    });

    function UserController($auth, $router) {
        var _this = this;
        _this.$auth = $auth;
        _this.$router = $router;

        _this.authenticate = function(provider) {
            _this.$auth.authenticate(provider)
                .then(function(res) {
                    _this.$router.navigate(['Build.index']);
                });
        };
    }
}(window.angular));
