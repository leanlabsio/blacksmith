var gulp = require("gulp"),
    del = require("del"),
    sass = require("gulp-sass"),
    ts = require("gulp-typescript");

var tsProject = ts.createProject('tsconfig.json');

var sassOptions = {
    includePaths: [
        "node_modules/foundation-sites/scss/",
        "node_modules/font-awesome/scss/",
    ]
};

var vendoredDeps = [
    "node_modules/es6-shim/es6-shim.min.js",
    "node_modules/systemjs/dist/system-polyfills.js",
    "node_modules/angular2/bundles/angular2-polyfills.js",
    "node_modules/systemjs/dist/system.src.js",
    "node_modules/rxjs/bundles/Rx.js",
    "node_modules/angular2/bundles/angular2.dev.js",
    "node_modules/angular2/bundles/router.dev.js",
    'node_modules/angular2/bundles/http.dev.js'
];

gulp.task("default", ["clean"]);

gulp.task("clean", function() {
    return del(["web"]);
});

gulp.task("copy", ["fonts"], function() {
    return gulp.src(vendoredDeps)
        .pipe(gulp.dest("web/js/"));
});

gulp.task("fonts", function() {
    return gulp.src("node_modules/font-awesome/fonts/*")
        .pipe(gulp.dest("web/fonts/"));
});

gulp.task('scripts', function() {
    var tsResult = tsProject.src() 
        .pipe(ts(tsProject));

    return tsResult.js.pipe(gulp.dest('web/js'));
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
    gulp.watch("src/**/*.ts", ["scripts"]);
    gulp.watch("src/**/*.html", ["html"]);
});
