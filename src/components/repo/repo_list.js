(function(angular) {
    'use strict';

    angular.module('bs.repo').directive('repo.list', function() {
        return {
            controller: ['$http', RepoListController],
            templateUrl: '/html/components/repo/repo_list.html',
            controllerAs: 'ctrl'
        };
    });

    function RepoListController($http) {
        var _this = this,
            repos = [];

        _this.$http = $http;

        $http.get('/repo').then(function(resp) {
            _this.repos = resp.data;
        });

        _this.enableBuild = function(repo) {
            var _this = this;

            _this.$http.post('/job', {
                'repository': repo.clone_url,
                'env': []
            }).then(function(resp) {
                console.log(resp);
            });
        };
    }


}(window.angular));
