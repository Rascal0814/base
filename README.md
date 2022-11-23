# 基于kratos框架开发
[kratos官方文档](https://go-kratos.dev/docs/)

### 目录说明
```text
base
├─third_party                           //第三方包
|      ├─README.md
|      ├─validate
|      ├─openapi
|      ├─google
|      ├─errors                     
├─pkg                                   //工具包
├─app                                   //微服务
|  ├─user                               //user服务
|  |  ├─internal                        //内部逻辑层
|  |  |    ├─service                    //处理层
|  |  |    ├─server                     //http_grpc连接实例
|  |  |    ├─data                       //数据层
|  |  |    ├─conf                       //conf_proto文件
|  |  |    ├─biz                        //业务层
|  |  ├─configs                         //服务配置文件
|  |  ├─cmd
|  |  |  ├─user
|  |  |  |  ├─main.go                   //启动文件
|  |  |  |  ├─wire.go                   //wire依赖注入
├─api                                   //proto文件
```
### 功能说明
- user  
``用户服务``
- login  
```登入```

### 工具使用
- proto生成文件  
    ```
    make api
    ```
- wire 依赖注入  
    ```text
    一般来说 service 依赖 biz 依赖 data  
    biz层会定义一些接口去让data层去实现方便后续替换数据层   
    依赖工具到wire文件下使用 wire 命令生成实例
    ```
- [kratos-cli](https://go-kratos.dev/docs/getting-started/usage)
    ```text
  1. 创建新的微服务
  kratos new app/user --nomod
  2. 生成service代码
  kratos proto server xxx.proto -t xxx(目录)
  3. 启动项目
  kratos run 
  4. debug项目
  需要制定configs文件夹 -conf xxx
  ```
### 服务发现
consul  
ui地址：http://localhost:8500/ui/