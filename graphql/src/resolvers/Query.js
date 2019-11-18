export default {
    async posts(parent, { onlyPublished }, { postService }, info) {
        const posts = await postService.getAllPosts(onlyPublished);

        return onlyPublished && posts.filter(post => post.publishedAt) || posts;
    },
    authors(parent, args, { authorService }, info) {
        return authorService.getAllAuthors();
    }
};