CREATE TABLE "languages" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY UNIQUE PRIMARY KEY NOT NULL,
  "language_code" varchar NOT NULL,
  "language_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "studios" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY UNIQUE PRIMARY KEY NOT NULL,
  "studio_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "genres" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY UNIQUE PRIMARY KEY NOT NULL,
  "genre_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "metas" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "title" varchar NOT NULL,
  "overview" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "languages" ("id");

CREATE UNIQUE INDEX ON "languages" ("language_code");

CREATE INDEX ON "studios" ("id");

CREATE UNIQUE INDEX ON "studios" ("studio_name");

CREATE INDEX ON "genres" ("id");

CREATE UNIQUE INDEX ON "genres" ("genre_name");

CREATE INDEX ON "metas" ("id");

CREATE INDEX ON "metas" ("title");