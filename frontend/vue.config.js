const { defineConfig } = require('@vue/cli-service')
const webpack = require('webpack')
module.exports = defineConfig({
  // 确保 transpileDependencies 是数组，若只需要布尔值，可改为 [/* 依赖名 */]
  transpileDependencies: [/* 若需要转译特定依赖，可在此添加依赖名 */],
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
        // 替换为实际的后端服务地址
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
      // 正确配置 ProgressPlugin，不传额外选项
      new webpack.ProgressPlugin()
    ]
  }
})