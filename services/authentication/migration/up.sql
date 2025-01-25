DROP TABLE IF EXISTS "tokens";
DROP TABLE IF EXISTS "passwords";
DROP TABLE IF EXISTS "login_history";
DROP TABLE IF EXISTS "users";

-- Users table
CREATE TABLE "users" (
    "user_id" BIGSERIAL NOT NULL PRIMARY KEY,
    "phone_number" TEXT NOT NULL UNIQUE,
    "user_role" TEXT NOT NULL,
	"verify" Bool NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL
);

-- Tokens table
CREATE TABLE "tokens" (
    "access_token" TEXT NOT NULL PRIMARY KEY,
    "refresh_token" TEXT NOT NULL,
    "user_id" BIGINT NOT NULL,
    "user_role" TEXT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL,
    "access_token_expire_at" TIMESTAMPTZ NOT NULL,
    "refresh_token_expire_at" TIMESTAMPTZ NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);

-- Passwords table 
CREATE TABLE "passwords" (
    "user_id" BIGINT NOT NULL PRIMARY KEY REFERENCES users (user_id),
    "password" TEXT NOT NULL
);

