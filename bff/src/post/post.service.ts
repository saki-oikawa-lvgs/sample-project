import { Injectable } from '@nestjs/common';
import { map } from 'rxjs';
import { HttpService, HttpModule } from '@nestjs/axios';
import { Test } from '@nestjs/testing';

@Injectable()
export class PostService {
  constructor(private readonly http: HttpService) {}
  getPosts() {
    const posts = this.http.get('http://127.0.0.1:1337/api/posts');
    //ここで取れていない？
    // frontend向けにデータを整形するロジック

    return posts.pipe(
      map((res) => {
        const formattedData = [];
        res.data.data.forEach((item) => {
          item.attributes['id'] = item.id;
          formattedData.push(item.attributes);
        });
        // console.log(formattedData);
        // console.log('formattedData');

        return formattedData;
      }),
    );
  }
}
// async function test() {
//   const moduleRef = await Test.createTestingModule({
//     imports: [HttpModule], // HttpModule（Axios）のHttpServiceを使いたいのでNestJSにインスタンスを作ってもらう
//     providers: [PostService], // 自分で定義したこのクラスのインスタンスを使いたいから作ってもらう
//   }).compile();

//   const repo = moduleRef.get<PostService>(PostService); // 作ってもらったインスタンスを取り出す
//   console.log(await repo.getPosts()); // どんなデータが取れてるかな？
// }
// test();
