/*
设置时区为东八区
*/
set time zone 'PRC'

/*
用户表创建
*/
DROP TABLE IF EXISTS "user";
CREATE TABLE "user" (
    "id" uuid PRIMARY KEY,
    "username" varchar(40),
    "phone" varchar(20) NOT NULL UNIQUE,
    "password" varchar(40)  DEFAULT NULL,
    "email" varchar(255) DEFAULT NULL,
    "avatar" BYTEA  DEFAULT NULL,
    "create_time" TIMESTAMPTZ DEFAULT NULL,
    "update_time" TIMESTAMPTZ DEFAULT NULL,
    "delete_time" TIMESTAMPTZ DEFAULT NULL,
    "nick_name" varchar(40)  DEFAULT NULL,
    "role_id" integer DEFAULT NULL,
    "status" integer DEFAULT NULL
);

comment on table "user" is '用户表';
comment on column "user".avatar is '用户头像';
comment on column "user".nick_name is '用户昵称';
comment on column "user".status is '用户状态((1:正常;2:禁用))';

INSERT INTO "user" VALUES (uuid_generate_v4(), 'root', '17314943215', '123456', NULL, NULL, now(), now(), NULL, '超级管理员', 1, 1);

/*
角色表创建
*/

CREATE TABLE IF NOT EXISTS "role" (
    "id" uuid PRIMARY KEY,
    "create_time" TIMESTAMPTZ DEFAULT NULL,
    "update_time" TIMESTAMPTZ DEFAULT NULL,
    "delete_time" TIMESTAMPTZ DEFAULT NULL,
    "name" varchar(40)  DEFAULT NULL,
    "access" integer[] DEFAULT NULL
    );

/*
菜单权限表创建
*/
DROP TABLE IF EXISTS "role_perm";
CREATE TABLE "role_perm"  (
                             "id" bigserial NOT NULL ,
                             "parent_id" bigint ,
                             "name" varchar(64) NOT NULL DEFAULT '' ,
                             "type" SMALLINT NOT NULL ,
                             "path" varchar(128)  NULL DEFAULT '' ,
                             "component" varchar(128)  NULL DEFAULT NULL ,
                             "perm" varchar(128)  NULL DEFAULT NULL ,
                             "visible" SMALLINT NOT NULL DEFAULT 1 ,
                             "sort" SMALLINT NULL DEFAULT 0 ,
                             "icon" varchar(64)  NULL DEFAULT '' ,
                             "redirect" varchar(128)  NULL DEFAULT NULL ,
                             "create_time" TIMESTAMPTZ NULL DEFAULT NULL,
                             "update_time" TIMESTAMPTZ NULL DEFAULT NULL 
);

comment on table role_perm is '菜单管理';
comment on column role_perm.parent_id is '父菜单ID';
comment on column role_perm.name is '菜单名称';
comment on column role_perm.type is '菜单类型(1:菜单；2:目录；3:外链；4:按钮)';
comment on column role_perm.path is '路由路径(浏览器地址栏路径)';
comment on column role_perm.component is '组件路径(vue页面完整路径，省略.vue后缀)';
comment on column role_perm.perm is '权限标识';
comment on column role_perm.visible is '显示状态(1-显示;0-隐藏)';
comment on column role_perm.sort is '排序';
comment on column role_perm.icon is '菜单图标';
comment on column role_perm.redirect is '跳转路径';
comment on column role_perm.create_time is '创建时间';
comment on column role_perm.update_time is '更新时间';

INSERT INTO "role_perm" VALUES (1, 0, '系统管理', 2, '/system', 'Layout', NULL, 1, 1, 'system', '/system/user', now(), now());