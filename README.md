# 一、安装
```bash
go get -u github.com/whrwsoftware/pay
```

# 二、文档目录

> ### 点击查看不同支付方式的使用文档。

* #### [Alipay](https://github.com/whrwsoftware/pay/blob/main/doc/alipay.md)
* #### [Wechat](https://github.com/whrwsoftware/pay/blob/main/doc/wechat_v3.md)
* #### [QQ](https://github.com/whrwsoftware/pay/blob/main/doc/qq.md)
* #### [Paypal](https://github.com/whrwsoftware/pay/blob/main/doc/paypal.md)
* #### [Apple](https://github.com/whrwsoftware/pay/blob/main/doc/apple.md)

---

<br>

# 三、其他说明

* 各支付方式接入，请仔细查看 `xxx_test.go` 使用方式
    * `pay/wechat/v3/client_test.go`
    * `pay/alipay/client_test.go`
    * `pay/qq/client_test.go`
    * `pay/paypal/client_test.go`
    * `pay/apple/verify_test.go`
    * 或 examples

* 开发过程中，请尽量使用正式环境，1分钱测试法！
