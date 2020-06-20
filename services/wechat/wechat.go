package wechat

import (
	"fmt"
	"myapp/common"
	"myapp/config"
)

// 获取并缓存access_token
// 获取用户信息接口
// 推送企业微信，创建企业微信
// 使用设计模式，开发企业微信，封装restful风格，尝试可以添加应用

// 定义变量缓存时间和token
var (
	GLOBALACCESSTOKEN string
	GLOEXPIRESIN      string
)

//GetQyInfo 获取企业微信参数
func GetQyInfo(wechatInfo config.Uxiaowechat) {
	fmt.Println("wechatInfo", wechatInfo)
}

//GenerageAccessToken 获取token
func GenerageAccessToken(URL, corpid, corpsecret string) (err error) {
	reqURL := URL + "?corpid=" + corpid + "&corpsecret=" + corpsecret
	res, err := common.PublicGet(reqURL, "", "")
	if err != nil {
		return err
	}
	fmt.Println("res", res)
	return err
}

//SetExpiresIn 设置有效期
func SetExpiresIn() {

}

//GetAccessToken 获取token
func GetAccessToken() {
	//获取最新的token
	if GLOBALACCESSTOKEN == "" {
		fmt.Println("没有缓存token")
	}
}
