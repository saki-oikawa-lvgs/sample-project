import { Module } from '@nestjs/common';
import { GraphQLModule } from '@nestjs/graphql';
// import { GraphQLModule } from '@nestjs/graphql';
import { ApolloDriver } from '@nestjs/apollo';
import * as path from 'path';
import { PostsModule } from './post/post.module';

@Module({
  imports: [
    GraphQLModule.forRoot({
      autoSchemaFile: path.join(process.cwd(), './schema.gql'),
      sortSchema: true,
      driver: ApolloDriver,
    }),
    PostsModule, // これ
  ],
})
export class AppModule {}
