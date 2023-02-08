mysql密码 root  rocketyou1@mysql
mysql workbench -> server -> users and privileges

每个项目需要单独的go.mod ,然后再引用的时候都需要以go.mod里面的路径为基础


[client]
host     = localhost
user     = debian-sys-maint
password = pYIwwGTJySoz4vOL
socket   = /var/run/mysqld/mysqld.sock
[mysql_upgrade]
host     = localhost
user     = debian-sys-maint
password = pYIwwGTJySoz4vOL
socket   = /var/run/mysqld/mysqld.sock

root = 'Rocketyou1rocketyou1#ibm'

 mysql -u root -p 登录不了，需要用sudo mysql -u root -p
 select user, plugin from mysql.user  用来查root的password是生成的还是自然字符

把密码改成自然的
 update mysql.user set authentication_string='Rocketyou1rocketyou1#ibm' , plugin= 'mysql_native_password' where user = 'root'

 show databases;
 show tables from XXXDB;