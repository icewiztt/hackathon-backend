CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "fullname" varchar UNIQUE NOT NULL,
  "password_encoded" varchar NOT NULL,
  "usertype" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tasks" (
  "id" serial PRIMARY KEY,
  "shortname" varchar UNIQUE NOT NULL,
  "problemname" varchar NOT NULL,
  "content" varchar NOT NULL,
  "subtasks" int NOT NULL,
  "answers" VARCHAR[6],
  "subtasks_score"  float[6],
  "official" BOOLEAN NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "submissions" (
  "id" serial PRIMARY KEY,
  "username" varchar NOT NULL,
  "fullname" varchar NOT NULL,
  "taskid" int NOT NULL,
  "taskname" varchar NOT NULL,
  "task_subtasks" int NOT NULL,
  "submission_answers" VARCHAR[6],
  "submission_results" BOOLEAN[6],
  "submission_score" float NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "submissions" ("fullname");

CREATE INDEX ON "submissions" ("fullname", "taskname");

COMMENT ON COLUMN "tasks"."subtasks" IS 'max 6 min 1';
