var webpack = require('webpack');

module.exports = {
    entry: {
        "app" : './src/boot.ts',
        "polyfills": './src/polyfills.ts'
    },
    output: {
        path: "web/js",
        filename: '[name].js',
    },
    resolve: {
        extensions: ['', '.js', '.ts', '.tsx']
    },
    plugins: [
        // Enable for not to produce compiled files in case of errors
        //new webpack.NoErrorsPlugin()
        new webpack.DefinePlugin({
            'ENV': JSON.stringify(process.env.NODE_ENV),
        }),
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
