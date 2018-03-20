(function(){
    'use strict';

    angular
        .module('goweb.main')
        .controller('MainController', MainController);

    MainController.$inject = ['$scope', 'MainService', '$state'];
    function MainController($scope, MainService, $state) {
        var vm = this;
        
        vm.error = "";
        vm.errors = [];
        

    }
})();
