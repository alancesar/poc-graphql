import { ApolloServer } from 'apollo-server';
import { importSchema } from 'graphql-import';
import resolvers from './resolvers/resolvers';
import authorService from './service/author';
import postService from './service/post';

const server = new ApolloServer({
  typeDefs: importSchema('src/schema.graphql'),
  resolvers,
  context: {
    authorService,
    postService,
  },
});

server.listen().then(({ url }) => {
  console.log(`🚀  Server ready at ${url}`);
});