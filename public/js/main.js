/* global document */

require.config({
	paths: {
		angular: '/public/js/libs/angular/angular.min',
		jquery: '/public/js/libs/jquery/dist/jquery.min',
		ngResource: '/public/js/libs/angular-resource/angular-resource.min'
	},
	shim: {
		angular: {
			exports: 'angular'
		},
		jquery:{
			exports: '$'
		},
		ngResource: ['angular']
	}
});

require(['jquery'], function () {
	'use strict';

	var tabs = document.querySelector('paper-tabs');
	tabs.addEventListener('core-select',function(){
		console.log('Selected: ' + tabs.selected);
	});

	var menu = document.querySelector('core-menu')
	menu.addEventListener('core-activate', function(){
		document.querySelector('core-drawer-panel').togglePanel();
	});

});