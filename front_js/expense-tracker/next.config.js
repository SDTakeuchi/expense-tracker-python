module.exports = {
  reactStrictMode: true,
  // for hot reload in dev env
  webpackDevMiddleware: config => {
    config.watchOptions = {
      poll: 800,
      aggregateTimeout: 300,
    }
    return config
  },
  // end
}
