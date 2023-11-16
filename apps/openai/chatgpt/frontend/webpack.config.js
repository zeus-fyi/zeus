const MonacoWebpackPlugin = require('monaco-editor-webpack-plugin');
module.exports = {
    plugins: [
        new MonacoWebpackPlugin({
            languages: ['json', 'javascript', 'typescript', 'html', 'css', 'markdown', 'xml', 'yaml'],
        })
    ]
};