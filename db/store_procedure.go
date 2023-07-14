package db

const StoreProcedure = `

---- 角色 1个----

/*查找指定角色下所有的子角色
（结果集中包含指定角色）
*/
create or replace function query_child_roles(roleid uuid)
  returns setof role
  language plpgsql
as
$$
begin
	--- 递归查询指定角色下所有的子角色
  return query WITH RECURSIVE roles AS
    (SELECT *
     FROM role
     WHERE role.id = roleid
     UNION
     SELECT role.*
     FROM role,
          roles
     WHERE role.parent_id = roles.id
       AND role.delete_time IS NULL)
    SELECT *
    FROM roles
    ORDER BY create_time;
end;
$$;

---- 组织 2个----

/*查找指定组织下所有的子组织
（结果集中包含指定组织）
*/
create or replace function query_child_organizations(orgid uuid)
  returns setof organization
  language plpgsql
as
$$
begin
	--- 递归查询指定组织下所有的子组织
  return query WITH RECURSIVE orgs AS
    (SELECT *
     FROM organization
     WHERE organization.id = orgid
     UNION
     SELECT organization.*
     FROM organization,
          orgs
     WHERE organization.parent_id = orgs.id
       AND organization.delete_time IS NULL)
    SELECT *
    FROM orgs
    ORDER BY create_time;
end;
$$;

/*查询指定组织及其子组织下的所有用户ID集合*/
create or replace function query_users_of_child_organization(orgid uuid)
  returns setof uuid
   language plpgsql
as
$$
begin

  return query select u.id
               from "public"."user" u
               where u.organization_id in (
                 WITH RECURSIVE orgs AS
                                  (SELECT organization.id
                                   FROM organization
                                   WHERE organization.id = orgid
                                   UNION
                                   SELECT organization.id
                                   FROM organization,
                                        orgs
                                   WHERE organization.parent_id = orgs.id
                                     AND organization.delete_time IS NULL)
                 SELECT orgs.id
                 FROM orgs
               );
end;
$$;



----- 设备 2个----

/*批量创建设备负责人*/
create or replace function new_directors(deviceid uuid, userids anyarray) returns void
  language plpgsql
as
$$
declare
  uid        uuid;
  now        bigint := FLOOR(EXTRACT(epoch FROM ((CURRENT_TIMESTAMP - TIMESTAMP '1970-01-01 00:00:00') * 1000)));
begin
  if userids != array[null] then
    foreach uid in array userids
     loop
      --       忽略重复添加的设备负责人
      perform * from director where device_id = deviceid and director.user_id = uid;
      if FOUND = false
      then
		--       忽略不存在的用户ID
        perform * from public."user" where id = uid;
        if FOUND = true
        then
          insert into director (id, created, updated, deleted, device_id, user_id)
          values (uuid_generate_v4(), now, now, null, deviceid, uid);
        end if;
      end if;
    end loop;
 end if;
end;
$$;

/*分配设备*/
create or replace function distribute_devices(org_id uuid, device_ids anyarray)
  returns void
language plpgsql
as $$
declare d_id uuid;
begin
if device_ids != array[null] then
  foreach d_id in array device_ids loop
    update device set organization_id = org_id where id = d_id;
  end loop;
end if;
end;
$$;


/*绑定用户和设备的关系*/
create or replace function bind_devices(userid uuid, deviceids anyarray)
  returns void
  language plpgsql
as
$$
declare
  now      bigint := FLOOR(EXTRACT(epoch FROM ((CURRENT_TIMESTAMP - TIMESTAMP '1970-01-01 00:00:00') * 1000)));
  deviceid varchar;
begin
--- 废除旧的绑定关系
  delete from "user_device" where user_id = userid;
 if deviceids != array[null] then
  foreach deviceid in array deviceids
    loop
      insert into "user_device"(user_id, device_id, created) values (userid, deviceid::uuid, now);
    end loop;
 end if;	
end;
$$;

/*绑定用户和组织的关系*/
create or replace function bind_organizations(userid uuid, orgids anyarray)
  returns void
  language plpgsql
as
$$
declare
  now      bigint := FLOOR(EXTRACT(epoch FROM ((CURRENT_TIMESTAMP - TIMESTAMP '1970-01-01 00:00:00') * 1000)));
  orgid varchar;
begin
  --- 废除旧的绑定关系
  delete from "user_organization" where user_id = userid;
  if orgids != array[null] then
  foreach orgid in array orgids
    loop
      insert into "user_organization"(user_id, organization_id, created) values (userid, orgid::uuid, now);
    end loop;
  end if;
end;
$$;



/*
查询用户真正能看到的设备*/
create or replace function query_real_devices_of_user(userid uuid)
  returns setof uuid
  language plpgsql
as
$$
declare
  deviceid uuid;
  orgid    uuid;
begin

  --     获取用户所能查看的每一个设备
  for deviceid in select device_id from user_device where user_id = userid
    loop
      --       raise notice '%',deviceid;
      select device.organization_id into orgid from device where id = deviceid;
      -- 查询用户和设备所属的组织是否存在关联
      perform * from user_organization where organization_id = orgid and user_id = userid;
      --       如果存在关联，则说明设备能被用户看到
      if FOUND = true
      then
        return next deviceid;
      end if;

    end loop;
end;
$$;


/*
查询真正能看到设备的用户
*/
create or replace function query_real_users_of_device(deviceid uuid)
  returns setof uuid
  language plpgsql
as
$$
declare
  orgid  uuid;
  userid uuid;
begin
  --   获取设备的组织id
  select organization_id into orgid from device where id = deviceid;
  --   查询和设备绑定关系的用户id
  for userid in select user_id from user_device where device_id = deviceid
    loop
      --       查询设备的组织id和用户是否存在关系
      perform * from user_organization where organization_id = orgid and user_id = userid;
      --       如果存在，则说明用户能够查看到设备
      if found = true then
        return next userid;
      end if;
    end loop;
end;
$$;

`
