-- public.follows definition

-- Drop table

-- DROP TABLE follows;

CREATE TABLE follows (
	user_id int4 NOT NULL,
	following_id int4 NOT NULL,
	created_at timestamptz DEFAULT now() NULL,
	CONSTRAINT follows_pkey PRIMARY KEY (user_id, following_id),
	CONSTRAINT follows_following_id_fkey FOREIGN KEY (following_id) REFERENCES profiles(user_id),
	CONSTRAINT follows_user_id_fkey FOREIGN KEY (user_id) REFERENCES profiles(user_id)
);