<div align=center>
<img src="https://i.loli.net/2020/12/14/cnJoF9r1BXY7Da5.png" width=200" height="200" />
<h1>Gin-Element-Admin</h1>
</div>
<div align=center>
<img src="https://img.shields.io/badge/golang-1.15-blue"/>
<img src="https://img.shields.io/badge/gin-1.4.0-lightBlue"/>
<img src="https://img.shields.io/badge/vue-2.6.10-brightgreen"/>
<img src="https://img.shields.io/badge/element--ui-2.14.1-green"/>
<img src="https://img.shields.io/badge/gorm-1.9.12-red"/>
</div>

<center>
Gin-Element-Admin 快速开发平台Fork自 gin-vue-admin 的vue2版本，并进行了大量修改。

鸣谢 https://github.com/flipped-aurora/gin-vue-admin
</center>

Gin-Element-Admin 后端使用Golang的gin，前端使用Vue的ElementUI，配色接近Quasar。目前项目仍在不断开发中，欢迎clone和fork试用，如果你感觉不错，请给个小小的 Star 鼓励一下！

#
Gin-Element-Admin 具有ElementUI和Quasar版本：

ElementUI版本：https://github.com/Junvary/gea

Quasar版本：https://github.com/Junvary/gin-quasar-admin

> Quasar版本旨在学习使用Quasar Framework，希望给初次使用 Quasar的小伙伴一些帮助。

# 更新内容：

> 目前项目的初始化方法与 gin-vue-admin 保持一致，但数据库初始化文件内容进行了相应的更改，推荐重新建库体验。初始化前请修改 gea-server 目录下 config.yaml 中 mysql 配置为自己的数据库信息。

> 由于 Gin-Element-Admin 移除了一些 gin-vue-admin 的功能，如需添加请自行参考原项目。

### 项目正在不断完善中，罗列一些近期更新的内容，方便参考。


| 内容 | Gin-Element-Admin  |
|  :----: | :----:  |
| 移除 | 各类示例模块，远程存储，email功能等 |
| web | 重写了前端部分为 gea-web |
| elementUI | 使用elementUI并开发为Quasar风格 |
| 菜单 | 重新划分菜单分组 |
| vuex | 登录只获取token，并保存在cookie中，其他状态每次通过权限捕获重新请求后端 |
| 代码生成器 | 优化了代码生成器，并包含表单验证 |
| devops | 增加 devops 模块 |




# 项目截图

![gea-login.png](https://i.loli.net/2020/12/14/SgkeMj1cis4VhJG.png)

![gea-dashboard.png](https://i.loli.net/2020/12/14/UVsQzXArxLNaRjF.png)

![gea-userinfo.png](https://i.loli.net/2020/12/14/5UcPGla1jzdgE8M.png)

![gea-server.png](https://i.loli.net/2020/12/14/yz78vanKPbuxHX2.png)












