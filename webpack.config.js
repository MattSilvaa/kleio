const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
    entry: './client/src/index.tsx',
    mode: 'development',
    resolve: {
        extensions: ['.tsx', '.ts', '.js'],
    },
    devServer: {
        static: {
            directory: path.resolve(__dirname, 'dist'),
        },
        port: 3000,
    },
    plugins: [new HtmlWebpackPlugin({ template: './client/src/index.html' })],
    module: {
        rules: [
            {
                test: /\.ts./,
                exclude: /node_modules/,
                use: {
                    loader: 'babel-loader',
                    options: {
                        presets: ['@babel/preset-react'],
                    },
                },
            },
        ],
    },
    watchOptions: {
        ignored: /node_modules/,
        poll: true,
    },
};
