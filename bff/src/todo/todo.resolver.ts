// src/todo/todo.resolver.ts
import { Query, Resolver } from '@nestjs/graphql';
import { TodoModel } from './models/todo.models';
import { TodoService } from './todo.service';
import { HttpService } from '@nestjs/axios';
// import { map } from 'rxjs';
import {
  ApolloClient,
  InMemoryCache,
  createHttpLink,
  gql,
} from '@apollo/client/core';
// import { global } from 'node-fetch';

// global.fetch = require('node-fetch');

@Resolver((of) => TodoModel)
export class TodoResolver {
  // constructor(
  //   // private PostService: TodoService,
  //   // private readonly http: HttpService,
  // ) {}

  @Query(() => [TodoModel], { name: 'todos', nullable: true })
  getTodos() {
    const Query = gql`
      query getTodos {
        getTodos {
          id
          text
        }
      }
    `;

    const apolloClient = new ApolloClient({
      uri: 'http://127.0.0.1:8080/query/',
      cache: new InMemoryCache(),
    });
    async function getQuery() {
      const result = await apolloClient
        .query({ query: Query })
        .then((result) => {
          return JSON.stringify(result, undefined, 2);
        })
        .catch((error) => {
          console.log(error);
        });
      console.log(JSON.stringify(result, undefined, 2));
    }
    getQuery();
  }
}
