-- Add new schema named "public"
CREATE SCHEMA IF NOT EXISTS "public";
-- Set comment to schema: "public"
COMMENT ON SCHEMA "public" IS 'standard public schema';
-- Create "games" table
CREATE TABLE "public"."games" (
  "game_id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "name" text NULL,
  "description" text NULL,
  "price" numeric NULL,
  "genre" text NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("game_id")
);
