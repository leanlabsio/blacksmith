var gulp = require("gulp"),
    del = require("del"),
    sass = require("gulp-sass");

var sassOptions = {
    includePaths: [
        "node_modules/foundation-sites/scss/",
        "node_modules/font-awesome/scss/",
    ]
};

var vendoredDeps = [
    "node_modules/angular2/bundles/angular2-polyfills.js",
];

gulp.task("default", ["clean"]);

gulp.task("clean", function() {
    return del(["web"]);
});

gulp.task("vendor", function() {
    return gulp.src(vendoredDeps)
        .pipe(gulp.dest("web/js/"));
});

gulp.task("fonts", function() {
    return gulp.src("node_modules/font-awesome/fonts/*")
        .pipe(gulp.dest("web/fonts/"));
});

gulp.task("html", function() {
    return gulp.src("src/**/*.html")
        .pipe(gulp.dest("web/html"));
});

gulp.task("css", function() {
    return gulp.src("src/scss/styles.scss")
        .pipe(sass(sassOptions).on("error", sass.logError))
        .pipe(gulp.dest("web/css"));
});

gulp.task("watch", function() {
    gulp.watch("src/**/*.html", ["html"]);
});
