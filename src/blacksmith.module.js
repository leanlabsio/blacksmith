(function(angular) {
    'use strict';

    var app = angular.module('bs', ['ngComponentRouter', 'bs.user', 'bs.repo', 'satellizer']);

    app.config(['$authProvider', function($authProvider) {
        $authProvider.github({
            clientId: window.bsconfig.github.oauth.clientid
        });
    }]);

    app.component('app', {
        template: '<ng-outlet></ng-outlet>',
        controller: ['$router', AppDirectiveController]
    });

    function AppDirectiveController($router) {
        $router.config([{
            path: '/',
            component: 'home',
            name: 'Home'
        }, {
            path: '/login',
            component: 'login',
            name: 'Login'
        }, {
            path: '/repos',
            component: 'repo.list',
            name: 'Repo.list'
        }]);
    }

    app.directive('home', function() {
        return {
            controller: HomeController,
            template: 'Blackmith main page'
        };
    });

    function HomeController() {
        console.log("fsdfsdf");
    }
}(window.angular));
