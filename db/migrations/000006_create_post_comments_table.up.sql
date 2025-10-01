-- public.post_comments definition

-- Drop table

-- DROP TABLE post_comments;

CREATE TABLE post_comments (
	comment_id int4 GENERATED ALWAYS AS IDENTITY( INCREMENT BY 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1 NO CYCLE) NOT NULL,
	user_id int4 NOT NULL,
	post_id int4 NOT NULL,
	"comment" varchar NULL,
	created_at timestamptz DEFAULT now() NULL,
	CONSTRAINT post_comments_pkey PRIMARY KEY (comment_id),
	CONSTRAINT post_comments_post_id_fkey FOREIGN KEY (post_id) REFERENCES posts(post_id),
	CONSTRAINT post_comments_user_id_fkey FOREIGN KEY (user_id) REFERENCES profiles(user_id)
);