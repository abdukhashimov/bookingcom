CREATE TABLE "user_type" (
  "id" integer PRIMARY KEY NOT NULL,
  "name" varchar
);

CREATE TABLE "users" (
  "id" varchar UNIQUE NOT NULL,
  "first_name" varchar,
  "last_name" varchar,
  "phone_number" varchar NOT NULL,
  "is_verified" boolean DEFAULT 'false',
  "long" FLOAT,
  "lat" FLOAT,
  "user_type" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT NOW(),
  "updated_at" timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE "faq" (
  "id" varchar PRIMARY KEY NOT NULL,
  "question" varchar,
  "answer" varchar,
  "slug" varchar NOT NULL,
  "lang" varchar NOT NULL,
  "active" boolean DEFAULT 'false',
  "created_at" timestamp NOT NULL DEFAULT NOW(),
  "updated_at" timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE "category" (
  "id" varchar PRIMARY KEY NOT NULL,
  "parent_id" varchar,
  "image" varchar,
  "active" boolean DEFAULT 'true',
  "slug" varchar,
  "lang" varchar,
  "information" varchar,
  "created_at" timestamp,
  "updated_at" timestamp
);
