import type { NextPage } from 'next'
import { useEffect, useState } from 'react'
import styles from '../styles/Home.module.css'
import { gql } from "@apollo/client";
import client from "../apollo-client";

const Home: NextPage = () => {

  const [todos, setTodos] = useState<any[]>([])

  useEffect(()=>{
    (async()=>{
      const { data } = await client.query({
        query: gql`
        query getTodos{
          getTodos{
            id
            text
          }
        }
        `,
      });

      setTodos(data.getTodos)
    })()
  },[])

  return (
    <div className={styles.container}>
      <main className={styles.main}>
        {
          todos.length === 0 ? "nothing" : todos.map(todos=><li key={todos.id}>{todos.text}</li>)
        }
      </main>
    </div>
  )
}

export default Home
