const prod = process.env.NODE_ENV === "production" || process.env.NODE_ENV === "test";
const webpack = require('webpack');

module.exports = {
    // publicPath: prod ? "/survey" : "/",
    productionSourceMap:  !prod,
};

// module.exports = {
//     publicPath: './',
//     outputDir: 'app/www',
//     productionSourceMap: false,
// };
