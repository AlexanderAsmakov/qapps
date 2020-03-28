CREATE TYPE operation_systems AS ENUM ('ios','android');

CREATE TABLE "apps" (
  "id" bigint PRIMARY KEY,
  "name" VARCHAR (255) NOT NULL DEFAULT '',
  "title" VARCHAR (255) NOT NULL DEFAULT '',
  "bundle_identifier" VARCHAR (255) NOT NULL DEFAULT '',
  "os" operation_systems NOT NULL DEFAULT 'ios',
  UNIQUE(bundle_identifier)
);

CREATE TABLE "apps_metadata" (
  "id" bigint PRIMARY KEY,
  "app_id" bigint NOT NULL,
  "name" VARCHAR (255) NOT NULL DEFAULT '',
  "keywords" text NOT NULL,
  "description" text NOT NULL,
  "release_notes" text NOT NULL,
  "marketing_url" VARCHAR (255) NOT NULL DEFAULT '',
  "privacy_url" VARCHAR (255) NOT NULL DEFAULT '',
  "support_url" VARCHAR (255) NOT NULL DEFAULT '',
  "author_id" bigint NOT NULL,
  "created_at" timestamp DEFAULT current_timestamp
);

CREATE TABLE "builds" (
  "id" bigint PRIMARY KEY,
  "app_id" bigint NOT NULL,
  "branch" VARCHAR (255) NOT NULL DEFAULT '',
  "comment" text NOT NULL,
  "bundle_version" VARCHAR (255) NOT NULL,
  "bundle_identifier" VARCHAR (255) NOT NULL,
  "bundle_name" VARCHAR (255) NOT NULL,
  "bundle_display_name" VARCHAR (255) NOT NULL,
  "file_name" VARCHAR (255) NOT NULL DEFAULT '',
  "build_key" VARCHAR (255) DEFAULT NULL,
  "is_release" boolean NOT NULL DEFAULT false,
  "created_at" timestamp DEFAULT current_timestamp
);