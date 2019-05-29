-- auto-generated definition
create table code_api_field
(
    id           bigint auto_increment
        comment '主键',
    field_id     varchar(32)  not null
        comment '字段Id',
    field_name   varchar(32)  not null
        comment '字段名称',
    field_type   varchar(32)  not null
        comment '字段类型',
    api_id       varchar(32)  not null
        comment '所属接口Id',
    field_desc   varchar(108) null
        comment '字段描述',
    request_type varchar(1)   not null
        comment '0;请求字段,1:响应字段',
    parent_id    varchar(32)  null
        comment '父字段Id, 不为Null则为子DTO字段',
    created_date timestamp    not null
        comment '创建时间',
    created_by   varchar(32)  not null
        comment '创建者',
    updated_date timestamp    not null
        comment '更新时间',
    updated_by   varchar(32)  not null
        comment '更新者',
    status       varchar(8)   not null
        comment '活动状态(1:draft,2:active,3:delete)',
    biz_status   varchar(8)   null
        comment '业务状态,预留',
    process_id   varchar(8)   null
        comment '工作流Id,预留',
    remark       varchar(32)  null
        comment '预留字段',
    entity       varchar(32)  null
        comment '实体字段，预留',
    constraint code_api_field_field_id_uindex
        unique (field_id),
    constraint id
        unique (id)
)
    comment '接口表';

alter table code_api_field
    add primary key (id);

