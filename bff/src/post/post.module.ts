import { Module } from '@nestjs/common';
import { PostsResolver } from './post.resolver';
import { PostService } from './post.service';
import { HttpModule } from '@nestjs/axios';

@Module({
  imports: [HttpModule],
  providers: [PostsResolver, PostService],
})
export class PostsModule {}
