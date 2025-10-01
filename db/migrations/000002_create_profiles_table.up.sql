-- public.profiles definition

-- Drop table

-- DROP TABLE profiles;

CREATE TABLE profiles (
	user_id int4 NOT NULL,
	username varchar NOT NULL,
	avatar varchar NULL,
	bio text NULL,
	created_at timestamptz DEFAULT now() NULL,
	updated_at timestamptz NULL,
	CONSTRAINT profiles_pkey PRIMARY KEY (user_id),
	CONSTRAINT profiles_username_key UNIQUE (username),
	CONSTRAINT profiles_user_id_fkey FOREIGN KEY (user_id) REFERENCES accounts(user_id)
);