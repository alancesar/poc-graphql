import rp from 'request-promise';

const AUTHOR_API_URL = 'http://localhost:8081/api/authors';

export const getAllAuthors = () => {
    return rp.get(AUTHOR_API_URL, {
        json: true,
    });
};

export const getAuthorById = (id) => {
    return rp.get(`${AUTHOR_API_URL}/${id}`, {
        json: true
    });
};

export const addAuthor = (author) => {
    return rp.post(AUTHOR_API_URL, {
        json: true,
        body: author,
    });
};

export const updateAuthor = (id, author) => {
    return rp.put(`${AUTHOR_API_URL}/${id}`, {
        json: true,
        body: author,
    });
}

export default {
    getAllAuthors,
    getAuthorById,
    addAuthor,
    updateAuthor,
};