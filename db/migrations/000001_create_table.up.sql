CREATE TABLE "cassette" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "user_id" uuid,
  "name" varchar,
  "created_at" timestamptz,
  "updated_at" timestamptz
);

CREATE TABLE "user" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "name" varchar,
  "icon_url" varchar,
  "created_at" timestamptz,
  "updated_at" timestamptz,
  "deleted_at" timestamptz
);

CREATE TABLE "songs" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "cassette_id" uuid,
  "user_id" uuid,
  "song_number" int,
  "song_time" int,
  "name" varchar,
  "url" varchar,
  "upload_user" uuid,
  "created_at" timestamptz,
  "updated_at" timestamptz
);

ALTER TABLE "cassette" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "songs" ADD FOREIGN KEY ("cassette_id") REFERENCES "cassette" ("id");

ALTER TABLE "songs" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");
