const path = require('path');
const mix = require('laravel-mix');

mix.copy('resources/css/app.css', 'public/css');
mix.js('resources/js/login.js', 'public/js')
    .extract(['axios', 'vuex']);
//设置manifest路径
mix.setPublicPath("public");
mix.webpackConfig({
    resolve: {
        alias: {
            "@js": path.resolve(__dirname, 'resources/js'),
            "@views": path.resolve(__dirname, 'resources/js/views'),
            "@public": path.resolve(__dirname, 'resources/js/views/public'),
            "@images": path.resolve(__dirname, 'resources/images'),
        }
    },
    //不需要webpack帮你打包的库,通过CDN打包
    externals: {
        'vue': 'Vue',
        'vue-router': 'VueRouter',
        'element-ui': 'ELEMENT'
    },
    output: {
        chunkFilename: "[name]-[chunkhash].js"
    }
});