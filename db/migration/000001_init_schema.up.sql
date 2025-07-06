CREATE TABLE "users" (
  "user_id" serial PRIMARY KEY,
  "email" varchar(255) UNIQUE NOT NULL,
  "firstname" varchar(64) NOT NULL,
  "lastname" varchar(64) NOT NULL,
  "email_verified" bool DEFAULT false,
  "hashed_password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE "refresh_tokens" (
  "refresh_token_id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" bool NOT NULL DEFAULT false,
  "expired_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE "verify_emails" (
  "email" varchar(255) UNIQUE NOT NULL,
  "token" varchar NOT NULL,
  "is_used" bool NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL
);

CREATE TABLE "urls" (
  "id" serial PRIMARY KEY,
  "original_url" varchar NOT NULL,
  "short_code" varchar(255) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'NOW()',
  "expires_at" timestamptz,
  "total_click" int NOT NULL default 0,
  "last_click_at" timestamptz,
  "unique_vistor" int NOT NULL default 0,
  "user_id" bigint
);

CREATE INDEX ON "users" USING BTREE ("email");

CREATE INDEX ON "refresh_tokens" USING BTREE ("user_id");

CREATE INDEX ON "verify_emails" USING BTREE ("email");

ALTER TABLE "refresh_tokens" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "verify_emails" ADD FOREIGN KEY ("email") REFERENCES "users" ("email") ON DELETE CASCADE ON UPDATE NO ACTION;
