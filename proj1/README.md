目标：用Golang实现用户注册登录功能

* 打开localhost:3000/signup，填写注册信息
    * 用户名(验证要求：手机号码或邮箱)
    * 密码(拥有大小写字母及数字，至少8位)
* 打开localhost:3000/signin，显示登录框(可填写用户名和密码)，登录框下方有按钮可以调到signup页面注册
    * 若登录成功 -> 跳转localhost:3000/profile，显示用户名及Login Succeeded
    * 若登录失败 -> 停留当前页面，显示Login Failed
* 登录状态保持在前端，未登录状态打开localhost:3000/profile，跳转至localhost:3000/signin

模块需求
* 任选一种web framework https://github.com/mingrammer/go-web-framework-stars
* 数据库使用postgres
* 登录状态缓存使用redis

To Do：
* 增加postgres的使用
* 使用redis

运行
<br>go build -o app
<br>./app

<br>
Code adapted from https://github.com/demo-apps/go-gin-app
