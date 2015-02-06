create database if not exists openinvoice default character set utf8;

use openinvoice;

create table if not exists user(
  id varchar(32),
  login_name varchar(64),
  password varchar(32),
  created_time bigint,
  modified_time bigint,
  deleted tinyint,
  primary key(id)
) engine=InnoDB;

create table if not exists company(
  id varchar(32),
  name text,
  zip varchar(16),
  address text,
  phone text,
  unit varchar(16),
  created_time bigint,
  modified_time bigint,
  deleted tinyint,
  primary key(id)
) engine=InnoDB;

create table if not exists trading(
  id varchar(32),
  company_id varchar(32),
  subject text,
  work_from bigint,
  work_to bigint,
  assignee varchar(32),
  product text,
  created_time bigint,
  modified_time bigint,
  deleted tinyint,
  primary key(id)
) engine=InnoDB;


create table if not exists trading_item(
  id varchar(32),
  trading_id varchar(32),
  sort_order tinyint,
  subject text,
  unit_price integer,
  amount integer,
  degree varchar(8),
  tax_type tinyint,
  memo text,
  created_time bigint,
  modified_time bigint,
  deleted tinyint,
  primary key(id, trading_id)
) engine=InnoDB;

create table if not exists settings(
  name varchar(32),
  value text,
  created_time bigint,
  modified_time bigint,
  deleted tinyint,
  primary key(name)
) engine=InnoDB;

create table if not exists session(
  access_token varchar(48),
  user_id varchar(32),
  scope text,
  expire_time bigint,
  created_time bigint,
  modified_time bigint,
  deleted tinyint,
  primary key(access_token)
) engine=InnoDB;

create table if not exists trading_id(
  date varchar(32),
  num int,
  primary key(date)
) engine=InnoDB;
