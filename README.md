# beego_admin
beego开发的通用后台功能 email:445864742@qq.com

#功能包含
#### 0.登录验证
#### 1.权限管理系统
#### 2.菜单管理(绑定权限)
#### 3.文件上传/阿里云oss
#### 4.微信分享、获取用户信息


## 执行步骤
1.clone代码

<code>$ git clone git@github.com:ZQCard/beego_admin.git</code>

2.下载所需要模块

<code>$ go mod tidy</code>

3.建立个人配置文件

<code>$ touch ./beego_admin/conf/my.ini</code>

4.导入sql文件到数据库(table.sql)

5.执行程序

<code>$ go run main.go</code>


