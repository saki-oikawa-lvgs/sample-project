import { IntrospectAndCompose } from '@apollo/gateway';
import { ApolloGatewayDriver, ApolloGatewayDriverConfig } from '@nestjs/apollo';
import { Module } from '@nestjs/common';
import { GraphQLModule } from '@nestjs/graphql';

@Module({
  imports: [
    GraphQLModule.forRoot<ApolloGatewayDriverConfig>({
      driver: ApolloGatewayDriver,
      server: {
        cors: true,
      },
      gateway: {
        supergraphSdl: new IntrospectAndCompose({
          // ここにサブグラフのGraphQLエンドポイントを追加していく
          subgraphs: [
            // { name: 'users', url: 'http://user-service/graphql' },
            { name: 'todos', url: 'http://localhost:8080/query' },
          ],
        }),
      },
    }),
  ],
})
export class AppModule {}
