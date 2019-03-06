### 基于GRPC的鉴权案例

#####目录结构

- autuentication  鉴权

- gateway 网关

- grpc_server 相关服务

- proto 

- unit 中间件

代码中写死了用户的权限相关内容，实际应用中可以从数据库或文件中获取

### 测试

```gotemplate
    获取Token 及用户角色
    http://localhost:3344/login?username=felix
    
    查看是否有权限
    
    http://localhost:3344/api/menu/2 （通过）
    http://localhost:3344/api/menus/1 （异常）
```