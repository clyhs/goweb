(function() {
    'use strict';

    angular
        .module('goweb')
        .run(
                ['$rootScope', '$state', '$stateParams', '$window',
                function ($rootScope, $state, $stateParams, $window) {
                    $rootScope.$state = $state;
                    $rootScope.$stateParams = $stateParams;
					
					
                    $rootScope.$on("$stateChangeStart", function(event, toState, toParams, fromState, fromParams) {
                        
                    });

                    $rootScope.$on("$stateChangeError", function(event, toState, toParams, fromState, fromParams) {
                        event.preventDefault(); 
                        console.log("$stateChangeError", event, toState, toParams, fromState, fromParams);
                        $state.transitionTo('error');
                    });

                    $rootScope.$on("$stateChangeSuccess", function(event, toState, toParams, fromState, fromParams) {
                    });
                }
    ]
        )
})();
