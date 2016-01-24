(function(angular) {
    'use strict';

    var app = angular.module('bs', ['ngComponentRouter', 'bs.home', 'bs.user', 'bs.repo', 'bs.build', 'satellizer']);

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
            component: 'bs.home',
            name: 'Home'
        }, {
            path: '/repos',
            component: 'repo.list',
            name: 'Repo.list'
        }, {
            path: '/build',
            component: 'bs.build',
            name: 'Build.index'
        }]);
    }
}(window.angular));
