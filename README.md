#webBlog
[小文博客地址](https://www.az1314.cn/)

## 初衷
之前博客项目[小文博客](https://www.az1314.cn/)基于自己通过swoole改造laravel框架运行lnmp+swoole环境中；
由于这个博客写的比较早，所以使用也很多不理想；
so,拿GO来重构下
这次后台前后端不分离，无权限控制，默认markdown编辑；ps:个人博客，无需分离，主要是懒~
前台前后端分离，所以前端Coder不精通后端的完全可以拿去用~当然我也会写一个默认前台，和我之前博客前台模板一样
## 技术选型
1. web:[gin](https://github.com/gin-gonic/gin)
2. orm:[gorm](https://github.com/jinzhu/gorm)
3. database:[mysql](https://github.com/go-sql-driver/mysql)
4. 文件存储:[七牛云存储](https://www.qiniu.com/)
5. Pjax
## 项目结构
```
-webBlog
    |-config 配置文件目录
    |-controller 控制器目录
        |-admin 后台控制器
        |-home 默认前台模板
        |-api 前台接口
    |-helper 公共函数目录
    |-model model层
    |-middlermare 中间件
    |-static 静态资源目录
    |-view 模板文件目录
        |-admin 后台模板
    |-main.go 程序执行入口
    |-go.mod go官方依赖管理
```
## 进度
- [√] 登录管理
- [√] 分类管理
- [√] 文章管理
- [√] 轮播图管理
- [√] 网站配置管理
- [√] 导航管理
- [√] 友链管理
- [√] markdown接入
- [√] 七牛云存储
- [ ] 评论功能
- [ ] 前台接口
- [ ] 前台模板渲染
- ...


## 安装部署
本项目使用go mod管理依赖包，所以你的go版本必须是1.11及以上

```
git clone https://github.com/xiaowen1108/webBlog.git
cd webBlog
go run main.go [-app_conf_file="配置文件地址"]
```

## 使用方法
### 使用说明
1. 修改app_bck.ini，设置mysql配置及其他可选配置，并重命名为app.ini
2. 七牛云存储，请在app.ini设置qiniu相关配置项，不设置也行，就是文章没封面图~~~
3. 系统运行会自动基于配置account生成后台管理账户
4. 打开site:8080/admin/login登录

## 效果图

![file](https://img.az1314.cn/20190111184054692.png)

![file](https://img.az1314.cn/20190111184150946.png)

![file](https://img.az1314.cn/20190111184245143.jpeg)
