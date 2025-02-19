#!/bin/bash
set -e

# 使用 root 用户执行的命令
psql -v ON_ERROR_STOP=1 --username "root" --dbname="postgres" <<-EOSQL
    CREATE DATABASE phalanx;
    CREATE ROLE phalanx_user WITH LOGIN PASSWORD 'ji394@Phalanx';
    GRANT ALL PRIVILEGES ON DATABASE phalanx TO phalanx_user;
    CREATE ROLE readwrite;
    GRANT CONNECT ON DATABASE phalanx TO readwrite;
EOSQL

# 切换到 phalanx 数据库，继续使用 root 用户执行
psql -v ON_ERROR_STOP=1 --username "root" --dbname="phalanx" <<-EOSQL
    GRANT USAGE, CREATE ON SCHEMA public TO readwrite;
    GRANT USAGE ON ALL SEQUENCES IN SCHEMA public TO readwrite;
    GRANT readwrite TO phalanx_user;
    -- 在此处添加其他针对 phalanx 数据库的命令
EOSQL

# 使用 phalanx_user 用户执行的命令
psql -v ON_ERROR_STOP=1 --username "phalanx_user" --dbname="phalanx" <<-EOSQL
    -- public.member definition
CREATE TABLE public.member (
    id bigserial NOT NULL,
    account varchar(64) NOT NULL,
    name varchar(128) NOT NULL,
    real_name varchar(128) NULL,
    email varchar(128) NOT NULL,
    spare_email varchar(128) NULL,
    mobile_phone varchar(16) NULL,
    city_phone varchar(16) NULL,
    city_phone_ext varchar(16) NULL,
    address varchar(16) NULL,
    is_enable bool NOT NULL,
    create_id int8 NOT NULL,
    create_time timestamp NOT NULL,
    modify_id int8 NOT NULL,
    modify_time timestamp NOT NULL,
    CONSTRAINT pk_member PRIMARY KEY (id)
);

-- public.menu definition
CREATE TABLE public.menu (
    id bigserial NOT NULL,
    title varchar(32) NOT NULL,
    icon varchar(32) NULL,
    url varchar(128) NOT NULL,
    parent int8 NULL,
    description varchar(128) NULL,
    sort int8 NULL,
    is_enable bool NOT NULL,
    is_show bool NOT NULL,
    create_id int8 NOT NULL,
    create_time timestamp NOT NULL,
    modify_id int8 NOT NULL,
    modify_time timestamp NOT NULL,
    CONSTRAINT pk_menu PRIMARY KEY (id)
);

-- public.resource_message definition
CREATE TABLE public.resource_message (
    id bigserial NOT NULL,
    message_key varchar(128) NOT NULL,
    message_value text NOT NULL,
    message_descr varchar(256) NOT NULL,
    create_id int8 NOT NULL,
    create_time timestamp NOT NULL,
    modify_id int8 NOT NULL,
    modify_time timestamp NOT NULL,
    CONSTRAINT resource_message_pkey PRIMARY KEY (id)
);

-- public.role definition
CREATE TABLE public.role (
    id bigserial NOT NULL,
    title varchar(32) NOT NULL,
    description varchar(128) NULL,
    sort int8 NULL,
    is_enable bool NOT NULL,
    create_id int8 NOT NULL,
    create_time timestamp NOT NULL,
    modify_id int8 NOT NULL,
    modify_time timestamp NOT NULL,
    CONSTRAINT pk_role PRIMARY KEY (id)
);

-- public.scan_tool definition
CREATE TABLE public.scan_tool (
    id bigserial NOT NULL,
    name varchar(64) NOT NULL,
    grpc_server varchar(64) NOT NULL,
    CONSTRAINT scan_tool_pkey PRIMARY KEY (id)
);

-- public.system_log definition
CREATE TABLE public.system_log (
    id serial4 NOT NULL,
    level varchar(128) NOT NULL,
    message json NOT NULL,
    timestamp timestamp NOT NULL,
    CONSTRAINT system_log_pkey PRIMARY KEY (id)
);

-- public.vulnerability definition
CREATE TABLE public.vulnerability (
    id bigserial NOT NULL,
    name varchar(64) NOT NULL,
    target json NULL,
    scan_type int4 NOT NULL,
    scan_status int4 NULL,
    description varchar(64) NULL,
    scan_time timestamp NOT NULL,
    create_id int8 NOT NULL,
    create_time timestamp NOT NULL,
    modify_id int8 NOT NULL,
    modify_time timestamp NOT NULL,
    CONSTRAINT vulnerability_pkey PRIMARY KEY (id)
);

-- public.forgot_temp definition
CREATE TABLE public.forgot_temp (
    id bigserial NOT NULL,
    member_id int8 NOT NULL,
    expire_time timestamp NOT NULL,
    code varchar(128) NOT NULL,
    redirect_path varchar(512) NULL,
    CONSTRAINT pk_forgot_temp PRIMARY KEY (id),
    CONSTRAINT fk_forgot_temp_member_id FOREIGN KEY (member_id) REFERENCES public.member(id)
);

-- public.member_history definition
CREATE TABLE public.member_history (
    id bigserial NOT NULL,
    member_id int8 NOT NULL,
    password varchar(128) NOT NULL,
    salt varchar(32) NOT NULL,
    error_count int2 NOT NULL DEFAULT 0,
    create_id int8 NOT NULL,
    create_time timestamp NOT NULL,
    modify_id int8 NOT NULL,
    modify_time timestamp NOT NULL,
    CONSTRAINT pk_member_history PRIMARY KEY (id),
    CONSTRAINT fk_member_history_member_id FOREIGN KEY (member_id) REFERENCES public.member(id)
);

-- public.member_role definition
CREATE TABLE public.member_role (
    id bigserial NOT NULL,
    role_id int8 NOT NULL,
    member_id int8 NOT NULL,
    CONSTRAINT pk_member_role PRIMARY KEY (id),
    CONSTRAINT fk_member_role_member_id FOREIGN KEY (member_id) REFERENCES public.member(id),
    CONSTRAINT fk_member_role_role_id FOREIGN KEY (role_id) REFERENCES public.role(id)
);

-- public.power definition
CREATE TABLE public.power (
    id bigserial NOT NULL,
    menu_id int8 NOT NULL,
    title varchar(32) NOT NULL,
    code varchar(32) NOT NULL,
    description varchar(128) NULL,
    sort int8 NULL,
    is_enable bool NOT NULL,
    create_id int8 NOT NULL,
    create_time timestamp NOT NULL,
    modify_id int8 NOT NULL,
    modify_time timestamp NOT NULL,
    CONSTRAINT pk_power PRIMARY KEY (id),
    CONSTRAINT fk_power_menu_id FOREIGN KEY (menu_id) REFERENCES public.menu(id)
);

-- public.role_power definition
CREATE TABLE public.role_power (
    id bigserial NOT NULL,
    role_id int8 NOT NULL,
    menu_id int8 NOT NULL,
    power_id int8 NULL,
    create_id int8 NOT NULL,
    create_time timestamp NOT NULL,
    modify_id int8 NOT NULL,
    modify_time timestamp NOT NULL,
    CONSTRAINT pk_role_power PRIMARY KEY (id),
    CONSTRAINT fk_role_power_menu_id FOREIGN KEY (menu_id) REFERENCES public.menu(id),
    CONSTRAINT fk_role_power_power_id FOREIGN KEY (power_id) REFERENCES public.power(id),
    CONSTRAINT fk_role_power_role_id FOREIGN KEY (role_id) REFERENCES public.role(id)
);

-- public.v_member_role source
CREATE OR REPLACE VIEW public.v_member_role
AS SELECT member_role.id,
    member_role.member_id,
    member_role.role_id,
    role.title AS role_title
   FROM public.member_role
     JOIN public.role ON role.id = member_role.role_id;

-- public.v_power source
CREATE OR REPLACE VIEW public.v_power
AS SELECT power.id,
    power.menu_id,
    menu.title AS menu_name,
    menu.sort AS menu_sort,
    power.title,
    power.code,
    power.description,
    power.sort,
    power.is_enable
   FROM public.power
     JOIN public.menu ON menu.id = power.menu_id
  ORDER BY menu.sort, power.sort;

-- public.v_role_power source
CREATE OR REPLACE VIEW public.v_role_power
AS SELECT role_power.id,
    role_power.role_id,
    role.title AS role_title,
    role_power.menu_id,
    menu.title AS menu_title,
    role_power.power_id,
    power.title AS power_title,
    power.code AS power_code
   FROM public.role_power
     JOIN public.role ON role.id = role_power.role_id
     JOIN public.menu ON menu.id = role_power.menu_id
     LEFT JOIN public.power ON power.id = role_power.power_id;

INSERT INTO member
(account, name, real_name, email, spare_email, mobile_phone, city_phone, city_phone_ext, address, is_enable, create_id, create_time, modify_id, modify_time)
VALUES
('system', 'SuperAdmin', 'SuperAdmin', 'system', '', '', '', '', '', true, 1, '2023-03-07 00:00:00.000', 1, '2023-04-19 04:30:30.412');

INSERT INTO member_history
(member_id, password, salt, error_count, create_id, create_time, modify_id, modify_time)
VALUES
(1, '84c67b3049a3eb8fd84e556bf05ad66cf151ea22c0211eb95be975cccacdcc016dbcd35a95b8530c75165dcff2ec8fdb197009a8539f466005961948130918cd', '453112bf6ef24884b069b8b1ff889f99', 0, 1, '2023-03-07 00:00:00.000', 1, '2023-03-07 00:00:00.000');

--

-- 2023/12/04 新增menu

INSERT INTO menu
(title, icon, url, parent, description, sort, is_enable, is_show, create_id, create_time, modify_id, modify_time)
VALUES('home', 'HomeIcon', '/', 0, '首頁', 0, true, true, 1, '2023-12-04', 1, '2023-12-04');
INSERT INTO "role"
(title, description, sort, is_enable, create_id, create_time, modify_id, modify_time)
VALUES('SystemAdmin', '', 0, true, 1, '2023-12-04', 1, '2023-12-04');
INSERT INTO role_power
(role_id, menu_id, power_id, create_id, create_time, modify_id, modify_time)
VALUES(1, 1, null, 1, '2023-12-04', 1, '2023-12-04');

INSERT INTO member_role
(role_id, member_id)
VALUES(1, 1);

-- 2023/12/08 新增menu資料
INSERT INTO menu
(title, icon, url, parent, description, sort, is_enable, is_show, create_id, create_time, modify_id, modify_time)
VALUES('task', 'TaskIcon', '/task', 0, '任務', 1, true, true, 1, '2023-12-04', 1, '2023-12-04');
INSERT INTO role_power
(role_id, menu_id, power_id, create_id, create_time, modify_id, modify_time)
VALUES(1, 2, null, 1, '2023-12-04', 1, '2023-12-04');

INSERT INTO menu
(title, icon, url, parent, description, sort, is_enable, is_show, create_id, create_time, modify_id, modify_time)
VALUES('reconnaissance', 'SatelliteAltIcon', '/reconnaissance', 2, '偵查', 2, true, true, 1, '2023-12-04', 1, '2023-12-04');
INSERT INTO role_power
(role_id, menu_id, power_id, create_id, create_time, modify_id, modify_time)
VALUES(1, 3, null, 1, '2023-12-04', 1, '2023-12-04');


INSERT INTO menu
(title, icon, url, parent, description, sort, is_enable, is_show, create_id, create_time, modify_id, modify_time)
VALUES('exploit', 'CoronavirusIcon', '/exploit', 2, '偵查', 2, true, true, 1, '2023-12-04', 1, '2023-12-04');
INSERT INTO role_power
(role_id, menu_id, power_id, create_id, create_time, modify_id, modify_time)
VALUES(1, 4, null, 1, '2023-12-04', 1, '2023-12-04');

INSERT INTO menu
(title, icon, url, parent, description, sort, is_enable, is_show, create_id, create_time, modify_id, modify_time)
VALUES('setting', 'SettingsIcon', '/setting', 0, '系統設定', 2, true, true, 1, '2023-12-22 05:31:14.126', 1, '2023-12-22 05:31:14.126');
INSERT INTO menu
(title, icon, url, parent, description, sort, is_enable, is_show, create_id, create_time, modify_id, modify_time)
VALUES('menu', 'ListIcon', '/menu', 5, '選單管理', 1, true, true, 1, '2023-12-22 05:31:39.477', 1, '2023-12-22 05:31:39.477');
INSERT INTO menu
(title, icon, url, parent, description, sort, is_enable, is_show, create_id, create_time, modify_id, modify_time)
VALUES('power', 'PolicyIcon', '/power', 5, '權限管理', 2, true, true, 1, '2023-12-22 05:31:47.418', 1, '2023-12-22 05:31:47.418');
INSERT INTO menu
(title, icon, url, parent, description, sort, is_enable, is_show, create_id, create_time, modify_id, modify_time)
VALUES('user', 'PeopleIcon', '/user', 5, '用戶管理', 4, true, true, 1, '2023-12-25 06:15:41.929', 1, '2023-12-25 06:15:41.929');
INSERT INTO menu
(title, icon, url, parent, description, sort, is_enable, is_show, create_id, create_time, modify_id, modify_time)
VALUES('role', 'PersonIcon', '/role', 5, '角色管理', 3, true, true, 1, '2023-12-25 06:16:13.451', 1, '2023-12-25 06:16:13.451');


INSERT INTO role_power
(role_id, menu_id, power_id, create_id, create_time, modify_id, modify_time)
VALUES(1, 5, NULL, 1, '2023-12-04 00:00:00.000', 1, '2023-12-04 00:00:00.000');
INSERT INTO role_power
(role_id, menu_id, power_id, create_id, create_time, modify_id, modify_time)
VALUES(1, 6, NULL, 1, '2023-12-04 00:00:00.000', 1, '2023-12-04 00:00:00.000');
INSERT INTO role_power
(role_id, menu_id, power_id, create_id, create_time, modify_id, modify_time)
VALUES(1, 7, NULL, 1, '2023-12-04 00:00:00.000', 1, '2023-12-04 00:00:00.000');
INSERT INTO role_power
(role_id, menu_id, power_id, create_id, create_time, modify_id, modify_time)
VALUES(1, 8, NULL, 1, '2023-12-04 00:00:00.000', 1, '2023-12-04 00:00:00.000');
INSERT INTO role_power
(role_id, menu_id, power_id, create_id, create_time, modify_id, modify_time)
VALUES(1, 9, NULL, 1, '2023-12-04 00:00:00.000', 1, '2023-12-04 00:00:00.000');


INSERT INTO "role"
(title, description, sort, is_enable, create_id, create_time, modify_id, modify_time)
VALUES('User', '', 0, true, 1, '2023-12-04', 1, '2023-12-04');

INSERT INTO role_power
(role_id, menu_id, power_id, create_id, create_time, modify_id, modify_time)
VALUES(2, 1, NULL, 1, '2023-12-04 00:00:00.000', 1, '2023-12-04 00:00:00.000');
INSERT INTO role_power
(role_id, menu_id, power_id, create_id, create_time, modify_id, modify_time)
VALUES(2, 2, NULL, 1, '2023-12-04 00:00:00.000', 1, '2023-12-04 00:00:00.000');
INSERT INTO role_power
(role_id, menu_id, power_id, create_id, create_time, modify_id, modify_time)
VALUES(2, 3, NULL, 1, '2023-12-04 00:00:00.000', 1, '2023-12-04 00:00:00.000');
INSERT INTO role_power
(role_id, menu_id, power_id, create_id, create_time, modify_id, modify_time)
VALUES(2, 4, NULL, 1, '2023-12-04 00:00:00.000', 1, '2023-12-04 00:00:00.000');


create table flow_control(
	id bigserial PRIMARY key NOT NULL,
	condition varchar(64) not null,
	next_flow  varchar(64) not null,
	sort int8 NULL
);


CREATE TABLE authorization_token (
    id bigserial PRIMARY key NOT null,
    mac varchar(256) not NULL,
    authorization_token varchar(256) NOT NULL,
    start_time timestamptz NOT NULL,
    end_time timestamptz NOT NULL,
    used int4 null,
    usage_count int4 null
);


CREATE TABLE setting (
	id bigserial PRIMARY key NOT null,
	setting json NOT NULL
);
INSERT INTO setting
(setting)
VALUES('{"externalIp":{"type":1,"ip":"192.168.70.3","gateway":"192.168.70.1"}}');

CREATE TABLE usage_record (
	id serial4 NOT NULL,
	duration float8 NOT NULL,
	CONSTRAINT usage_record_pkey PRIMARY KEY (id)
);

EOSQL
