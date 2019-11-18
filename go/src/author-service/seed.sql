CREATE  KEYSPACE IF NOT EXISTS blog WITH REPLICATION = {
    'class' : 'SimpleStrategy',
    'replication_factor': 1
    };

USE blog;

CREATE TABLE "blog"."authors" (
	id uuid,
	name text,
	email text,
	username text,
	PRIMARY KEY (id)
);

INSERT INTO blog.authors (id, name, email, username)
VALUES (d5182478-c9be-4358-a91f-0f4761ac56f2, 'Alan Cesar', 'alan.cesar@email.com', 'alancesar');

INSERT INTO blog.authors (id, name, email, username)
VALUES (c273cd32-8b37-49b5-8e63-e2d8639c132e, 'Sidney Magal', 'sidney.magal@email.com', 'sidney');
