/*
 用户表创建
 */
DROP TABLE IF EXISTS "user";

CREATE TABLE "user" (
    "id" uuid PRIMARY KEY,
    "username" varchar(40) NOT NULL UNIQUE,
    "phone" varchar(20) NOT NULL UNIQUE,
    "password" varchar(40) DEFAULT NULL,
    "email" varchar(255) DEFAULT NULL,
    "avatar" varchar(40) DEFAULT NULL,
    "create_time" TIMESTAMPTZ DEFAULT NULL,
    "update_time" TIMESTAMPTZ DEFAULT NULL,
    "deletedAt" TIMESTAMPTZ DEFAULT NULL,
    "nick_name" varchar(40) DEFAULT NULL,
    "role_id" uuid DEFAULT NULL,
    "status" integer DEFAULT 1,
    "organization_id" uuid DEFAULT NULL
);

comment on table "user" is '用户表';

comment on column "user".avatar is '用户头像';

comment on column "user".nick_name is '用户昵称';

comment on column "user".status is '用户状态((1:正常;2:禁用))';

INSERT INTO
    "user"
VALUES
    (
        uuid_generate_v4(),
        'root',
        '17314943215',
        'E10ADC3949BA59ABBE56E057F20F883E',
        NULL,
        NULL,
        now(),
        now(),
        NULL,
        '超级管理员',
        uuid('5138e5da-4161-485a-bd1f-5f5a10de2f80'),
        1,
        NULL
    );

INSERT INTO
    "user"
VALUES
    (
        uuid_generate_v4(),
        'admin',
        '18994326818',
        'E10ADC3949BA59ABBE56E057F20F883E',
        NULL,
        NULL,
        now(),
        now(),
        NULL,
        '管理员',
        uuid('166108d4-75bf-4688-b2e6-1643a2df8d03'),
        1,
        uuid('7cdb912c-05db-4078-8625-2c3b3f8ee2b4')
    );

/*
 角色表创建
 */
DROP TABLE IF EXISTS "role";

CREATE TABLE IF NOT EXISTS "role" (
    "id" uuid PRIMARY KEY,
    "create_time" TIMESTAMPTZ DEFAULT NULL,
    "update_time" TIMESTAMPTZ DEFAULT NULL,
    "deletedAt" TIMESTAMPTZ DEFAULT NULL,
    "name" varchar(40) DEFAULT NULL,
    "access" integer[] DEFAULT NULL,
    "parent_id" uuid DEFAULT NULL
);

comment on column role.access is '菜单和按钮权限集合';

INSERT INTO
    "role"
VALUES
    (uuid('5138e5da-4161-485a-bd1f-5f5a10de2f80'), now(), now(), NULL, '超级管理员', NULL, NULL);

INSERT INTO
    "role"
VALUES
    (uuid('166108d4-75bf-4688-b2e6-1643a2df8d03'), now(), now(), NULL, '管理员', ARRAY[1,2,3,4,5,6,7], uuid('5138e5da-4161-485a-bd1f-5f5a10de2f80'));

INSERT INTO
    "role"
VALUES
    (
        uuid('0592b9d9-b8d6-45e6-86eb-2a441da0a5b9'),
        now(),
        now(),
        NULL,
        '用户',
        ARRAY [1,2,3,4,5,6,7],
        2,
        uuid('5138e5da-4161-485a-bd1f-5f5a10de2f80')
    );

/*
 菜单权限表创建
 */
DROP TABLE IF EXISTS "role_perm";

CREATE TABLE "role_perm" (
    "id" serial PRIMARY KEY,
    "parent_id" bigint,
    "name" varchar(64) NOT NULL DEFAULT '',
    "type" SMALLINT NOT NULL,
    "path" varchar(128) NULL DEFAULT '',
    "component" varchar(128) NULL DEFAULT NULL,
    "perm" varchar(128) NULL DEFAULT NULL,
    "visible" bool NOT NULL DEFAULT true,
    "sort" SMALLINT NULL DEFAULT 0,
    "icon" varchar(64) NULL DEFAULT '',
    "redirect" varchar(128) NULL DEFAULT NULL,
    "create_time" TIMESTAMPTZ DEFAULT NULL,
    "update_time" TIMESTAMPTZ DEFAULT NULL,
    "deletedAt" TIMESTAMPTZ DEFAULT NULL
);

comment on table role_perm is '菜单管理';

comment on column role_perm.parent_id is '父菜单ID';

comment on column role_perm.name is '菜单名称';

comment on column role_perm.type is '菜单类型(1:菜单；2:目录；3:外链；4:按钮)';

comment on column role_perm.path is '路由路径(浏览器地址栏路径)';

comment on column role_perm.component is '组件路径(vue页面完整路径，省略.vue后缀)';

comment on column role_perm.perm is '权限标识';

comment on column role_perm.visible is '显示状态(true-显示;false-隐藏)';

comment on column role_perm.sort is '排序';

comment on column role_perm.icon is '菜单图标';

comment on column role_perm.redirect is '跳转路径';

comment on column role_perm.create_time is '创建时间';

comment on column role_perm.update_time is '更新时间';

INSERT INTO
    "role_perm"
VALUES
    (
        DEFAULT,
        0,
        '系统管理',
        2,
        '/system',
        'Layout',
        NULL,
        true,
        1,
        'system',
        '/system/user',
        now(),
        now()
    );

INSERT INTO
    "role_perm"
VALUES
    (
        DEFAULT,
        1,
        '用户管理',
        1,
        '/system/user',
        'system/user/index',
        NULL,
        true,
        1,
        NULL,
        NULL,
        now(),
        now()
    );

INSERT INTO
    "role_perm"
VALUES
    (
        DEFAULT,
        1,
        '角色管理',
        1,
        '/system/role',
        'system/role/index',
        NULL,
        true,
        1,
        NULL,
        NULL,
        now(),
        now()
    );

INSERT INTO
    "role_perm"
VALUES
    (
        DEFAULT,
        1,
        '组织管理',
        1,
        '/system/organization',
        'system/organization/index',
        NULL,
        true,
        1,
        NULL,
        NULL,
        now(),
        now()
    );

INSERT INTO
    "role_perm"
VALUES
    (
        DEFAULT,
        2,
        '编辑用户按钮',
        4,
        NULL,
        NULL,
        'sys:user:edit',
        true,
        NULL,
        NULL,
        NULL,
        now(),
        now()
    );

    INSERT INTO
    "role_perm"
VALUES
    (
        DEFAULT,
        2,
        '添加用户按钮',
        4,
        NULL,
        NULL,
        'sys:user:add',
        true,
        NULL,
        NULL,
        NULL,
        now(),
        now()
    );

    INSERT INTO
    "role_perm"
VALUES
    (
        DEFAULT,
        2,
        '删除用户按钮',
        4,
        NULL,
        NULL,
        'sys:user:delete',
        true,
        NULL,
        NULL,
        NULL,
        now(),
        now()
    );

/*
 组织表创建
 */
DROP TABLE IF EXISTS "organization";

CREATE TABLE "organization" (
    "id" uuid PRIMARY KEY,
    "parent_id" uuid,
    "name" varchar(800) NOT NULL DEFAULT '',
    "province" varchar(128),
    "city" varchar(128),
    "district" varchar(128),
    "address" varchar(128),
    "create_time" TIMESTAMPTZ DEFAULT NULL,
    "update_time" TIMESTAMPTZ DEFAULT NULL,
    "deletedAt" TIMESTAMPTZ DEFAULT NULL
);

comment on column organization.parent_id is '父组织ID';

comment on column organization.name is '组织名称';

comment on column organization.province is '省份';

comment on column organization.city is '城市';

comment on column organization.district is '市区';

comment on column organization.create_time is '创建时间';

comment on column organization.update_time is '更新时间';

INSERT INTO
    "organization"
VALUES
    (
        uuid_generate_v4(),
        NULL,
        '总公司',
        '江苏省',
        '苏州市',
        '吴江区',
        '',
        now(),
        now(),
        NULL
    );

INSERT INTO
    "organization"
VALUES
    (
        uuid_generate_v4(),
        uuid('2836ba22-3a0b-4ff3-9ce5-9b6d4d7b7739'),
        '上海分公司',
        '上海市',
        '上海市',
        '青浦区',
        '',
        now(),
        now(),
        NULL
    );