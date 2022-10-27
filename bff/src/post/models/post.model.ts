import { Field, ObjectType } from '@nestjs/graphql';

@ObjectType()
export class PostModel {
  @Field()
  id: string;

  @Field()
  Text: string;
}
