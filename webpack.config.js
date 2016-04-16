module.exports = {
    entry: './src/boot.ts',
    output: {
        path: "web/js",
        filename: 'bundle.js'
    },
    resolve: {
        extensions: ['', '.js', '.ts', '.tsx']
    },
    module: {
        loaders: [
            {
                test: /\.tsx?$/,
                loader: 'ts-loader'
            },
            {
                test: /\.html$/,
                loader: 'raw-loader'
            }

        ]
    }
}
