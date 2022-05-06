# bbsproject
a bbs for Graduation Project

### :warning:本项目是一个究极缝合，[源项目github地址](https://github.com/jiawood/simple-bbs)
由于毕设选题的时候牛皮吹大了，eng着头皮在github上找了个项目魔改:clown_face:。
#### 主要制作的内容：使用go语言作为后端语言，提供API接口满足前端的行为（注册，登录，状态验证，发帖，评论）
项目本身属于半吊子魔改，所以实际使用起来还是会存在一些问题，但是正常操作不会有啥事儿:tea:

## 前端项目 :point_right: 1:`npm install`  2:`npm run serve`
## 后端项目我使用wsl 作为服务器，第一次运行时还需要
在comtroller中找到db.create语句，在此前插入如下代码
``` 
 if db.HasTable(&tablename) == false {
      db.CreateTable(&tablename)
}
```
## 总而言之，魔改的项目，配置起来有点复杂，本次github也是初次使用，希望是一个好的开端吧🤡
