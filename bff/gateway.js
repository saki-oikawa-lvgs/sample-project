/* eslint-disable*/
const { ApolloServer } = require('apollo-server'); 
const { ApolloGateway } = require('@apollo/gateway');

// Initialize an ApolloGateway instance and pass it an array of
// your subgraph names and URLs
const gateway = new ApolloGateway({
  serviceList: [
    { name: 'todo-go', url: 'http://localhost:8083/query' },
    { name: 'post-go', url: 'http://localhost:8082/query' },
  ],
});

// Pass the ApolloGateway to the ApolloServer constructor
const server = new ApolloServer({
  gateway,

  // Disable subscriptions (not currently supported with ApolloGateway)
  subscriptions: false,
});

server.listen(4001).then(({ url }) => {
  console.log(`🚀 Server ready at ${url}`);
});