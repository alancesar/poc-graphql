export default {
    posts({ id }, args, { postService }, info) {
        return postService.getPostsByAuthor(id)
    }
};