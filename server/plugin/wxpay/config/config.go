package config

type Wxpay struct {
	MchID                      string // 商户ID
	AppID                      string // 绑定小程序的APPID
	NotifyUrl                  string // 支付回调域名
	MchCertificateSerialNumber string // 商户证书序列号
	MchAPIv3Key                string // 商户APIv3密钥
	MchAPIv2Key                string // 商户APIv2密钥
	CertPath                   string // Cert证书文件所在地址
	KeyPath                    string // Key证书文件所在地址
}
