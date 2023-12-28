CREATE TYPE "statuses" AS ENUM (
  'approved',
  'waiting for approval',
  'rejected'
);

CREATE TABLE "admin" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "user_name" varchar UNIQUE NOT NULL,
  "first_name" varchar NOT NULL,
  "middle_name" varchar DEFAULT null,
  "last_name" varchar NOT NULL,
  "id_number" varchar UNIQUE NOT NULL,
  "phone" varchar(10) UNIQUE NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "mentor" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "user_name" varchar UNIQUE NOT NULL,
  "first_name" varchar NOT NULL,
  "middle_name" varchar DEFAULT null,
  "last_name" varchar NOT NULL,
  "phone" varchar(10) UNIQUE NOT NULL,
  "id_number" varchar UNIQUE NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "student" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "user_name" varchar UNIQUE NOT NULL,
  "first_name" varchar NOT NULL,
  "middle_name" varchar DEFAULT null,
  "last_name" varchar NOT NULL,
  "roll_number" varchar UNIQUE NOT NULL,
  "branch" varchar NOT NULL,
  "section" varchar(1) NOT NULL,
  "course" varchar NOT NULL,
  "phone" varchar(10) UNIQUE NOT NULL,
  "mentor" bigserial UNIQUE,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "opportunity" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "link" varchar UNIQUE NOT NULL,
  "status" statuses DEFAULT 'waiting for approval',
  "created_by" bigserial NOT NULL,
  "approved_by" bigserial,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "opportunity_status_history" (
  "id" bigserial PRIMARY KEY,
  "opportunity_id" bigserial,
  "previous_status" statuses NOT NULL,
  "status" statuses NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE INDEX ON "admin" ("email");

CREATE INDEX ON "admin" ("user_name");

CREATE INDEX ON "admin" ("id_number");

CREATE INDEX ON "mentor" ("email");

CREATE INDEX ON "mentor" ("user_name");

CREATE INDEX ON "mentor" ("id_number");

CREATE INDEX ON "student" ("email");

CREATE INDEX ON "student" ("user_name");

CREATE INDEX ON "student" ("roll_number");

CREATE INDEX ON "student" ("phone");

CREATE INDEX ON "opportunity" ("status");

CREATE INDEX ON "opportunity" ("created_by");

CREATE INDEX ON "opportunity" ("approved_by");

CREATE INDEX ON "opportunity_status_history" ("status");

COMMENT ON COLUMN "admin"."middle_name" IS 'can be given or left unfilled';

COMMENT ON COLUMN "admin"."id_number" IS 'teacher id number from the id card';

COMMENT ON COLUMN "mentor"."middle_name" IS 'can be given or left unfilled';

COMMENT ON COLUMN "mentor"."id_number" IS 'teacher id number from the id card';

COMMENT ON COLUMN "student"."middle_name" IS 'can be given or left unfilled';

ALTER TABLE "student" ADD FOREIGN KEY ("mentor") REFERENCES "mentor" ("id");

ALTER TABLE "opportunity" ADD FOREIGN KEY ("approved_by") REFERENCES "admin" ("id");

ALTER TABLE "opportunity_status_history" ADD FOREIGN KEY ("opportunity_id") REFERENCES "opportunity" ("id");


