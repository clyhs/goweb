(function(){
    'use strict';

    angular
        .module('goweb', [
		        'goweb.core',
		        'goweb.layout',
				'goweb.login',
                'goweb.filters',
                'angular-jwt',
                'base64',
                'selectize',
                'ui.router'
				
        ]);

})();
