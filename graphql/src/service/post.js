import rp from 'request-promise';

const POST_API_URL = 'http://localhost:8080/api/posts';

export const getAllPosts = () => {
    return rp.get(POST_API_URL, {
        json: true
    });
};


export const getPostById = (id) => {
    return rp.get(`${POST_API_URL}/${id}`, {
        json: true
    });
};

export const getPostsByAuthor = (authorID) => {
    return rp.get(`${POST_API_URL}/search?author=${authorID}`, {
        json: true
    });
};

export default {
    getAllPosts,
    getPostById,
    getPostsByAuthor,
};