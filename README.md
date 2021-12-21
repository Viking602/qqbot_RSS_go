# QQ订阅机器人

## 简介
目前还是一个非常简陋的版本，部分功能还在开发中

如果有好的建议欢迎提issues

计划未来增加web页面

## 使用方法
### 1.所需环境
#### 数据库
MySQL 5.7 

创建数据库后运行table.sql导入所需数据表

创建后运行以下语句添加bot
```mysql
insert into bot_info (botuid, botname, owneruid, status, createdatetime) value ('botQQ账号', '名称', '管理者QQ', 1, now());
```
### 2.配置相关

#### config.yaml
```yaml
mysql:
  # 数据库用户名
  user :
  # 数据库密码
  password :
  # 数据库链接地址
  host :
  # 数据库端口
  port :
  # 数据库名称
  db :

qq_bot:
  # 启动端口
  port :
  # 是否启用环境变量配置 Docker环境下推荐启动此配置（默认为环境变量配置）
  env: true

```
#### 环境变量

```shell
数据库地址
export RSS_MYSQL_HOST=

数据库端口
export RSS_MYSQL_PORT=

数据库名称
export RSS_MYSQL_NAME=

数据库用户名
export RSS_MYSQL_USER=

数据库密码
export RSS_MYSQL_PASSWORD=

启动端口
export RSS_BOT_PORT=
```
### 3.Docker部署(推荐)
下载镜像
```shell
docker pull viking602/qqbot-rss
```
启动容器
```shell
docker run --name qqbot-rss -p 8080:8080 -d -e RSS_MYSQL_HOST="数据库连接地址" -e RSS_MYSQL_PORT="端口" -e RSS_MYSQL_NAME="数据库名称" -e RSS_MYSQL_USER="用户名" -e RSS_MYSQL_PASSWORD="密码" -e RSS_BOT_PORT="启动端口" viking602/qqbot-rss
```

### 机器人配置
推荐使用[go-cqhttp](https://github.com/Mrs4s/go-cqhttp)
#### 配置方法
添加一个反向websocket连接方式
```yaml
  - ws-reverse:
      # 反向WS Universal 地址
      # 注意 设置了此项地址后下面两项将会被忽略
      # 注意这里替换成自己的地址
      universal: ws://127.0.0.1:8080/ws
      # 反向WS API 地址
      api: ""
      # 反向WS Event 地址
      event: ""
      # 重连间隔 单位毫秒
      reconnect-interval: 3000
      middlewares:
        <<: *default # 引用默认中间件
```

### 4.指令相关
rss-help 帮助信息

rss-all    查询本群订阅信息

rss-about    关于

rss-status    运行状态

rss-add RSS格式URL    添加RSS订阅

rss-live 房间号(暂时仅支持B站)    添加直播间订阅

rss-del 订阅名称    删除RSS订阅

rss-live-del 房间号    删除直播订阅