// File currently only being used as a reference

import Head from "next/head";
import styles from "../styles/Home.module.css";
import ClientOnly from "../components/ClientOnly";
import ApolloTest from "../components/ApolloTest";

export default function ClientSide() {
  return (
    <div className={styles.container}>
      <Head>
        <title>Apollo Test</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main className={styles.main}>
        <h1>Apollo Test</h1>
        <ClientOnly>
          <ApolloTest />
        </ClientOnly>
      </main>
    </div>
  );
}
