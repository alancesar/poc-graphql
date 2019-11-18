export default {
    addAuthor(parent, { data }, { authorService }, info) {
        return authorService.addAuthor(data)
    },
    updateAuthor(parent, { id, data }, { authorService }, info) {
        return authorService.updateAuthor(id, data);
    },
}