const webpack = require("webpack");

module.exports = {
  devServer: {
    proxy: {
      "^/api": {
        target: "http://localhost:9000",
        changeOrigin: true,
        logLevel: "debug",
        //pathRewrite: { "^/api": "/" },
      },
    },
  },
};
