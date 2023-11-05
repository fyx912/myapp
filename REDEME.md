### 一、目录结构
-  main.go ：包含应用程序的主入口点。 
-  handlers/ ：包含处理HTTP请求的处理程序或控制器。 
-  models/ ：包含应用程序中使用的数据模型或结构。 
-  routes/ ：包含路由定义或URL映射。 
-  middlewares/ ：包含应用程序中使用的中间件函数或处理程序。 
-  utils/ ：包含实用函数或辅助包。 
-  tests/ ：包含应用程序的测试文件。 
-  configs/ ：包含应用程序的配置文件。 
-  docs/ ：包含API文档或其他相关文档。 
-  scripts/ ：包含构建脚本或其他实用脚本。 
-  vendor/ ：包含应用程序的依赖包。 
-  README.md ：包含有关应用程序的信息和文档

- **cmd/**：该目录包含主应用程序文件，通常以应用程序名称命名。 
- **pkg/**：该目录包含打算被外部包导入的包或模块。 
- **internal/**：该目录包含不打算被外部包导入的内部包或模块。这些内部包对于项目的其他部分是私有的。 
  - **internal/handlers/**: 包含处理不同API端点的HTTP请求的处理程序或控制器。 
  - **internal/models/**: 包含应用程序中使用的数据模型或结构。 
  - **internal/middlewares/**: 包含用于API的自定义中间件函数。 
  - **internal/services/**: 包含API使用的业务逻辑或服务。 
  - **internal/repositories/**: 包含与数据库交互的数据访问层或仓库。 
  - **internal/utils/**: 包含在整个应用程序中使用的实用函数或辅助方法。
- **configs/**：该目录包含应用程序的配置文件，例如特定环境的设置或数据库配置。 
- **scripts/**：该目录包含与项目相关的任何脚本或自动化工具。 
- **test/**：该目录包含项目的测试文件。 
- **vendor/**：该目录存放项目的依赖项，通常使用Go Modules或Dep等依赖管理工具进行管理

### 二、依赖
```
go get -u github.com/gin-gonic/gin
go get -u github.com/jinzhu/gorm
go get -u github.com/go-sql-driver/mysql
go get -u gopkg.in/yaml.v2
go get github.com/google/uuid
```
### 三、运行命令
```
go mod init    初始化
go mog tidy    拉取缺少得模块，移除不用得模块
go mod vendor  将依赖复制到 vendor下
```
go.mod 提供了module, require、replace和exclude 四个命令
- module 语句指定包的名字（路径）
- require 语句指定的依赖项模块
- replace 语句可以替换依赖项模块
- exclude 语句可以忽略依赖项模块
#### 3.1、运行项目
```
go run main.go
```