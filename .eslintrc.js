module.exports = {
    parser: '@typescript-eslint/parser',
    plugins: ['@typescript-eslint'],
    extends: [
        'eslint:recommended',
        'plugin:@typescript-eslint/recommended',
        'plugin:react/recommended',
    ],
    parserOptions: {
        ecmaVersion: 2021,
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
    },
    overrides: [
        {
            files: ['**/*.tsx', "**/*.ts"], // Include TypeScript JSX files
            rules: {
                // Add rules specific to .tsx files here
            },
        },
    ],
};