const VueLoaderPlugin = require("vue-loader/lib/plugin");

const path = require("path");

const webpack = require("webpack");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const CleanWebpackPlugin = require("clean-webpack-plugin");
const CopyWebpackPlugin = require("copy-webpack-plugin");
const env = process.env.NODE_ENV;
const WebpackBar = require("webpackbar");

const config = {
  mode: env,
  output: {
    path: path.resolve(__dirname, "dist"),
    filename: "[name].[hash].js"
  },
  devServer: {
    hot: true,
    host: "0.0.0.0",
    port: 4040,
    contentBase: path.resolve(__dirname, "public"),
    watchOptions: {
      poll: true
    }
  },
  resolve: {
    alias: {
      vue: "vue/dist/vue.js"
    },
    extensions: ["*", ".js", ".vue"]
  },
  module: {
    rules: [
      {
        test: /\.(js|vue)$/,
        exclude: /node_modules/,
        loader: "eslint-loader",
        enforce: "pre",
        options: {
          fix: true
        }
      },
      {
        test: /\.vue$/,
        loader: "vue-loader",
        exclude: /node_modules/
      },
      {
        test: /\.js$/,
        loader: "babel-loader",
        exclude: /node_modules/,
        include: [path.join(__dirname, "src")]
      },
      {
        test: /\.scss$/,
        use: [
          "style-loader", // creates style nodes from JS strings
          "css-loader", // translates CSS into CommonJS
          "sass-loader" // compiles Sass to CSS, using Node Sass by default
        ]
      },
      {
        test: /\.css$/,
        use: [
          "style-loader",
          process.env.NODE_ENV !== "production"
            ? "vue-style-loader"
            : MiniCssExtractPlugin.loader,
          "css-loader"
        ]
      },
      {
        test: /\.(woff(2)?|ttf|eot|svg)(\?v=\d+\.\d+\.\d+)?$/,
        use: [
          {
            loader: "file-loader"
          }
        ]
      },
      {
        test: /\.(gif|png|jpe?g|svg)$/i,
        use: [
          "file-loader",
          {
            loader: "image-webpack-loader",
            options: {
              bypassOnDebug: true, // webpack@1.x
              disable: true // webpack@2.x and newer
            }
          }
        ]
      }
    ]
  },
  plugins: [
    new WebpackBar(),
    new VueLoaderPlugin(),
    new CleanWebpackPlugin("dist", {}),
    new webpack.HotModuleReplacementPlugin({}),
    new HtmlWebpackPlugin({
      filename: path.join(__dirname, "dist", "index.html"),
      template: path.join(__dirname, "static", "index.html"),
      inject: true,
      hash: true
    }),
    new MiniCssExtractPlugin({
      filename: "style.[contenthash].css"
    }),
    new CopyWebpackPlugin([
      {
        from: path.join(__dirname, "static"),
        to: path.join(__dirname, "dist"),
        toType: "dir"
      }
    ])
  ]
};

module.exports = config;
