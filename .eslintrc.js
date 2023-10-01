module.exports = {
    parser: '@typescript-eslint/parser',
    plugins: ['@typescript-eslint'],
    extends: [
        'eslint:recommended',
        'plugin:@typescript-eslint/recommended',
        'plugin:react/recommended',
    ],
    ignorePatterns: ['**/*.js'],
    parserOptions: {
        ecmaVersion: 2022,
        sourceType: 'module',
        ecmaFeatures: {
            jsx: true,
        },
    },
    settings: {
        react: {
            version: 'detect',
        },
    },
    rules: {
        'quotes': ['error', 'single'], // Enforce single quotes for strings
        'semi': ['error', 'never'], // no semicolons
    },
    overrides: [
        {
            files: ['**/*.tsx', '**/*.ts'], // Include TypeScript JSX files
            rules: {
                // Add rules specific to .tsx files here
            },
        },
    ],
};