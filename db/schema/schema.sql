CREATE TABLE "cassette" (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "user_id" uuid NOT NULL,
    "name" varchar NOT NULL,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL
);
CREATE TABLE "user" (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "name" varchar NOT NULL,
    "icon_url" varchar NOT NULL,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "deleted_at" timestamptz
);
CREATE TABLE "songs" (
    "id" uuid PRIMARY KEY,
    "cassette_id" uuid NOT NULL,
    "user_id" uuid NOT NULL,
    "song_number" int NOT NULL,
    "song_time" int,
    "name" varchar NOT NULL,
    "url" varchar NOT NULL,
    "upload_user" uuid NOT NULL,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL
);
ALTER TABLE "cassette"
ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");
ALTER TABLE "songs"
ADD FOREIGN KEY ("cassette_id") REFERENCES "cassette" ("id");
ALTER TABLE "songs"
ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");