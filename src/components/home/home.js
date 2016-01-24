(function(angular){
    angular.module('bs.home', []).directive('bs.home', function() {
        return {
            controller: HomeController,
            templateUrl: '/html/components/home/home.html',
            controllerAs: 'ctrl'
        };
    });

    function HomeController() {
        var _this = this;
    }
}(window.angular));
