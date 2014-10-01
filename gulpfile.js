var gulp = require('gulp');

var sass = require('gulp-sass'),
	minifyCss = require('gulp-minify-css'),
	autoprefixer = require('gulp-autoprefixer')
	uglify = require('gulp-uglify');

gulp.task('css', function() {
	return gulp.src('public/sass/*.scss')
		.pipe(sass({
			sourceComments: 'map',
			errLogToConsole: true
		}))
		.pipe(autoprefixer('last 2 version', 'ie 9', 'opera 12.1', 'ios 6', 'android 4'))
		//.pipe(minifyCss())
		.pipe(gulp.dest('public/dist'));
});

gulp.task('scripts', function() {
	return gulp.src('public/js/*.js')
		.pipe(uglify())
		.pipe(gulp.dest('public/dist'));
});

gulp.task('watch', function() {
	gulp.watch('public/sass/*.scss', ['css']);
	//gulp.watch('public/js/*.js', ['scripts']);
});

gulp.task('default', ['css', 'scripts']);