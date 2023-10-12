-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-10-12T13:44:57.772Z

CREATE TABLE "anime_movie" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "original_title" varchar NOT NULL,
  "aired" timestamptz NOT NULL,
  "release_year" integer NOT NULL,
  "duration" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "studios" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY UNIQUE PRIMARY KEY NOT NULL,
  "studio_name" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_studios" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "studio_id" integer
);

CREATE TABLE "genres" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY UNIQUE PRIMARY KEY NOT NULL,
  "genre_name" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_genres" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "genre_id" integer
);

CREATE TABLE "languages" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY UNIQUE PRIMARY KEY NOT NULL,
  "language_code" varchar UNIQUE NOT NULL,
  "language_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_metas" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "language_id" integer NOT NULL,
  "meta_id" bigserial NOT NULL
);

CREATE TABLE "metas" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "title" varchar NOT NULL,
  "overview" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "users" (
  "id" BIGSERIAL UNIQUE NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "is_email_verified" bool NOT NULL DEFAULT false,
  "full_name" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("id", "username")
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "anime_movie" ("id");

CREATE INDEX ON "anime_movie" ("original_title");

CREATE INDEX ON "anime_movie" ("release_year");

CREATE UNIQUE INDEX ON "anime_movie" ("original_title", "duration", "aired");

CREATE INDEX ON "studios" ("id");

CREATE INDEX ON "studios" ("studio_name");

CREATE INDEX ON "anime_studios" ("anime_id");

CREATE INDEX ON "anime_studios" ("studio_id");

CREATE INDEX ON "genres" ("id");

CREATE INDEX ON "genres" ("genre_name");

CREATE INDEX ON "anime_genres" ("anime_id");

CREATE INDEX ON "anime_genres" ("genre_id");

CREATE INDEX ON "languages" ("id");

CREATE INDEX ON "languages" ("language_code");

CREATE INDEX ON "anime_metas" ("id");

CREATE UNIQUE INDEX ON "anime_metas" ("anime_id", "language_id");

CREATE INDEX ON "metas" ("id");

CREATE INDEX ON "metas" ("title");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("full_name");

ALTER TABLE "anime_studios" ADD FOREIGN KEY ("studio_id") REFERENCES "studios" ("id");

ALTER TABLE "anime_genres" ADD FOREIGN KEY ("genre_id") REFERENCES "genres" ("id");

ALTER TABLE "anime_metas" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_movie" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_studios" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_movie" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_genres" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_movie" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_metas" ADD FOREIGN KEY ("language_id") REFERENCES "languages" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_metas" ADD FOREIGN KEY ("meta_id") REFERENCES "metas" ("id") ON DELETE CASCADE;

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username") ON DELETE CASCADE;
