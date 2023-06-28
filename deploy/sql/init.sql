/*
用户表创建
*/

CREATE TABLE IF NOT EXISTS "user" (
  "id" uuid PRIMARY KEY,
  "username" varchar(40) COLLATE "utf8mb3_general_ci",
  "phone" varchar(20) COLLATE "utf8mb3_general_ci" DEFAULT NULL,
  "password" varchar(40) COLLATE "utf8mb3_general_ci" DEFAULT NULL,
  "email" varchar(255) COLLATE "utf8mb3_general_ci" DEFAULT NULL,
  "avatar" BYTEA COLLATE "utf8mb3_general_ci" DEFAULT NULL,
  "create_time" date DEFAULT NULL,
  "update_time" date DEFAULT NULL,
  "delete_time" date DEFAULT NULL,
  "nick_name" varchar(40) COLLATE "utf8mb3_general_ci" DEFAULT NULL,
  "role_id" integer DEFAULT NULL,
  "organization_id" uuid NOT NULL,
  "status" integer DEFAULT NULL,
);