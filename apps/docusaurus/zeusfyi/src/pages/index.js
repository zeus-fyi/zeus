import React from 'react';
import clsx from 'clsx';
import Link from '@docusaurus/Link';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import Layout from '@theme/Layout';

import '@docsearch/css';
import styles from './index.module.css';
import HomepageFeatures from "../components/HomepageFeatures";

function HomepageHeader() {
    const {siteConfig} = useDocusaurusContext();
    return (
        <header className={clsx('hero hero--primary', styles.heroBanner)}>
            <link rel="preconnect" href="https://B479Q2S8TS-dsn.algolia.net" crossOrigin/>
            <div className="container">
                <h1 className="hero__title">{siteConfig.title}</h1>
                <p className="hero__subtitle">{siteConfig.tagline}</p>
                <div className={styles.buttons}>
                    <Link
                        className="button button--secondary button--lg"
                        to="/docs/zk8s/intro">
                        Platform & APIs
                    </Link>
                    <Link
                        className="button button--secondary button--lg"
                        to="/docs/lb/intro">
                        RPC Load Balancer
                    </Link>
                </div>
            </div>
        </header>
    );
}

export default function Home() {
    const {siteConfig} = useDocusaurusContext();
    return (
        <Layout
            wrapperClassName={styles.backgroundHome}
            title={`${siteConfig.title} documentation`}
            description="zeusfyi documentation">
            <HomepageHeader/>
            <main>
                <HomepageFeatures/>
            </main>
        </Layout>
    );
}
