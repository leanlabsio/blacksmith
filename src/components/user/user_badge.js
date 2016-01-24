(function(angular){
    'use strict';

    angular.module('bs.user').directive('user.badge', function() {
        return {
            controller: UserBadgeController,
            templateUrl: '/html/components/user/user_badge.html',
            controllerAs: 'ctrl'
        };
    });

    function UserBadgeController() {
        var _this = this;
    }
}(window.angular));
