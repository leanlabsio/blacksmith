var gulp = require("gulp"),
    del  = require("del"),
    sass = require("gulp-sass");

var sassOptions = {
    includePaths: ["node_modules/foundation-sites/scss/"]
};

gulp.task("default", ["clean"]);

gulp.task("clean", function() {
    return del(["web"]);
});

gulp.task("copy", function() {
    return gulp.src("node_modules/angular/angular.js")
        .pipe(gulp.dest("web/js/"));
});

gulp.task("js", function() {
    return gulp.src("src/*.js")
        .pipe(gulp.dest("web/js"));
});

gulp.task("css", function() {
    return gulp.src("src/scss/styles.scss")
        .pipe(sass(sassOptions).on("error", sass.logError))
        .pipe(gulp.dest("web/css"));
});
