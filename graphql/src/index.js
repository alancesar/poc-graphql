import { ApolloServer, PubSub } from 'apollo-server';
import { importSchema } from 'graphql-import';
import resolvers from './resolvers/resolvers';
import authorService from './service/author';
import postService from './service/post';

const pubsub = new PubSub()

const server = new ApolloServer({
  typeDefs: importSchema('src/schema.graphql'),
  resolvers,
  context: {
    authorService,
    postService,
    pubsub,
  },
});

server.listen().then(({ url }) => {
  console.log(`🚀  Server ready at ${url}`);
});
