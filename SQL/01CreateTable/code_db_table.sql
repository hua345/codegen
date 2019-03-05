-- auto-generated definition
create table code_db_table
(
  id           bigint auto_increment
  comment '主键',
  table_id     varchar(32)                           not null
  comment '数据库表Id',
  table_name   varchar(64)                           not null
  comment '数据库表名称',
  project_id   varchar(32)                           not null
  comment '所属项目Id',
  created_date timestamp default current_timestamp() not null
  comment '创建时间',
  created_by   varchar(32)                           not null
  comment '创建者',
  updated_date timestamp default current_timestamp() not null
  comment '更新时间',
  updated_by   varchar(32)                           not null
  comment '更新者',
  status       varchar(8)                            not null
  comment '活动状态(1:draft,2:active,3:delete)',
  biz_status   varchar(8)                            null
  comment '业务状态,预留',
  process_id   varchar(8)                            null
  comment '工作流Id,预留',
  remark       varchar(32)                           null
  comment '预留字段',
  entity       varchar(32)                           null
  comment '实体字段，预留',
  constraint id
  unique (id),
  constraint table_id
  unique (table_id)
)
  comment '数据库表';

alter table code_db_table
  add primary key (id);