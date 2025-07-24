const { defineConfig } = require('@vue/cli-service')
const webpack = require('webpack')
module.exports = defineConfig({
  transpileDependencies: [],
  devServer: {
    port: 3332,
    proxy: {
      '/auth': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        pathRewrite: {
          '^/auth': '/auth'
        }
      },
      '/students': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        pathRewrite: {
          '^/students': '/students'
        },
        logLevel: 'debug'
      }
    }
  },
  configureWebpack: {
    plugins: [
      new webpack.ProgressPlugin(),
      new webpack.DefinePlugin({
        __VUE_PROD_HYDRATION_MISMATCH_DETAILS__: false,
        __VUE_OPTIONS_API__: true,
        __VUE_PROD_DEVTOOLS__: false
      })
    ]
  }
})