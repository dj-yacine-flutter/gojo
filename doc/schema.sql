-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-11-08T14:15:13.436Z

CREATE TABLE "users" (
  "id" BIGSERIAL UNIQUE NOT NULL,
  "role" varchar NOT NULL DEFAULT 'member',
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
  "id" uuid UNIQUE PRIMARY KEY NOT NULL,
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "verify_emails" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "secret_code" varchar NOT NULL,
  "is_used" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "expired_at" timestamptz NOT NULL DEFAULT (now() + interval '15 minutes')
);

CREATE TABLE "anime_movies" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "original_title" varchar NOT NULL,
  "aired" timestamptz NOT NULL,
  "release_year" integer NOT NULL,
  "rating" varchar NOT NULL DEFAULT ('PG-13 - Teens 13 or older'),
  "duration" interval NOT NULL DEFAULT ('00h 00m 00s'),
  "portriat_poster" varchar NOT NULL,
  "portriat_blur_hash" varchar NOT NULL,
  "landscape_poster" varchar NOT NULL,
  "landscape_blur_hash" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_series" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "original_title" varchar NOT NULL,
  "aired" timestamptz NOT NULL,
  "release_year" integer NOT NULL,
  "rating" varchar NOT NULL DEFAULT ('PG-13 - Teens 13 or older'),
  "portriat_poster" varchar NOT NULL,
  "portriat_blur_hash" varchar NOT NULL,
  "landscape_poster" varchar NOT NULL,
  "landscape_blur_hash" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "languages" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY UNIQUE PRIMARY KEY NOT NULL,
  "language_code" varchar UNIQUE NOT NULL,
  "language_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "studios" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY UNIQUE PRIMARY KEY NOT NULL,
  "studio_name" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "genres" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY UNIQUE PRIMARY KEY NOT NULL,
  "genre_name" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "metas" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "title" varchar NOT NULL,
  "overview" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_movie_studios" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "studio_id" integer NOT NULL
);

CREATE TABLE "anime_serie_studios" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "studio_id" integer NOT NULL
);

CREATE TABLE "anime_movie_genres" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "genre_id" integer NOT NULL
);

CREATE TABLE "anime_serie_genres" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "genre_id" integer NOT NULL
);

CREATE TABLE "anime_movie_metas" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "language_id" integer NOT NULL,
  "meta_id" bigserial NOT NULL
);

CREATE TABLE "anime_serie_metas" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "language_id" integer NOT NULL,
  "meta_id" bigserial NOT NULL
);

CREATE TABLE "anime_serie_seasons" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "season_number" integer UNIQUE NOT NULL,
  "portriat_poster" varchar NOT NULL,
  "portriat_blur_hash" varchar NOT NULL,
  "landscape_poster" varchar NOT NULL,
  "landscape_blur_hash" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_season_metas" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "season_id" bigserial NOT NULL,
  "language_id" integer NOT NULL,
  "meta_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_episodes" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "episode_number" integer UNIQUE NOT NULL,
  "season_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_episode_metas" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "episode_id" bigserial NOT NULL,
  "language_id" integer NOT NULL,
  "meta_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_season_episodes" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "season_id" bigserial NOT NULL,
  "episode_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_servers" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "episode_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_episode_servers" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "episode_id" bigserial NOT NULL,
  "server_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_server_sub_videos" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "server_id" bigserial NOT NULL,
  "video_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_server_dub_videos" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "server_id" bigserial NOT NULL,
  "video_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_videos" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "language_id" integer NOT NULL,
  "autority" varchar NOT NULL,
  "referer" varchar NOT NULL,
  "link" varchar NOT NULL,
  "quality" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_movie_servers" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_movie_server_sub_videos" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "server_id" bigserial NOT NULL,
  "video_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_movie_server_dub_videos" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "server_id" bigserial NOT NULL,
  "video_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_movie_videos" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "language_id" integer NOT NULL,
  "autority" varchar NOT NULL,
  "referer" varchar NOT NULL,
  "link" varchar NOT NULL,
  "quality" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

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

CREATE TABLE "anime_movie_torrents" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "file_name" varchar NOT NULL,
  "language_id" integer NOT NULL,
  "torrent_hash" varchar NOT NULL,
  "torrent_file" varchar NOT NULL,
  "seeds" integer NOT NULL,
  "peers" integer NOT NULL,
  "leechers" integer NOT NULL,
  "size_bytes" bigserial NOT NULL,
  "quality" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_movie_server_torrents" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "server_id" bigserial NOT NULL,
  "torrent_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_torrents" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "file_name" varchar NOT NULL,
  "language_id" integer NOT NULL,
  "torrent_hash" varchar NOT NULL,
  "torrent_file" varchar NOT NULL,
  "seeds" integer NOT NULL,
  "peers" integer NOT NULL,
  "leechers" integer NOT NULL,
  "size_bytes" bigserial NOT NULL,
  "quality" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_server_torrents" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "server_id" bigserial NOT NULL,
  "torrent_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_media" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "media_type" varchar NOT NULL,
  "media_url" varchar NOT NULL,
  "media_author" varchar NOT NULL,
  "media_platform" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_media" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "media_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_season_media" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "season_id" bigserial NOT NULL,
  "media_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_movie_media" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "media_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("full_name");

CREATE INDEX ON "sessions" ("id");

CREATE INDEX ON "anime_movies" ("id");

CREATE INDEX ON "anime_movies" ("original_title");

CREATE INDEX ON "anime_movies" ("release_year");

CREATE UNIQUE INDEX ON "anime_movies" ("original_title", "duration", "aired");

CREATE INDEX ON "anime_series" ("id");

CREATE INDEX ON "anime_series" ("original_title");

CREATE INDEX ON "anime_series" ("release_year");

CREATE UNIQUE INDEX ON "anime_series" ("original_title", "aired");

CREATE INDEX ON "languages" ("id");

CREATE INDEX ON "languages" ("language_code");

CREATE INDEX ON "studios" ("id");

CREATE INDEX ON "studios" ("studio_name");

CREATE INDEX ON "genres" ("id");

CREATE INDEX ON "genres" ("genre_name");

CREATE INDEX ON "metas" ("id");

CREATE INDEX ON "metas" ("title");

CREATE INDEX ON "anime_movie_studios" ("id");

CREATE UNIQUE INDEX ON "anime_movie_studios" ("anime_id", "studio_id");

CREATE INDEX ON "anime_serie_studios" ("id");

CREATE UNIQUE INDEX ON "anime_serie_studios" ("anime_id", "studio_id");

CREATE INDEX ON "anime_movie_genres" ("id");

CREATE UNIQUE INDEX ON "anime_movie_genres" ("anime_id", "genre_id");

CREATE INDEX ON "anime_serie_genres" ("id");

CREATE UNIQUE INDEX ON "anime_serie_genres" ("anime_id", "genre_id");

CREATE INDEX ON "anime_movie_metas" ("id");

CREATE UNIQUE INDEX ON "anime_movie_metas" ("anime_id", "language_id");

CREATE INDEX ON "anime_serie_metas" ("id");

CREATE UNIQUE INDEX ON "anime_serie_metas" ("anime_id", "language_id");

CREATE INDEX ON "anime_serie_seasons" ("id");

CREATE UNIQUE INDEX ON "anime_serie_seasons" ("anime_id", "season_number");

CREATE INDEX ON "anime_serie_season_metas" ("id");

CREATE UNIQUE INDEX ON "anime_serie_season_metas" ("season_id", "language_id");

CREATE INDEX ON "anime_serie_episodes" ("id");

CREATE UNIQUE INDEX ON "anime_serie_episodes" ("episode_number", "season_id");

CREATE INDEX ON "anime_serie_episode_metas" ("id");

CREATE UNIQUE INDEX ON "anime_serie_episode_metas" ("episode_id", "language_id");

CREATE INDEX ON "anime_serie_season_episodes" ("season_id");

CREATE INDEX ON "anime_serie_season_episodes" ("episode_id");

CREATE UNIQUE INDEX ON "anime_serie_season_episodes" ("season_id", "episode_id");

CREATE INDEX ON "anime_serie_servers" ("id");

CREATE INDEX ON "anime_serie_episode_servers" ("server_id");

CREATE INDEX ON "anime_serie_episode_servers" ("episode_id");

CREATE UNIQUE INDEX ON "anime_serie_episode_servers" ("episode_id", "server_id");

CREATE INDEX ON "anime_serie_server_sub_videos" ("server_id");

CREATE INDEX ON "anime_serie_server_sub_videos" ("video_id");

CREATE UNIQUE INDEX ON "anime_serie_server_sub_videos" ("server_id", "video_id");

CREATE INDEX ON "anime_serie_server_dub_videos" ("server_id");

CREATE INDEX ON "anime_serie_server_dub_videos" ("video_id");

CREATE UNIQUE INDEX ON "anime_serie_server_dub_videos" ("server_id", "video_id");

CREATE INDEX ON "anime_serie_videos" ("id");

CREATE INDEX ON "anime_movie_servers" ("id");

CREATE INDEX ON "anime_movie_server_sub_videos" ("server_id");

CREATE INDEX ON "anime_movie_server_sub_videos" ("video_id");

CREATE UNIQUE INDEX ON "anime_movie_server_sub_videos" ("server_id", "video_id");

CREATE INDEX ON "anime_movie_server_dub_videos" ("server_id");

CREATE INDEX ON "anime_movie_server_dub_videos" ("video_id");

CREATE UNIQUE INDEX ON "anime_movie_server_dub_videos" ("server_id", "video_id");

CREATE INDEX ON "anime_movie_videos" ("id");

CREATE INDEX ON "anime_resources" ("id");

CREATE UNIQUE INDEX ON "anime_resources" ("tmdb_id", "imdb_id", "wikipedia_url", "official_website");

CREATE INDEX ON "anime_serie_resources" ("anime_id");

CREATE INDEX ON "anime_serie_resources" ("resource_id");

CREATE INDEX ON "anime_movie_resources" ("anime_id");

CREATE INDEX ON "anime_movie_resources" ("resource_id");

CREATE INDEX ON "anime_movie_torrents" ("id");

CREATE UNIQUE INDEX ON "anime_movie_torrents" ("file_name", "language_id", "torrent_hash", "torrent_file", "size_bytes");

CREATE INDEX ON "anime_movie_server_torrents" ("server_id");

CREATE INDEX ON "anime_movie_server_torrents" ("torrent_id");

CREATE UNIQUE INDEX ON "anime_movie_server_torrents" ("server_id", "torrent_id");

CREATE INDEX ON "anime_serie_torrents" ("id");

CREATE UNIQUE INDEX ON "anime_serie_torrents" ("file_name", "language_id", "torrent_hash", "torrent_file", "size_bytes");

CREATE INDEX ON "anime_serie_server_torrents" ("server_id");

CREATE INDEX ON "anime_serie_server_torrents" ("torrent_id");

CREATE UNIQUE INDEX ON "anime_serie_server_torrents" ("server_id", "torrent_id");

CREATE INDEX ON "anime_media" ("id");

CREATE INDEX ON "anime_serie_media" ("id");

CREATE UNIQUE INDEX ON "anime_serie_media" ("anime_id", "media_id");

CREATE INDEX ON "anime_serie_season_media" ("id");

CREATE UNIQUE INDEX ON "anime_serie_season_media" ("season_id", "media_id");

CREATE INDEX ON "anime_movie_media" ("id");

CREATE UNIQUE INDEX ON "anime_movie_media" ("anime_id", "media_id");

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username") ON DELETE CASCADE;

ALTER TABLE "verify_emails" ADD FOREIGN KEY ("username") REFERENCES "users" ("username") ON DELETE CASCADE;

ALTER TABLE "anime_movie_metas" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_movies" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_studios" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_movies" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_genres" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_movies" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_servers" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_movies" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_resources" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_movies" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_media" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_movies" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_metas" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_series" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_studios" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_series" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_genres" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_series" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_seasons" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_series" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_resources" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_series" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_media" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_series" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_metas" ADD FOREIGN KEY ("language_id") REFERENCES "languages" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_metas" ADD FOREIGN KEY ("language_id") REFERENCES "languages" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_videos" ADD FOREIGN KEY ("language_id") REFERENCES "languages" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_videos" ADD FOREIGN KEY ("language_id") REFERENCES "languages" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_torrents" ADD FOREIGN KEY ("language_id") REFERENCES "languages" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_torrents" ADD FOREIGN KEY ("language_id") REFERENCES "languages" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_season_metas" ADD FOREIGN KEY ("language_id") REFERENCES "languages" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_episode_metas" ADD FOREIGN KEY ("language_id") REFERENCES "languages" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_metas" ADD FOREIGN KEY ("meta_id") REFERENCES "metas" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_metas" ADD FOREIGN KEY ("meta_id") REFERENCES "metas" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_season_metas" ADD FOREIGN KEY ("meta_id") REFERENCES "metas" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_episode_metas" ADD FOREIGN KEY ("meta_id") REFERENCES "metas" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_genres" ADD FOREIGN KEY ("genre_id") REFERENCES "genres" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_genres" ADD FOREIGN KEY ("genre_id") REFERENCES "genres" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_studios" ADD FOREIGN KEY ("studio_id") REFERENCES "studios" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_studios" ADD FOREIGN KEY ("studio_id") REFERENCES "studios" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_resources" ADD FOREIGN KEY ("resource_id") REFERENCES "anime_resources" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_resources" ADD FOREIGN KEY ("resource_id") REFERENCES "anime_resources" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_server_sub_videos" ADD FOREIGN KEY ("server_id") REFERENCES "anime_movie_servers" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_server_dub_videos" ADD FOREIGN KEY ("server_id") REFERENCES "anime_movie_servers" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_server_torrents" ADD FOREIGN KEY ("server_id") REFERENCES "anime_movie_servers" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_server_sub_videos" ADD FOREIGN KEY ("video_id") REFERENCES "anime_movie_videos" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_server_dub_videos" ADD FOREIGN KEY ("video_id") REFERENCES "anime_movie_videos" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_episodes" ADD FOREIGN KEY ("season_id") REFERENCES "anime_serie_seasons" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_season_episodes" ADD FOREIGN KEY ("season_id") REFERENCES "anime_serie_seasons" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_season_metas" ADD FOREIGN KEY ("season_id") REFERENCES "anime_serie_seasons" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_season_media" ADD FOREIGN KEY ("season_id") REFERENCES "anime_serie_seasons" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_servers" ADD FOREIGN KEY ("episode_id") REFERENCES "anime_serie_episodes" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_season_episodes" ADD FOREIGN KEY ("episode_id") REFERENCES "anime_serie_episodes" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_episode_servers" ADD FOREIGN KEY ("episode_id") REFERENCES "anime_serie_episodes" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_episode_metas" ADD FOREIGN KEY ("episode_id") REFERENCES "anime_serie_episodes" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_episode_servers" ADD FOREIGN KEY ("server_id") REFERENCES "anime_serie_servers" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_server_sub_videos" ADD FOREIGN KEY ("server_id") REFERENCES "anime_serie_servers" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_server_dub_videos" ADD FOREIGN KEY ("server_id") REFERENCES "anime_serie_servers" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_server_torrents" ADD FOREIGN KEY ("server_id") REFERENCES "anime_serie_servers" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_server_sub_videos" ADD FOREIGN KEY ("video_id") REFERENCES "anime_serie_videos" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_server_dub_videos" ADD FOREIGN KEY ("video_id") REFERENCES "anime_serie_videos" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_server_torrents" ADD FOREIGN KEY ("torrent_id") REFERENCES "anime_movie_torrents" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_server_torrents" ADD FOREIGN KEY ("torrent_id") REFERENCES "anime_serie_torrents" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_media" ADD FOREIGN KEY ("media_id") REFERENCES "anime_media" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_media" ADD FOREIGN KEY ("media_id") REFERENCES "anime_media" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_season_media" ADD FOREIGN KEY ("media_id") REFERENCES "anime_media" ("id") ON DELETE CASCADE;
