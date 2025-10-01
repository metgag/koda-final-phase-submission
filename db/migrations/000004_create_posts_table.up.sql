-- public.posts definition

-- Drop table

-- DROP TABLE posts;

CREATE TABLE posts (
	post_id int4 GENERATED ALWAYS AS IDENTITY( INCREMENT BY 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1 NO CYCLE) NOT NULL,
	body text NULL,
	image varchar NULL,
	user_id int4 NOT NULL,
	created_at timestamptz DEFAULT now() NULL,
	CONSTRAINT posts_pkey PRIMARY KEY (post_id),
	CONSTRAINT posts_user_id_fkey FOREIGN KEY (user_id) REFERENCES profiles(user_id)
);