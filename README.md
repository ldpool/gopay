<div align=center><img width="240" height="240" alt="Logo was Loading Faild!" src="logo.png"/></div>

# GoPay

### 微信、支付宝、QQ、通联支付、拉卡拉、PayPal 的 Golang 版本 SDK

---

# 一、安装

```bash
go get -u github.com/ldpool/gopay
```

#### 查看 GoPay 版本

[版本更新记录](https://github.com/ldpool/gopay/blob/main/release_note.txt)

```go
import (
    "github.com/ldpool/gopay"
    "github.com/go-pay/xlog"
)

func main() {
    xlog.Info("GoPay Version: ", gopay.Version)
}
```

---

<br>

# 二、文档目录

> ### 点击查看不同支付方式的使用文档。方便的话，请留下您认可的小星星，十分感谢！

- #### [支付宝支付](https://github.com/ldpool/gopay/blob/main/doc/alipay.md)
- #### [微信支付](https://github.com/ldpool/gopay/blob/main/doc/wechat_v3.md)
- #### [QQ 支付](https://github.com/ldpool/gopay/blob/main/doc/qq.md)
- #### [通联支付](https://github.com/ldpool/gopay/blob/main/doc/allinpay.md)
- #### [拉卡拉支付](https://github.com/ldpool/gopay/blob/main/doc/lakala.md)
- #### [Paypal 支付](https://github.com/ldpool/gopay/blob/main/doc/paypal.md)
- #### [Apple 支付校验](https://github.com/ldpool/gopay/blob/main/doc/apple.md)
- #### [扫呗支付](https://github.com/ldpool/gopay/blob/main/doc/saobei.md)

---

<br>

# 三、其他说明

- 如需自定义 Log 输出，请调用以下方法设置自定义 Logger，实现 `xlog.XLogger` 接口即可。
  - `xlog.SetDebugLog()`
  - `xlog.SetInfoLog()`
  - `xlog.SetWarnLog()`
  - `xlog.SetErrLog()`
- 各支付方式接入，请仔细查看 `xxx_test.go` 使用方式
  - `gopay/wechat/v3/client_test.go`
  - `gopay/alipay/client_test.go`
  - `gopay/qq/client_test.go`
  - `gopay/allinpay/client_test.go`
  - `gopay/lakala/client_test.go`
  - `gopay/paypal/client_test.go`
  - `gopay/apple/verify_test.go`
  - 或 examples
- 接入 gopay 示例项目(可参考接入使用方式)：[gopay-platform](https://github.com/ldpool/gopay-platform)
