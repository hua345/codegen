-- auto-generated definition
create table code_project
(
  id             bigint auto_increment
  comment '主键',
  project_id     varchar(32)                           not null
  comment '项目Id',
  project_name   varchar(52)                           not null
  comment '项目名称',
  project_status varchar(16)                           null
  comment '项目状态',
  created_date   timestamp default current_timestamp() not null
  comment '创建时间',
  created_by     varchar(32)                           not null
  comment '创建者',
  updated_date   timestamp default current_timestamp() not null
  comment '更新时间',
  updated_by     varchar(32)                           not null
  comment '更新者',
  status         varchar(8)                            not null
  comment '活动状态(1:draft,2:active,3:delete)',
  biz_status     varchar(8)                            null
  comment '业务状态,预留',
  process_id     varchar(8)                            null
  comment '工作流Id,预留',
  remark         varchar(32)                           null
  comment '预留字段',
  entity         varchar(32)                           null
  comment '实体字段，预留',
  constraint code_project_id_uindex
  unique (id),
  constraint code_project_project_id_uindex
  unique (project_id)
)
  comment '项目表';

alter table code_project
  add primary key (id);

