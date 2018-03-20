(function(){
    'use strict';

    angular
        .module('goweb', [
                'goweb.core',
				'goweb.services',
                'goweb.layout',
                'goweb.filters',
                'angular-jwt',
                'base64',
                'selectize',
                'ui.router'
        ]);

})();
