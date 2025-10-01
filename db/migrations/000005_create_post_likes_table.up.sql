-- public.post_likes definition

-- Drop table

-- DROP TABLE post_likes;

CREATE TABLE post_likes (
	user_id int4 NOT NULL,
	post_id int4 NOT NULL,
	created_at timestamptz DEFAULT now() NULL,
	CONSTRAINT post_likes_pkey PRIMARY KEY (user_id, post_id),
	CONSTRAINT post_likes_post_id_fkey FOREIGN KEY (post_id) REFERENCES posts(post_id),
	CONSTRAINT post_likes_user_id_fkey FOREIGN KEY (user_id) REFERENCES profiles(user_id)
);