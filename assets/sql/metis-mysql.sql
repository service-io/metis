-- create a table for survey
drop table if exists `survey`;
create table `survey`
(
  `id`          bigint       not null auto_increment,
  `title`       varchar(255) not null comment '主题',
  `description` varchar(1000)         default null comment '描述',
  `start_at`    timestamp    not null default current_timestamp comment '开始时间',
  `end_at`      timestamp    not null default current_timestamp comment '结束时间',
  `status`      tinyint      not null default 0 comment '0: 发布 1: 暂存 2: 已结束 3: 已失效',
  `top`         tinyint      not null default 0 comment '0: 不置顶 1: 置顶',
  `creator_by`  bigint                default null comment '创建人员 ID',
  `updater_by`  bigint                default null comment '更新人员 ID',
  `created_at`  timestamp             default current_timestamp comment '创建时间',
  `updated_at`  timestamp             default current_timestamp on update current_timestamp comment '更新时间',
  primary key (`id`) using btree
) engine = innodb
  auto_increment = 100
  default charset = utf8mb4 comment ='调查问卷主表';

-- create a table for question
drop table if exists `question`;
create table `question`
(
  `id`          bigint       not null auto_increment,
  `title`       varchar(255) not null comment '问题主题',
  `description` varchar(1000)         default null comment '问题描述',
  `type`        tinyint      not null default 1 comment '1: 单选 2: 多选 3: 填空',
  `sort`        tinyint               default 0 comment '排序',
  `required`    tinyint               default 0 comment '0: 必填 1: 非必填',
  `extra`       varchar(2000)         default '' comment '附加数据',
  `creator_by`  bigint                default null comment '创建人员 ID',
  `updater_by`  bigint                default null comment '更新人员 ID',
  `created_at`  timestamp             default current_timestamp comment '创建时间',
  `updated_at`  timestamp             default current_timestamp on update current_timestamp comment '更新时间',
  primary key (`id`) using btree
) engine = innodb
  auto_increment = 100
  default charset = utf8mb4 comment ='问题主表';

-- create a table for option
drop table if exists `option`;
create table `option`
(
  `id`          bigint        not null auto_increment,
  `question_id` bigint        not null default 0 comment '问题 ID',
  `status`      tinyint       not null default 0 comment '0: 启用 1: 禁用',
  `name`        varchar(100)           default null comment '候选名',
  `content`     varchar(2000) not null default '' comment '候选内容',
  `answer`      tinyint       not null default 0 comment '是否是答案',
  `extra`       varchar(2000)          default '' comment '附加数据',
  `creator_by`  bigint                 default null comment '创建人员 ID',
  `updater_by`  bigint                 default null comment '更新人员 ID',
  `created_at`  timestamp              default current_timestamp comment '创建时间',
  `updated_at`  timestamp              default current_timestamp on update current_timestamp comment '更新时间',
  primary key (`id`) using btree
) engine = innodb
  auto_increment = 100
  default charset = utf8mb4 comment ='选项主表';

-- create a table for survey result
drop table if exists `survey_result`;
create table `survey_result`
(
  `id`        bigint not null auto_increment,
  `survey_id` bigint not null comment '关联调查问卷主表 ID',
  `ip`        varchar(40)   default '::7f00:0001' comment 'ip 地址',
  `agent`     varchar(300)  default '' comment '代理信息',
  `extra`     varchar(2000) default '' comment '附加数据',
  `result`    json          default ('{}') comment '结果数据',
  `answer_at` timestamp     default current_timestamp comment '回答时间',
  primary key (`id`) using btree
) engine = innodb
  auto_increment = 100
  default charset = utf8mb4 comment ='问卷结果主表';

-- create a table for survey template
drop table if exists `survey_template`;
create table `survey_template`
(
  `id`         bigint  not null auto_increment,
  `survey_id`  bigint  not null comment '关联调查问卷主表 ID',
  `type`       tinyint not null default 0 comment '0: PC 1: MOBILE 2: EMBED',
  `content`    text    not null default ('') comment '渲染模版',
  `creator_by` bigint           default null comment '创建人员 ID',
  `updater_by` bigint           default null comment '更新人员 ID',
  `created_at` timestamp        default current_timestamp comment '创建时间',
  `updated_at` timestamp        default current_timestamp on update current_timestamp comment '更新时间',
  primary key (`id`) using btree
) engine = innodb
  auto_increment = 100
  default charset = utf8mb4 comment ='问卷模版主表';

-- create a table for ship of survey and question
drop table if exists `survey_question_ship`;
create table `survey_question_ship`
(
  `survey_id`   bigint not null comment '关联调查问卷主表 ID',
  `question_id` bigint not null comment '关联问题主表 ID',
  unique key (`survey_id`, `question_id`) using btree
) engine = innodb
  default charset = utf8mb4 comment ='问卷-问题关系';
