create table bot_info
(
    BotId          int auto_increment
        primary key,
    BotUid         varchar(255)                              null,
    BotName        varchar(255)                              null,
    OwnerUid       varchar(255)                              null,
    Status         int(1) unsigned default 0                 null,
    CreateDateTime datetime        default CURRENT_TIMESTAMP null
);

create table group_info
(
    GroupId        int auto_increment
        primary key,
    GroupCode      int           null,
    BotId          int           null,
    Status         int default 0 null,
    CreateDateTime datetime      null
);

create table msg_info
(
    Id             int auto_increment
        primary key,
    GroupId        int          null,
    Title          varchar(255) null,
    CreateDateTime datetime     null
);

create table room_info
(
    RoomId         int auto_increment
        primary key,
    RoomCode       int(255)     null comment '房间号',
    BotId          int(255)     null comment '机器人ID',
    GroupId        int(255)     null comment '群ID',
    RssTypeId      int(255)     null comment '类型',
    Status         int(1)       null comment '状态',
    CreateDateTime datetime     null comment '创建时间',
    RoomName       varchar(255) null,
    CreateUserId   bigint       null
);

create table rss_type
(
    RssTypeId      int auto_increment
        primary key,
    TypeName       varchar(255) null,
    CreateDateTime datetime     null
);

create table send_info
(
    Id             int auto_increment
        primary key,
    Url            varchar(255) null,
    MsgInfo        varchar(255) null,
    GroupId        int          null,
    CreateDateTime datetime     null,
    UpdateTime     datetime     null
);

create table url_info
(
    Id             int auto_increment
        primary key,
    Url            varchar(255)     null,
    UrlName        varchar(255)     null,
    Status         int(1) default 0 null,
    GroupId        int              null,
    BotId          int              null,
    CreateDateTime datetime         null,
    RssTypeId      int              null comment '订阅类型',
    CreateUserId   bigint           null
);

