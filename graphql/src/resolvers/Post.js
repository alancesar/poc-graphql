export default {
    id({ postId }, args, context, info) {
        return postId
    },
    author({ authorId }, args, { authorService }, info) {
        return authorService.getAuthorById(authorId);
    },
};