-- Create "games" table
CREATE TABLE "public"."games" (
  "game_id" text NOT NULL,
  "name" text NULL,
  "description" text NULL,
  "price" numeric NULL,
  "genre" text NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("game_id")
);
