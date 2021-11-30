module.exports = {
  root: true,
  env: {
    node: true
  },
  extends: [
    'plugin:vue/essential',
    '@vue/standard',
    '@vue/typescript/recommended'
  ],
  plugins: ['import'],
  parserOptions: {
    ecmaVersion: 2021
  },
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'no-extra-semi': 'off',
    'space-before-function-paren': ['error', {
      anonymous: 'always',
      named: 'never',
      asyncArrow: 'always'
    }],
    semi: ['error', 'always'],
    '@typescript-eslint/camelcase': ['error', { allow: ['__webpack_public_path__'] }],
    'import/order': ['error']
  },
  overrides: [
    {
      files: [
        '**/__tests__/*.{j,t}s?(x)',
        '**/tests/unit/**/*.spec.{j,t}s?(x)'
      ],
      env: {
        jest: true
      },
      rules: {
        'import/order': 'off'
      }
    }, {
      files: ['vue.config.js', 'styleguide.config.js'],
      rules: { '@typescript-eslint/no-var-requires': 'off' }
    }
  ]
};
