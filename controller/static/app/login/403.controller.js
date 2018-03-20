(function(){
	'use strict';

	angular
	    .module('goweb.login')
	    .controller('AccessDeniedController', AccessDeniedController);

	AccessDeniedController.$inject = ['$stateParams'];
	function AccessDeniedController($stateParams) {
            var vm = this;
	}
})();
