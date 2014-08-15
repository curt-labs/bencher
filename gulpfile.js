var gulp = require('gulp');
var bower = require('gulp-bower');
var install = require('gulp-install');
var gutil = require('gulp-util');
var clean = require('gulp-clean');
var compass = require('gulp-compass');
var minifyCSS = require('gulp-minify-css');
var path = require('path');
var jshint = require('gulp-jshint');
var stylish = require('jshint-stylish');

gulp.task('bower',function(){
	return bower('./public/js/libs').pipe(gulp.dest('./public/js/libs'));
});
gulp.task('npm',function(){
	gulp.src(['./package.json']).pipe(install());
});

gulp.task('clean-npm',function(){
	return gulp.src('node_modules',{read:false}).pipe(clean());
});
gulp.task('clean-bower',function(){
	return gulp.src('public/js/libs',{read:false}).pipe(clean());
});

gulp.task('compass',function(){
	gulp.src('./public/sass/*.scss')
	.pipe(compass({
		css: 'public/css',
		sass:'public/sass',
		image: 'public/img'
	})).pipe(minifyCSS());
});

gulp.task('lint', function() {
  return gulp.src(['./public/js/**/*.js', '!public/js/libs/'])
    .pipe(jshint())
    .pipe(jshint.reporter(stylish))
    .pipe(jshint.reporter('fail'));
});

gulp.task('default',['compass']);
gulp.task('clean',['clean-npm','clean-bower']);