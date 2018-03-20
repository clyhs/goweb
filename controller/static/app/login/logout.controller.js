(function(){
	'use strict';

	angular
		.module('goweb.login')
		.controller('LogoutController', LogoutController);

    LogoutController.$inject = ['AuthService', '$state'];
	function LogoutController(AuthService, $state) {
            var vm = this;
            AuthService.logout();
            $state.transitionTo('login');
        }
})();

