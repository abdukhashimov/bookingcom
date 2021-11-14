CREATE TABLE "user_type" (
  "id" integer PRIMARY KEY NOT NULL,
  "name" varchar
);

CREATE TABLE "users" (
  "id" varchar UNIQUE NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "phone_number" varchar NOT NULL,
  "is_verified" boolean DEFAULT 'false',
  "location" varchar,
  "user_type" integer,
  "created_at" timestamp,
  "updated_at" timestamp
);
