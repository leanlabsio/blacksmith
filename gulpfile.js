var gulp = require("gulp");

var sassOptions = {
    includePaths: [
        "node_modules/foundation-sites/scss/",
        "node_modules/font-awesome/scss/",
    ]
};

var vendoredDeps = [
    "node_modules/angular2/bundles/angular2-polyfills.js",
];

gulp.task("vendor", function() {
    return gulp.src(vendoredDeps)
        .pipe(gulp.dest("web/js/"));
});
