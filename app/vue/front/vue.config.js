// vue.config.js
const CopyWebpackPlugin = require('copy-webpack-plugin')
process.env.VUE_APP_VERSION = require('./package.json').version
const callback = (content) => {
  content_str = content.toString()
  for( key in process.env ) {
    content_str = content_str.replace('process.env.' + key, process.env[key])
  }
  return Buffer.from(content_str)
}
module.exports = {
  publicPath: process.env.VUE_APP_ASSETS_ORIGIN,
  productionSourceMap: false,
  assetsDir: process.env.NODE_ENV === 'production'
    ? `assets/${process.env.VUE_APP_VERSION}/`
    : '',
  filenameHashing: false,
  configureWebpack: {
    plugins: [
      new CopyWebpackPlugin(
        [
          {
            from: '/var/www/front/src/sw/'+process.env.NODE_ENV+'/firebase-messaging-sw.js',
            to: '/var/www/front/dist/firebase-messaging-sw.js',
            transform: callback
          }
        ]
      ),
    ]
  }
}
