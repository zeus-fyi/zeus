// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');

/** @type {import('@docusaurus/types').Config} */
const config = {
    title: 'zeusfyi',
    tagline: 'Adaptive RPC Load Balancing',
    favicon: 'img/icon.svg',

    // Set the production url of your site here
    url: 'https://cloud.zeus.fyi',
    // Set the /<baseUrl>/ pathname under which your site is served
    // For GitHub pages deployment, it is often '/<projectName>/'
    baseUrl: '/',

    // GitHub pages deployment config.
    // If you aren't using GitHub pages, you don't need these.
    organizationName: 'zeusfyi', // Usually your GitHub org/user name.
    projectName: 'zeus', // Usually your repo name.

    onBrokenLinks: 'warn',
    onBrokenMarkdownLinks: 'warn',

    // Even if you don't use internalization, you can use this field to set useful
    // metadata like html lang. For example, if your site is Chinese, you may want
    // to replace "en" with "zh-Hans".
    i18n: {
        defaultLocale: 'en',
        locales: ['en'],
    },

    presets: [
        [
            'classic',
            /** @type {import('@docusaurus/preset-classic').Options} */
            ({
                docs: {
                    sidebarPath: require.resolve('./sidebars.js'),
                    // Please change this to your repo.
                    // Remove this to remove the "edit this page" links.
                    editUrl:
                        'https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/',
                },
                blog: {
                    showReadingTime: true,
                    // Please change this to your repo.
                    // Remove this to remove the "edit this page" links.
                    editUrl:
                        'https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/',
                },
                theme: {
                    customCss: require.resolve('./src/css/custom.css'),
                },
            }),
        ],
    ],

    themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
        ({
            // Replace with your project's social card
            image: 'img/icon.png',
            navbar: {
                title: 'zeusfyi',
                logo: {
                    alt: 'zeusfyi logo',
                    src: 'img/icon.svg',
                },
                items: [
                    {
                        label: 'Sign Up',
                        position: 'right',
                        href: 'https://marketplace.quicknode.com/add-on/zeusfyi-4',
                    },
                    {
                        label: 'Medium',
                        href: 'https://medium.com/zeusfyi',
                    },
                    {
                        label: 'GitHub',
                        position: 'right',
                        href: 'https://github.com/zeus-fyi/zeus',
                    },
                    {
                        label: 'Discord',
                        href: 'https://discord.gg/g3jtumw7B7',
                        position: 'right',
                    },
                ],
            },
            footer: {
                style: 'dark',
                links: [
                    {
                        title: 'Documentation',
                        items: [
                            {
                                label: 'Introduction',
                                to: '/docs/intro',
                            },
                        ],
                    },
                    {
                        title: 'Support',
                        items: [
                            {
                                label: 'Discord',
                                href: 'https://discord.gg/g3jtumw7B7',
                            },
                        ],
                    },
                    {
                        title: 'Social Media',
                        items: [
                            {
                                label: 'LinkedIn',
                                href: 'https://www.linkedin.com/company/zeusfyi',
                            },
                            {
                                label: 'Twitter',
                                href: 'https://twitter.com/zeus_fyi',
                            },
                        ],
                    },
                    {
                        title: 'Resources',
                        items: [
                            {
                                label: 'Medium',
                                href: 'https://medium.com/zeusfyi',
                            },
                            {
                                label: 'GitHub',
                                href: 'https://github.com/zeus-fyi/zeus',
                            },
                        ],
                    },
                ],
                copyright: `zeusfyi © ${new Date().getFullYear()}`,
            },
            prism: {
                theme: lightCodeTheme,
                darkTheme: darkCodeTheme,
            },
        }),
};

module.exports = config;
