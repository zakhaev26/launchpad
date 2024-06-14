module.exports = function(api) {
    // Configure caching
    api.cache(true);
  
    return {
      presets: [
        ['@babel/preset-env', { targets: "defaults" }],
        '@babel/preset-typescript'
      ],
      plugins: [
        // Add any Babel plugins you need here
      ]
    };
  };
  