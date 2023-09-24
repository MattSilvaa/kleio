const path = require('path');

module.exports = {
    entry: 'src/index.tsx',
    mode: 'development',
    resolve: {
        extensions: ['.tsx', '.ts', '.js'],
    },
    devServer: {
        static: {
            directory: path.resolve(__dirname, 'dist'),
        },
        compress: true,
        port: 3000,

    },
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
