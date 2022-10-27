import { Args, Query, Resolver } from '@nestjs/graphql';
import { PostModel } from './models/post.model';
import { PostService } from './post.service';
import { HttpService, HttpModule } from '@nestjs/axios';
import { map } from 'rxjs';

@Resolver((of) => PostModel)
export class PostsResolver {
  constructor(
    private PostService: PostService,
    private readonly http: HttpService,
  ) {}

  @Query(() => [PostModel], { name: 'posts', nullable: true })
  getPosts() {
    const posts = this.http.get('http://127.0.0.1:1337/api/posts');
    // const posts = this.http.get('http://127.0.0.1:8080/api/query/getTodos');

    // frontend向けにデータを整形するロジック

    return posts.pipe(
      map((res) => {
        const formattedData = [];
        res.data.data.forEach((item) => {
          item.attributes['id'] = item.id;
          // item.attributes['Text'] = item.Text;

          formattedData.push(item.attributes);
        });
        console.log(formattedData);

        return formattedData;
      }),
    );
  }
  // getPost() {
  //   console.log(this.getPosts());
  //   return this.getPosts();
  // }
  // async getPosts() {
  //   return [
  //     {
  //       id: '1',
  //       title: 'NestJS is so good.',
  //     },
  //     {
  //       id: '2',
  //       title: 'GraphQL is so good.',
  //     },
  //   ];
  // }

  // getPosts() {
  //   console.log(this.PostService.getPosts());
  //   return this.PostService.getPosts();
  // }
  // getPosts() {
  //   const posts = this.http.get('http://127.0.0.1:1337/api/posts');
  //   //ここで取れていない？
  //   // frontend向けにデータを整形するロジック

  //   return posts.pipe(
  //     map((res) => {
  //       const formattedData = [];
  //       res.data.data.forEach((item) => {
  //         item.attributes['id'] = item.id;
  //         formattedData.push(item.attributes);
  //       });
  //       // console.log('formattedData');

  //       return formattedData;
  //     }),
  //   );
  // }
}
