CREATE TABLE "anime_resources" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "tmdb_id" integer NOT NULL,
  "imdb_id" varchar NOT NULL,
  "official_website" varchar NOT NULL,
  "wikipedia_url" varchar NOT NULL,
  "crunchyroll_url" varchar NOT NULL,
  "social_media" varchar[] NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_resources" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial UNIQUE NOT NULL,
  "resource_id" bigserial UNIQUE NOT NULL
);

CREATE TABLE "anime_movie_resources" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial UNIQUE NOT NULL,
  "resource_id" bigserial UNIQUE NOT NULL
);


CREATE INDEX ON "anime_resources" ("id");

CREATE UNIQUE INDEX ON "anime_resources" ("tmdb_id", "imdb_id", "wikipedia_url", "official_website");

CREATE INDEX ON "anime_serie_resources" ("anime_id");

CREATE INDEX ON "anime_serie_resources" ("resource_id");

CREATE INDEX ON "anime_movie_resources" ("anime_id");

CREATE INDEX ON "anime_movie_resources" ("resource_id");


ALTER TABLE "anime_movie_resources" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_movies" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_resources" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_series" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_resources" ADD FOREIGN KEY ("resource_id") REFERENCES "anime_resources" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_resources" ADD FOREIGN KEY ("resource_id") REFERENCES "anime_resources" ("id") ON DELETE CASCADE;
