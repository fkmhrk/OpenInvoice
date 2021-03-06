create database if not exists openinvoice default character set utf8;

use openinvoice;

create table if not exists user(
  id varchar(32),
  login_name varchar(64),
  password varchar(32),
  display_name text,
  role text,
  tel text,
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
  title_type tinyint,
  work_from bigint,
  work_to bigint,
  total bigint,
  assignee varchar(32),
  product text,
  memo text,
  quotation_date bigint,
  quotation_number varchar(32),
  bill_date bigint,
  bill_number varchar(32),
  delivery_date bigint,
  delivery_number varchar(32),   
  tax_rate decimal(6,3),
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
  amount float,
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
  role text,
  expire_time bigint,
  created_time bigint,
  modified_time bigint,
  deleted tinyint,
  primary key(access_token)
) engine=InnoDB;

create table if not exists session_refresh(
  token varchar(48),
  user_id varchar(32),
  role text,
  expire_time bigint,
  created_time bigint,
  modified_time bigint,
  deleted tinyint,
  primary key(token)
) engine=InnoDB;

create table if not exists trading_id(
  date varchar(32),
  num int,
  primary key(date)
) engine=InnoDB;

create table if not exists env(
  id varchar(32),
  value text,
  created_time bigint,
  modified_time bigint,
  deleted tinyint,
  primary key(id)
) engine=InnoDB;

create table if not exists seq(
  seq_type tinyint,
  year int,
  value int,
  created_time bigint,
  modified_time bigint,
  deleted tinyint,  
  primary key(seq_type, year)
) engine=InnoDB;

/* default admin user */
insert into user(id,login_name,password,display_name,role,tel,created_time,modified_time,deleted)
 values('idadmin', 'admin', 'c8ea932352eae95d3ae36891cdf42eb4', 'admin', 'Admin','',unix_timestamp(), unix_timestamp(), 0);
