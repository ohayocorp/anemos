import path from 'path';
import { themes as prismThemes } from 'prism-react-renderer';
import type { Config } from '@docusaurus/types';
import type * as Preset from '@docusaurus/preset-classic';

const config: Config = {
    title: 'Anemos',
    tagline: 'More than a package manager for Kubernetes',
    favicon: 'img/favicon.ico',

    url: 'https://anemos.ohayocorp.com',
    baseUrl: '/anemos/docs/',

    onBrokenLinks: 'throw',
    onBrokenMarkdownLinks: 'warn',

    i18n: {
        defaultLocale: 'en',
        locales: ['en'],
    },

    markdown: {
        preprocessor: ({ filePath, fileContent }) => {
            return fileContent.replaceAll('{{CODE_PATH}}', path.resolve(__dirname, path.join('..', 'docs-code')).replaceAll(/\\/g, '/'));
        },
        mermaid: true,
    },

    presets: [
        [
            'classic',
            {
                docs: {
                    sidebarPath: './sidebars.ts',
                    routeBasePath: '/',
                },
                blog: false,
                theme: {
                    customCss: './src/css/custom.css',
                },
            } satisfies Preset.Options,
        ],
    ],

    themes: [
        '@docusaurus/theme-mermaid',
    ],

    themeConfig: {
        image: 'img/anemos-social-card.jpg',
        navbar: {
            title: 'Anemos',
            logo: {
                alt: 'Anemos Logo',
                src: 'img/logo.png',
            },
            items: [
                {
                    type: 'docSidebar',
                    sidebarId: 'docsSidebar',
                    position: 'left',
                    label: 'Docs',
                },
                {
                    href: 'https://github.com/ohayocorp/anemos',
                    label: 'GitHub',
                    position: 'right',
                },
            ],
        },
        footer: {
            style: 'dark',
            links: [
                {
                    title: 'Community',
                    items: [
                        {
                            label: 'Stack Overflow',
                            href: 'https://stackoverflow.com/questions/tagged/anemos',
                        },
                        {
                            label: 'GitHub',
                            href: 'https://github.com/ohayocorp/anemos',
                        },
                    ],
                },
                {
                    title: 'Social',
                    items: [
                        {
                            label: 'X',
                            href: 'https://x.com/ohayocorp',
                        },
                    ],
                },
            ],
            copyright: `Copyright © ${new Date().getFullYear()} Ohayocorp, Inc. Built with Docusaurus.`,
        },
        prism: {
            theme: prismThemes.oneLight,
            darkTheme: prismThemes.oneDark,
            additionalLanguages: ["diff", "bash", "powershell"],
        },
        algolia: {
            appId: '72G6IVIRR8',
            apiKey: '09b05eef254dc39321153599ad7e15d8',
            indexName: 'Anemos Docs',
            contextualSearch: true,
            insights: true,
        },
    } satisfies Preset.ThemeConfig,
};

export default config;
