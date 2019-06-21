module.exports = {
    mode: "development",
    // mode: "production",

    entry: {
        main: "./src/main.ts",
    },

    output: {
        publicPath: "/",
    },

    optimization: {
        namedChunks: true,
        splitChunks: {
            automaticNameDelimiter: "-",
            cacheGroups: {
                vendors: false,
            },
        },
    },

    module: {
        rules: [
            {
                test: /\.ts$/,
                use: "ts-loader",
            },
            {
                test: /\.scss$/,
                use: [
                    "style-loader",
                    {
                        loader: "css-loader",
                        options: {
                            sourceMap: false,
                            importLoaders: 2,
                        },
                    },
                    {
                        loader: "sass-loader",
                        options: {
                            sourceMap: false,
                            includePaths: ["./node_modules"],
                        },
                    },
                ],
            },
        ],
    },

    resolve: {
        extensions: [".ts", ".js"],
    },
};
