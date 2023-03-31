## 1. 基本介绍

### 1.1 项目介绍

> asynq-scheduled是一个基于 [Asynq](https://github.com/hibiken/asynq)开发的计划任务调度程序，通过数据库读取已配置任务，实现动态管理。

[在线预览](http://demo.fyly.cc:7202/): 测试后台

测试用户名：admin

测试密码：123456

## 2. 使用说明

```
- node版本 > v18.8
- IDE推荐：Visual Studio Code
```
> 注：scheduled、worker使用Redis实现消息代理，统一配置。

### 2.1 scheduled项目
Asynq任务生产者，只能单节点部署。

```bash
# 克隆项目
git clone https://github.com/wd16yuan/asynq-scheduled.git
# 进入scheduled文件夹
cd scheduled

# 使用 go mod 并安装go依赖包
go generate

# 编译 
go build -o scheduled main.go

# 运行二进制
./scheduled 
```

### 2.2 worker项目
Asynq任务消费者，可多节点部署。

```bash
# 克隆项目
git clone https://github.com/wd16yuan/asynq-scheduled.git
# 进入worker文件夹
cd worker

# 使用 go mod 并安装go依赖包
go generate

# 编译 
go build -o worker main.go

# 运行二进制
./worker 
```