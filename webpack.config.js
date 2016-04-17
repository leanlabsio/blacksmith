var webpack = require('webpack');

module.exports = {
    entry: './src/boot.ts',
    output: {
        path: "web/js",
        filename: 'bundle.js'
    },
    resolve: {
        extensions: ['', '.js', '.ts', '.tsx']
    },
    plugins: [
        // Enable for not to produce compiled files in case of errors
        //new webpack.NoErrorsPlugin()
    ],
    devtool: "source-map",
    module: {
        loaders: [
            {
                test: /\.ts?$/,
                loader: 'ts-loader'
            },
            {
                test: /\.html$/,
                loader: 'raw-loader'
            }
        ]
    }
};
