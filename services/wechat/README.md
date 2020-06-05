# 企业微信开发整理

服务端API `https://work.weixin.qq.com/api/doc/90000/90135/90664`

详细代码见 `wechat.go`

> 企业微信开发分为服务端API和客户端API

一、企业微信开发前需要了解几个问题

默认你已经知道了下面的问题

1、企业微信开发，一般对应的是创建相关应用（这里创建应用取名：优效）

2、![微信的后台获取的参数](https://work.weixin.qq.com/api/doc/90000/90135/90665)

```
corpid
userid
部门id
tagid
agentid
secret
access_token
```
3、这里的开发主要是三个地方的处理

- 获取access_token
- 身份验证-网页授权登录
- 消息推送

二、获取access_token

调用企业微信接口首先需要获取access_token，获取access_token的方法很简单

1、首先在企业微信后台获取企业微信的企业微信ID(corpid)和应用的密钥(Secret)

2、调用企业微信接口

```
https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=ID&corpsecret=SECRET

返回参数：
{
   "errcode": 0,
   "errmsg": "ok",
   "access_token": "accesstoken000001",
   "expires_in": 7200
}

```
3、缓存access_token

定义两个全局变量，一个用来存放access_token,一个用来存放access_token失效时间

```
GLOBALTOKEN 存放access_token
GLOBALEXPIRE 存放access_token失效时间
```

- 获取access_token
- 查询GLOBALTOKEN是否有值，无值直接调用接口获取，若有值
- 通过GLOBALEXPIRE是否失效，有效则直接返回，无效则调用接口获取

三、网页授权登录

很多时候，我们需要在手机端查看信息，开发自己的页面展示，如果每次都登录太麻烦，可以直接调用企业微信接口获取用户信息，在再此之前需要通过认证

参数  |	必须  |	说明
---|---|---
appid |	是	| 企业的CorpID
redirect_uri  |	是	| 授权后重定向的回调链接地址，请使用urlencode对链接进行处理
response_type |	是	| 返回类型，此时固定为：code
scope |	是 | 应用授权作用域。企业自建应用固定填写：snsapi_base
state |	否 | 重定向后会带上state参数，企业可以填写a-zA-Z0-9的参数值，长度不可超过128个字节
#wechat_redirect  |	是 |  终端使用此参数判断是否需要带上身份信息

```
假定当前
企业CorpID：wxCorpId
访问链接：http://api.3dept.com/cgi-bin/query?action=get
根据URL规范，将上述参数分别进行UrlEncode，得到拼接的OAuth2链接为：
https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxCorpId&redirect_uri=http%3a%2f%2fapi.3dept.com%2fcgi-bin%2fquery%3faction%3dget&response_type=code&scope=snsapi_base&state=#wechat_redirect
员工点击后，页面将跳转至 
http://api.3dept.com/cgi-bin/query?action=get&code=AAAAAAgG333qs9EdaPbCAP1VaOrjuNkiAZHTWgaWsZQ&state=
企业可根据code参数调用获得员工的userid

```
> 注意到，构造OAuth2链接中参数的redirect_uri是经过UrlEncode的

通过构造OAuth2链接，可以进行微信认证，拿到code就可以获取用户信息

```
// access_token 第二步获取的access_token
// code 第三步获取的code,code有效期为5分钟，只能使用一次
https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=ACCESS_TOKEN&code=CODE

{
   "errcode": 0,
   "errmsg": "ok",
   "UserId":"USERID",
   "DeviceId":"DEVICEID"
}

```
获取userid后，在调用读取成员信息接口

```
https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&userid=USERID

// 返回的信息（展示部分）
{
    "errcode": 0,
    "errmsg": "ok",
    "userid": "zhangsan",
    "name": "李四",
    "department": [1, 2],
    "order": [1, 2],
}

```

注意：上面如果每次打开企业微信，重定向获取code会有时间的消耗，而且转发、刷新后会导致页面获取用户信息失败，最好进行下面的处理

1、前端拿到code后，返回给后端，后端返回用户信息、cookies信息和有效期（服务端设置），前端缓存在本地，作为用户登录的凭证
2、前端直接使用自己页面的链接，如果没有cookies信息，在前端重定向到认证页面，获取code,调用接口，缓存信息
3、获取用户信息，登录页面

这样处理就比较灵活，可以跟据页面设置有效期，页面的打开速度也会快了很多。


四、推送信息

企业微信支持的消息类型有

- 文本消息
- 图片消息
- 语音消息
- 视频消息
- 文件消息
- 文本卡片消息
- 图文消息
- 图文消息（mpnews）
- markdown消息
- 小程序通知消息
- 任务卡片消息

1、推送信息给个人

```
https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=ACCESS_TOKEN

对应的请求参数，不通类型不通可以通过企业微信查看(https://work.weixin.qq.com/api/doc/90000/90135/90236)

// 返回参数
 {
   "errcode" : 0,
   "errmsg" : "ok",
   "invaliduser" : "userid1|userid2",
   "invalidparty" : "partyid1|partyid2",
   "invalidtag": "tagid1|tagid2"
 }

```
如果所有的用户不存在，才会返回失败，部分不存在会提示在invaliduser字段中

2、推送信息给群聊

企业微信必须通过api创建群聊，然后才能进行推送

创建群聊：
```
https://qyapi.weixin.qq.com/cgi-bin/appchat/create?access_token=ACCESS_TOKEN

//请求参数
{
    "name" : "NAME",
    "owner" : "userid1",
    "userlist" : ["userid1", "userid2", "userid3"],
    "chatid" : "CHATID"
}

//返回参数
{
    "errcode" : 0,
    "errmsg" : "ok",
    "chatid" : "CHATID"
}

```

参数 |	是否必须	| 说明
---|---|---
access_token|	是|	调用接口凭证
name|	否	|群聊名，最多50个utf8字符，超过将截断
owner|	否|	指定群主的id。如果不指定，系统会随机从userlist中选一人作为群主
userlist|	是|	群成员id列表。至少2人，至多2000人
chatid|	否|	群聊的唯一标志，不能与已有的群重复；字符串类型，最长32个字符。只允许字符0-9及字母a-zA-Z。如果不填，系统会随机生成群id


推送群聊信息
```
https://qyapi.weixin.qq.com/cgi-bin/appchat/send?access_token=ACCESS_TOKEN

//对应的请求参数，不通类型不通可以通过企业微信查看(https://work.weixin.qq.com/api/doc/90000/90135/90248)

{
    "chatid": "CHATID",
    "msgtype":"text",
    "text":{
        "content" : "你的快递已到\n请携带工卡前往邮件中心领取"
    },
    "safe":0
}

```










