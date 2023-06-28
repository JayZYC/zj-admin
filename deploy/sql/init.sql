/*
用户表创建
*/
DROP TABLE  IF NOT EXISTS "user";
CREATE TABLE "user" (
    "id" uuid PRIMARY KEY,
    "username" varchar(40),
    "phone" varchar(20) DEFAULT NULL,
    "password" varchar(40)  DEFAULT NULL,
    "email" varchar(255) DEFAULT NULL,
    "avatar" BYTEA  DEFAULT NULL,
    "create_time" date DEFAULT NULL,
    "update_time" date DEFAULT NULL,
    "delete_time" date DEFAULT NULL,
    "nick_name" varchar(40)  DEFAULT NULL,
    "role_id" integer DEFAULT NULL,
    "status" integer DEFAULT NULL
);

INSERT INTO "user" VALUES (uuid_generate_v4(), 'root', '17314943215', '123456', NULL, NULL, now(), now(), NULL, '超级管理员', 1, 1);

/*
角色表创建
*/

CREATE TABLE IF NOT EXISTS "role" (
    "id" uuid PRIMARY KEY,
    "create_time" date DEFAULT NULL,
    "update_time" date DEFAULT NULL,
    "delete_time" date DEFAULT NULL,
    "name" varchar(40)  DEFAULT NULL,
    "access" integer[] DEFAULT NULL
    );