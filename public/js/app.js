/* global define */
define(['angular'], function (angular) {
	'use strict';
	return angular.module('app', [], function($interpolateProvider){
		$interpolateProvider.startSymbol('[[');
		$interpolateProvider.endSymbol(']]');
	});
});