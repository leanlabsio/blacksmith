(function(angular) {
    'use strict';

    angular.module('bs.repo').directive('repo.list', function() {
        return {
            controller: ['$http', RepoListController],
            template: 'repo list template here'
        };
    });

    function RepoListController($http) {
        $http.get('/repo').then(function(resp) {
            console.log(resp);
        });
    }
}(window.angular));
