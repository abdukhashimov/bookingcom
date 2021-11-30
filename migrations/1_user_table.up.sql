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
  "slug" varchar,
  "lang" varchar,
  "active" bool DEFAULT 'false',
  "created_at" timestamp NOT NULL DEFAULT NOW(),
  "udpated_at" timestamp NOT NULL DEFAULT NOW()
);