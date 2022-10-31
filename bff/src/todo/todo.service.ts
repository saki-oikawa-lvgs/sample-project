import { Injectable } from '@nestjs/common';
import { map } from 'rxjs';
import { HttpService } from '@nestjs/axios';
import {
  ApolloClient,
  InMemoryCache,
  ApolloProvider,
  createHttpLink,
  useQuery,
  gql,
} from '@apollo/client';
//使わない
@Injectable()
export class TodoService {
  constructor(private readonly http: HttpService) {}
  getTodos() {
    // const todos = this.http.get('http://127.0.0.1:8080/query/getTodos');
    const link = createHttpLink({
      uri: 'http://127.0.0.1:8080/query/',
      credentials: 'include',
    });

    const client = new ApolloClient({
      cache: new InMemoryCache(),
      link: link,
    });

    console.log(client);
    return client;
    // return todos.pipe(
    //   map((res) => {
    //     const formattedData = [];
    //     res.data.data.forEach((item) => {
    //       item.attributes['id'] = item.id;
    //       formattedData.push(item.attributes);
    //     });
    //     console.log('formattedData');
    //     console.log(formattedData);
    //     return formattedData;
    //   }),
    // );
  }
}
