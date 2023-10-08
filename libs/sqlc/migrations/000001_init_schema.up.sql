CREATE TABLE "channels" (
        "id" serial PRIMARY KEY,
        "twitch_name" varchar NOT NULL,
        "twitch_id" int UNIQUE NOT NULL
);

CREATE TABLE "commands" (
        "id" serial PRIMARY KEY,
        "channel_id" int NOT NULL,
        "command" varchar NOT NULL,
        "answer" varchar NOT NULL
);

COMMENT ON COLUMN "commands"."command" IS 'Move to another table, replace with tag_id';

COMMENT ON COLUMN "commands"."answer" IS 'Move to another table, replace with answer_id';

ALTER TABLE "commands" ADD FOREIGN KEY ("channel_id") REFERENCES "channels" ("id");
