package template

var Yaml = `project_name: {{.Name}}
http_port: 8088
grpc_port: 9988
registry:  172.18.15.182:2379
qps: 100
#debug or prod
run_mode: debug
cos_root_path: /
log:
  path: ./log/
  name: app
  level: debug #日志等级 默认不配置是debug debug->info->warn->error
redis:
  -
    name: default
    host: 127.0.0.1
    port: 6379
    password:
    db:
  -
    name: redis2
    host: 127.0.0.1
    port: 6379
    password:
    db:
mysql:
  -
    key: gomicro #默认缺省 default
    driver: mysql
    host: 172.18.15.43
    port: 3306
    database: gomicro
    username: developer
    password: DN3v74JB
    charset: utf8
    prefix:
    policies: 1 #从库的负载策略 默认轮询访问 0:轮询访问; 1：随机访问;2：权重随机;3：权重轮询;4：最小连接数
    policies_weight:
    max_idle_conns:
    max_open_conns:
    conn_max_lifetime:
    master_slave: #主从关键字
      -
        key: master #主从库关键字 主库/从库
        driver: mysql
        host: 172.18.15.43
        port: 3306
        database: gomicro
        username: developer
        password: DN3v74JB
        charset: utf8
        prefix: #表名前缀
        policies_weight: 2 # 权重比
app_center:
  app_host: http://123.59.58.212:8001
  app_url: /app/getTencentCosConfig
`
