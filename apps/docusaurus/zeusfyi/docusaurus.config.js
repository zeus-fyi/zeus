// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');

/** @type {import('@docusaurus/types').Config} */
const config = {
    title: 'Zeusfyi',
    tagline: 'Show Me How To Use...',
    favicon: 'img/icon.svg',

    // Set the production url of your site here
    url: 'https://docs.zeus.fyi',
    baseUrl: '/',
    onBrokenLinks: 'throw',
    onBrokenMarkdownLinks: 'throw',
    plugins: [
        [
            '@docusaurus/plugin-google-gtag',
            {
                trackingID: 'G-D8FGVE4D6N',
            },
        ],
    ],
    // Even if you don't use internalization, you can use this field to set useful
    // metadata like html lang. For example, if your site is Chinese, you may want
    // to replace "en" with "zh-Hans".
    i18n: {
        defaultLocale: 'en',
        locales: ['en'],
    },
    markdown: {
        mermaid: true,
    },
    themes: ['@docusaurus/theme-mermaid'],
    presets: [
        [
            '@docusaurus/preset-classic',
            ({
                sitemap: {
                    changefreq: 'hourly',
                    priority: 0.5,
                    filename: 'sitemap.xml',
                },
                docs: {
                    sidebarPath: require.resolve('./sidebars.js'),
                },
                theme: {
                    customCss: require.resolve('./src/css/custom.css'),
                },

            }),
        ],
        [
            'redocusaurus',
            {
                // Plugin Options for loading OpenAPI files
                specs: [
                    {
                        spec: '/Users/alex/go/Olympus/Zeus/openapi_v3_1/mockingbird.yaml',
                        route: '/docs/mockingbird/api/',
                    },
                ],
                // Theme Options for modifying how redoc renders them
                theme: {
                    // Change with your site colors
                    primaryColor: '#1890ff',
                },
            },
        ],
    ],

    // algolia if not using preset-classic
    // themes: ['@docusaurus/theme-search-algolia'],
    themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
        ({
            // Replace with your project's social card
            image: 'img/icon.png',
            navbar: {
                title: 'Zeusfyi',
                logo: {
                    alt: 'zeusfyi logo',
                    src: 'img/icon.svg',
                },
                items: [
                    {
                        label: 'Login',
                        position: 'right',
                        href: 'https://cloud.zeus.fyi/login',
                    },
                    {to: '/docs/mockingbird/intro', label: 'Mockingbird', position: 'left'},
                    {to: '/docs/zK8s/intro', label: 'Platform & APIs', position: 'left'},
                    {to: '/docs/lb/intro', label: 'Adaptive RPC Load Balancer', position: 'left'},
                    {
                        label: 'LinkTree',
                        position: 'right',
                        href: 'https://linktr.ee/zeusfyi',
                    },
                ],
            },
            announcementBar: {
                id: 'support_us',
                content:
                    '<strong>⭐️ If you like Zeusfyi, give it a <a href="https://github.com/zeus-fyi/zeus" target="_blank" rel="noopener noreferrer">star on GitHub</a> and follow us on <a href="https://twitter.com/zeus_fyi" target="_blank" rel="noopener noreferrer">Twitter</a></strong>',
                backgroundColor: '#B2E5E4', // Light turquoise/teal background
                textColor: '#1C6865',      // Darker teal text
                isCloseable: true,
            },
            algolia: {
                appId: 'B479Q2S8TS',
                apiKey: 'a4cb394096549d2cef4c16cb32a6435b',
                indexName: 'zeusfyi_docs',
            },
            footer: {
                style: 'dark',
                links: [
                    {
                        title: 'Support',
                        items: [
                            {
                                label: 'Discord',
                                href: 'https://discord.gg/g3jtumw7B7',
                            },
                            {
                                label: 'Status',
                                href: 'https://status.zeus.fyi/'
                            },
                            {
                                label: 'Solutions Engineering',
                                href: 'https://calendly.com/zeusfyi/solutions-engineering'
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
                                href: 'https://medium.zeus.fyi/',
                            },
                            {
                                label: 'GitHub',
                                href: 'https://github.com/zeus-fyi/zeus',
                            },
                        ],
                    },
                ],
                copyright: `Zeusfyi, Inc © ${new Date().getFullYear()}`,
            },
            prism: {
                theme: lightCodeTheme,
                darkTheme: darkCodeTheme,
            },
        }),
};

module.exports = config;
