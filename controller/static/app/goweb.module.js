(function(){
    'use strict';

    angular
        .module('goweb', [
                'goweb.core',
                'goweb.services',
                'goweb.layout',
                'goweb.login',
				'goweb.main',
                'goweb.filters',
                'angular-jwt',
                'base64',
                'selectize',
                'ui.router'
        ]);

})();
