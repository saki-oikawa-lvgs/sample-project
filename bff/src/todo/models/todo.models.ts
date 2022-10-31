import { Field, ObjectType } from '@nestjs/graphql';

@ObjectType()
export class TodoModel {
  @Field()
  id: string;

  @Field()
  Text: string;
}
