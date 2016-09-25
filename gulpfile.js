var gulp = require('gulp');

var sass = require('gulp-sass'),
	minifyCss = require('gulp-minify-css'),
	autoprefixer = require('gulp-autoprefixer')
	uglify = require('gulp-uglify'),
	concat = require('gulp-concat');	

gulp.task('css', function() {
	return gulp.src('sass/*.scss')
		.pipe(sass({
			sourceComments: 'map',
			errLogToConsole: true
		}))
		.pipe(autoprefixer('last 2 versions'))
		.pipe(minifyCss())
		.pipe(gulp.dest('public/dist'));
});

gulp.task('scripts', function() {
	return gulp.src('js/*.js')
		.pipe(concat('bundle.js'))
		.pipe(gulp.dest('public/dist'));
});

gulp.task('watch', function() {
	gulp.watch('sass/*.scss', ['css']);
	//gulp.watch('public/js/*.js', ['scripts']);
});

gulp.task('default', ['css', 'scripts']);