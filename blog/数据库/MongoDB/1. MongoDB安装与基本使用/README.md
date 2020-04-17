

# MongoDB 安装与基本使用

[官方文档](https://docs.mongodb.com/manual/administration/install-community/)



[MongoDB版本与系统平台环境要求](https://docs.mongodb.com/manual/administration/production-notes/#prod-notes-supported-platforms)

越新的MongoDB版本，要求的系统环境越高。旧的MongoDB 版本可能在新的系统版本下无法运行。所以需要查看版本与系统环境的兼容。



### Mac OS 环境下，使用 homebrew 包管理工具进行安装。

#### 1. 需要 Xcode

如果没有安装，复制下面命令在 mac os 终端里执行安装 Xcode

```shell
xcode-select --install
```



#### 2. 需要 Homebrew

[安装 Homebrew 官方文档](https://brew.sh/#install)

复制下面命令在 mac os 终端里执行

```shell
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
```



复制下面命令在 mac os 终端里执行，指定mongodb 的 brew 仓库

```shell
brew tap mongodb/brew
```



#### 3. 安装 Mongodb

*以 4.2 版本为例*

复制下面命令在 mac os 终端里执行，`brew install` 后面跟包名字，`@`后面指定包版本

```
brew install mongodb-community@4.2
```



> 如果已经安装过其他版本，可能会出现 ChecksumMismatchError 错误，解决办法 [Troubleshooting ChecksumMismatchError](https://docs.mongodb.com/manual/tutorial/install-mongodb-on-os-x/#troubleshooting-checksumerror)。



![image-20200417154356619](./images/brew-install-mongodb-community@4.2.png)



安装完成，启动。

手动启动：`mongod --config /usr/local/etc/mongod.conf` 手动启动

启动 MondoDB 后台服务： `new services start mongoldb/brew/mongodb-community` 

停止 MongoDB 服务：`brew services stop mongodb-community@4.2` 

重启 MongoDB 服务：`brew services restart mongodb-community@4.2`



安装后会生成下列默认文件：

默认配置文件：`/usr/local/etc/mongod.conf`

默认日志目录路径：`/usr/local/var/log/mongodb`

默认数据目录路径：`/usr/local/var/mongodb`



#### 4. 其他问题：

1. 如果提示权限不足（Permission denied，关键字）

一般是对应目录没有权限，使用 `sudo chmod -R 777 目录` 修改目录权限，例如：`sudo chmod -R 777 /usr/local/var/mongodb`。

或尝试在命令前面加上 `sudo`，例如： `sudo mongod --config /usr/local/etc/mongod.conf`。



2. 使用 `mongod` 或 `sudo mongod` 命令直接启动报错

因为在 Mac OS 10.15 以上版本，不能创建 `/data/db` 目录，所以最好改成 `mongod --config /usr/local/etc/mongod.conf` 启动



3. 如果原来安装过旧版本的 Mongodb，启动时可能报下面错误：

```
2020-04-17T15:50:50.046+0800 F  CONTROL  [initandlisten] ** IMPORTANT: UPGRADE PROBLEM: Found an invalid featureCompatibilityVersion document (ERROR: BadValue: Invalid value for version, found 3.6, expected '4.2' or '4.0'. Contents of featureCompatibilityVersion document in admin.system.version: { _id: "featureCompatibilityVersion", version: "3.6" }. See http://dochub.mongodb.org/core/4.0-feature-compatibility.). If the current featureCompatibilityVersion is below 4.0, see the documentation on upgrading at http://dochub.mongodb.org/core/4.0-upgrade-fcv.

```

![image-20200417170520265](./images/error-fearuteCompatibilityVersion.png)



解决方式：

只需将数据目录路径 `/usr/local/var/mongodb` 里面内容清空在重新启动 mongod 即可。

数据是原来旧版本留下的，所以导致报错。



