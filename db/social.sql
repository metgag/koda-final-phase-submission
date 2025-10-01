CREATE TABLE "accounts" (
  "user_id" integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL
);

CREATE TABLE "profiles" (
  "user_id" integer PRIMARY KEY REFERENCES accounts(user_id),
  "username" varchar UNIQUE NOT NULL,
  "avatar" varchar,
  "bio" text,
  "created_at" timestamptz DEFAULT now(),
  "updated_at" timestamptz
);

CREATE TABLE "follows" (
  "user_id" integer NOT NULL REFERENCES profiles(user_id),
  "following_id" integer NOT NULL REFERENCES profiles(user_id),
  "created_at" timestamptz DEFAULT now(),
  PRIMARY KEY(user_id, following_id)
);

CREATE TABLE "posts" (
  "post_id" integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  "body" text,
  "image" varchar,
  "user_id" integer NOT NULL REFERENCES profiles(user_id),
  "created_at" timestamptz DEFAULT now()
);

CREATE TABLE "post_likes" (
  "user_id" integer NOT NULL REFERENCES profiles(user_id),
  "post_id" integer NOT NULL REFERENCES posts(post_id),
  "created_at" timestamptz DEFAULT now(),
  PRIMARY KEY(user_id, post_id)
);

CREATE TABLE "post_comments" (
  "comment_id" integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  "user_id" integer NOT NULL REFERENCES profiles(user_id),
  "post_id" integer NOT NULL REFERENCES posts(post_id),
  "comment" varchar,
  "created_at" timestamptz DEFAULT now()
);
