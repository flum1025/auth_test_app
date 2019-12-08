module.exports = {
  root: true,
  env: {
    node: true,
  },
  extends: [
    'plugin:vue/strongly-recommended',
    'eslint:recommended',
    'plugin:vue/essential',
    '@vue/airbnb',
  ],
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'no-param-reassign': [
      'error',
      {
        ignorePropertyModificationsFor: [
          'state',
        ],
      },
    ],
  },
  parserOptions: {
    parser: 'babel-eslint',
  },
};
