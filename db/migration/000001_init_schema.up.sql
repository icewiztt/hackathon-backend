CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "username" varchar NOT NULL,
  "fullname" varchar NOT NULL,
  "password_encoded" varchar NOT NULL,
  "usertype" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tasks" (
  "id" serial PRIMARY KEY,
  "shortname" varchar NOT NULL,
  "problemname" varchar NOT NULL,
  "content" varchar NOT NULL,
  "subtasks" int NOT NULL,
  "answers" VARCHAR[6],
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "submissions" (
  "id" serial PRIMARY KEY,
  "from_user_id" int NOT NULL,
  "to_task_id" int NOT NULL,
  "task_subtasks" int NOT NULL,
  "submission_answers" VARCHAR[6],
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "submissions" ADD FOREIGN KEY ("from_user_id") REFERENCES "users" ("id");

ALTER TABLE "submissions" ADD FOREIGN KEY ("to_task_id") REFERENCES "tasks" ("id");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "tasks" ("problemname");

CREATE INDEX ON "submissions" ("from_user_id");

CREATE INDEX ON "submissions" ("to_task_id");

CREATE INDEX ON "submissions" ("from_user_id", "to_task_id");

COMMENT ON COLUMN "tasks"."subtasks" IS 'max 6 min 1';
