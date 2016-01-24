(function(angular) {
    'use strict';

    angular.module('bs.build').directive('bs.build', function() {
        return {
            controller: BuildController,
            templateUrl: '/html/components/build/build.html',
            controllerAs: 'ctrl'
        };
    });

    function BuildController() {
        var _this = this;
    }
}(window.angular));
