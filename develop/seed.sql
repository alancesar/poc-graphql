CREATE TABLE "blog"."posts" (
	post_id uuid,
	author_id uuid,
	title text,
	slug text,
	body text,
	description text,
	created_at TIMESTAMP,
	published_at TIMESTAMP,
	categories list<text>,
	primary key (post_id)
);

CREATE TABLE "blog"."authors" (
	id uuid,
	name text,
	email text,
	username text,
	PRIMARY KEY (id)
);


INSERT INTO blog.posts (post_id, author_id, title, slug, body, description, created_at)
	VALUES(de1f28af-b42d-42e6-8564-02a4b02d3211, d5182478-c9be-4358-a91f-0f4761ac56f2, 'My very first post!', 'my-very-vrist-post', 'welcome to my first post!', 'Hello again', '2019-11-10 23:46:21+0000');

INSERT INTO blog.posts (post_id, author_id, title, slug, body, description, created_at)
	VALUES(e5765b93-a8fc-4f5c-80cf-8f602dd96ac7, d5182478-c9be-4358-a91f-0f4761ac56f2, 'My second post!', 'my-second-post', 'welcome to my last post!', 'Hello world again', '2019-11-10 23:46:23+0000');

INSERT INTO blog.posts (post_id, author_id, title, slug, body, description, created_at, published_at)
	VALUES(uuid(), c273cd32-8b37-49b5-8e63-e2d8639c132e, 'Sandra Rosa Madalena', 'sandra-rosa-madalena', 'Quero vê-la sorrir', 'Quero vê-la chorar', '2019-11-14 20:00:00+0000', '2019-11-14 21:00:00+0000');

INSERT INTO blog.authors (id, name, email, username)
	VALUES (d5182478-c9be-4358-a91f-0f4761ac56f2, 'Alan Cesar', 'alan.cesar@email.com', 'alancesar');

INSERT INTO blog.authors (id, name, email, username)
	VALUES (c273cd32-8b37-49b5-8e63-e2d8639c132e, 'Sidney Magal', 'sidney.magal@email.com', 'sidney');
