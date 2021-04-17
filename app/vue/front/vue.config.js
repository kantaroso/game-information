// vue.config.js
process.env.VUE_APP_VERSION = require('./package.json').version
module.exports = {
    publicPath: process.env.VUE_APP_ASSETS_ORIGIN,
    productionSourceMap: false,
    assetsDir: process.env.NODE_ENV === 'production'
    ? process.env.VUE_APP_VERSION + '/'
    : '',
    filenameHashing: false,
}
