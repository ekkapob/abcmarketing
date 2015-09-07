var browserSync = require('browser-sync').create(),
    gulp        = require('gulp'),
    concat      = require('gulp-concat'),
    plumber     = require('gulp-plumber'),
    minifyCss   = require('gulp-minify-css'),
    sass        = require('gulp-sass'),
    sourcemaps  = require('gulp-sourcemaps');
    // wait        = require('gulp-wait');

gulp.task('default', ['serve'], function(){
});

gulp.task('serve', ['sass'], function(){
  browserSync.init({
    proxy: 'localhost:8080'
  });

  gulp.watch('./assets/stylesheets/*.scss', ['sass']);
  gulp.watch(['./public/css/*.css', './templates/*.html'], browserSync.reload);
  gulp.watch(['./templates/*.ace'], function(){
    setTimeout(browserSync.reload, 200);
  });
});

gulp.task('sass', function(){
  return gulp.src('./assets/stylesheets/*.scss')
            .pipe(plumber())
            .pipe(sourcemaps.init())
            .pipe(sass.sync())
            .pipe(minifyCss())
            .pipe(concat('styles.css'))
            .pipe(sourcemaps.write())
            .pipe(gulp.dest('./public/css'));
});

