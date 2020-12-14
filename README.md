# Gin-Element-Admin

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

Gin-Element-Admin 快速开发平台Fork自 gin-vue-admin 的master vue2版本，后端使用Golang的gin，前端使用Vue的ElementUI，删除了原项目的部分功能，对前端内容进行了全面重写，前端配色风格接近于Quasar Framework

*目前项目仍在不断开发中，如果你感觉不错，欢迎给个小小的Star！*

> 鸣谢 https://github.com/flipped-aurora/gin-vue-admin

#
ElementUI版本：https://github.com/Junvary/gea

Quasar版本：https://github.com/Junvary/gin-quasar-admin

> Quasar版本旨在学习在Vue中使用Quasar Framework，听说ElementUI原班人马已经回归，更要支持了！

> 目前项目的初始化方法与 gin-vue-admin 保持一致，但数据库初始化文件内容进行了相应的更改，推荐重新建库体验。初始化前请修改 gea-server 目录下 config.yaml 中 mysql 配置为自己的数据库信息。

> 由于 Gin-Element-Admin 移除了一些 gin-vue-admin 的功能，如需添加请自行参考原项目。

### 项目正在不断完善中，后续此文档将持续更新与原项目的区别，方便参考。


| 内容 |  gin-vue-admin   | Gin-Element-Admin  |
|  :----:  |  :----:  | :----:  |
| elementUI  | gva风格  | Quasar风格 |
| vuex  | 通过插件保存本地  | 登录只获取token保存cookie，其他状态重新请求 |




# 项目截图

![gea-login.png](https://i.loli.net/2020/12/14/SgkeMj1cis4VhJG.png)

![gea-dashboard.png](https://i.loli.net/2020/12/14/UVsQzXArxLNaRjF.png)

![gea-userinfo.png](https://i.loli.net/2020/12/14/5UcPGla1jzdgE8M.png)

![gea-server.png](https://i.loli.net/2020/12/14/yz78vanKPbuxHX2.png)












