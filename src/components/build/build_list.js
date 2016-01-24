(function(angular) {
    'use strict';

    angular.module('bs.build').directive('build.list', function() {
        return {
            templateUrl: '/html/components/build/build_list.html',
            controller: ['$http', BuildListController],
            controllerAs: 'ctrl'
        };
    });

    function BuildListController($http) {
        var _this = this,
            builds = [];

        _this.$http = $http;

        $http.get('/builds').then(function(resp) {
            _this.builds = resp.data;
        });
    }
}(window.angular));
