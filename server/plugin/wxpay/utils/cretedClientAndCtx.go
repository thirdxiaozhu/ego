package utils

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxpay/global"
)

func CreateClientAndCtx() (error, context.Context, *payment.Payment) {
	PaymentService, err := payment.NewPayment(&payment.UserConfig{
		AppID:       global.GlobalConfig.AppID,                      // 小程序、公众号或者企业微信的appid
		MchID:       global.GlobalConfig.MchID,                      // 商户号 appID
		MchApiV3Key: global.GlobalConfig.MchAPIv3Key,                // 微信V3接口调用必填
		Key:         global.GlobalConfig.MchAPIv2Key,                // 微信V2接口调用必填
		CertPath:    global.GlobalConfig.CertPath,                   // 商户后台支付的Cert证书路径
		KeyPath:     global.GlobalConfig.KeyPath,                    // 商户后台支付的Key证书路径
		SerialNo:    global.GlobalConfig.MchCertificateSerialNumber, // 商户支付证书序列号
		NotifyURL:   global.GlobalConfig.NotifyUrl,
		HttpDebug:   false, // 订单模式不支持debug 请勿打开
		Log: payment.Log{
			Level: "debug",
			// 可以重定向到你的目录下，如果设置File和Error，默认会在当前目录下的wechat文件夹下生成日志
			File:   "info.log",
			Error:  "error.log",
			Stdout: false, //  是否打印在终端
		},
		Http: payment.Http{
			Timeout: 30.0,
			BaseURI: "https://api.mch.weixin.qq.com",
		},
	})
	return err, context.Background(), PaymentService
}
