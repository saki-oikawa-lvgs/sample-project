import type { NextPage } from 'next'
import { useEffect, useState } from 'react'
import styles from '../styles/Home.module.css'
import { gql } from "@apollo/client";
import client from "../apollo-client";

const Home: NextPage = () => {

  const [posts, setPosts] = useState<any[]>([])

  useEffect(()=>{
    (async()=>{
      const { data } = await client.query({
        query: gql`
        query getPosts{
          getPosts{
            id
            text
          }
        }
        `,
      });

      setPosts(data.getPosts)
    })()
  },[])

  return (
    <div className={styles.container}>
      <main className={styles.main}>
        {
          posts.length === 0 ? "nothing" : posts.map(posts=><li key={posts.id}>{posts.text}</li>)
        }
      </main>
    </div>
  )
}

export default Home
