# xie-blog
![Static Badge](https://img.shields.io/badge/go-1.21.4-skyblue)
![Static Badge](https://img.shields.io/badge/gorm-1.25.10-green)
![Static Badge](https://img.shields.io/badge/gin-1.10.0-blue)
![Static Badge](https://img.shields.io/badge/jwt-5.2.1-yellow)

## 简介📖

这是一个基于 go 的 gin 框架构建的开源项目，旨在为开发者提供一个高效、稳定、美观的管理系统解决方案。项目采用现代化的前后端技术栈，注重代码的可维护性和开发效率。 

**前端项目地址**: [xie-blog-web](https://github.com/Jaynxe/xie-blog-web)
## 技术栈💻
 Go, Gin, GORM, JWT, Redis, CORS, Swagger, Stringer, Validator, Snowflake 算法, Logrus, 七牛云存储, Gomail

## 特性⭐

- 后端使用 Go 和 Gin 开发高效、稳定的服务。
- 使用 GORM 进行数据库操作，确保数据的完整性和一致性。
- 使用 JWT 进行用户认证和授权，并使用 Redis 存储随机生成的密钥。
- 支持 CORS 接口跨域、Swagger 接口文档生成、统一定义错误码，使用 Stringer 生成错误码字符串。
- 使用 Validator 实现数据绑定和验证。
- 设计并实现基于 Snowflake 算法的用户 ID 生成与加密，实现用户密码哈希处理与登录密码验证，确保密码安全。
- 使用 Logrus 进行日志记录，并实现按时间和等级分割日志文件的功能。
- 实现了七牛云存储和本地文件存储功能。
- 使用 Gomail 发送邮箱验证码，通过 Session 存储和验证验证码，实现验证码登录功能。

## 安装与运行

1. 克隆项目：
    ```sh
    git clone https://github.com/Jaynxe/xie-blog.git
    ```

2. 安装依赖：
    ```sh
    go mod tidy
    ```

3. 运行服务：
    ```sh
    go run main.go
    ```

## 联系方式

- **作者**: 谢安
- **邮箱**: xjy158191@126.com
