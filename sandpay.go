package sandpay

import (
	"github.com/shop-r1/sandpay/pay"
	"github.com/shop-r1/sandpay/pay/params"
	"github.com/shop-r1/sandpay/pay/request"
	"github.com/shop-r1/sandpay/pay/response"
	"net/url"
	"strings"
	"time"
)

var Client SandPay

func init() {
	//TODO 改为配置文件
	Client.Config = pay.Config{
		Version:         "1.0",
		MerId:           "S6887745",
		PrivatePath:     "/Users/linwenxiang/Downloads/certs/server.key", //私钥文件
		CertPath:        "/Users/linwenxiang/Downloads/certs/server.crt", //公钥文件
		EncryptCertPath: "/Users/linwenxiang/Downloads/certs/sand.cer",   //导出的公钥文件
		ApiHost:         "https://cashier.sandpay.com.cn",
		NotifyUrl:       "https://www.baidu.com",
		FrontUrl:        "https://www.baidu.com",
		//ApiHost:   	 "http://61.129.71.103:8003",
		//ApiHost:       "https://cashier.sandpay.com.cn/gateway/api",
	}
	pay.LoadCertInfo(&Client.Config)
}

type SandPay struct {
	Config pay.Config
}

// 微信统一下单接口
func (sandPay *SandPay) OrderPayWx(params params.OrderPayParams) (resp response.OrderPayResponse, err error) {
	config := sandPay.Config
	timeString := time.Now().Format("20060102150405")

	header := request.Header{}
	header.SetMethod(`sandpay.trade.pay`).SetVersion(`1.0`).SetAccessType("1")
	header.SetChannelType("07").SetMid(config.MerId).SetProductId("00002020").SetReqTime(timeString)
	body := request.OrderPayBody{
		//PayTool: "0402",
		OrderCode:   params.OrderNo,
		TotalAmount: params.GetTotalAmountToString(),
		Subject:     params.Subject,
		Body:        params.Body,
		TxnTimeOut:  params.TxnTimeOut,
		PayMode:     params.PayMode,
		PayExtra:    params.PayExtra.ToJson(),
		ClientIp:    params.ClientIp,
		NotifyUrl:   sandPay.Config.NotifyUrl,
		FrontUrl:    sandPay.Config.FrontUrl,
		Extends:     params.Extends,
	}

	signDataJsonString := pay.GenerateSignString(body, header)
	sign, _ := pay.PrivateSha1SignData(signDataJsonString)
	postData := pay.GeneratePostData(signDataJsonString, sign)

	data, err := pay.PayPost(config.ApiHost+"/qr/api/order/create", postData)
	if err != nil {
		return
	}
	resp.SetData(data.Data)
	return resp, err
}

// pc微信统一下单接口
func (sandPay *SandPay) OrderPayQrWx(params params.OrderPayParams) (resp response.OrderPayResponse, err error) {
	config := sandPay.Config
	timeString := time.Now().Format("20060102150405")

	header := request.Header{}
	header.SetMethod(`sandpay.trade.precreate`).SetVersion(`1.0`).SetAccessType("1")
	header.SetChannelType("07").SetMid(config.MerId).SetProductId("00000005").SetReqTime(timeString)
	body := request.OrderPayBody{
		PayTool: "0402",
		OrderCode:   params.OrderNo,
		TotalAmount: params.GetTotalAmountToString(),
		Subject:     params.Subject,
		Body:        params.Body,
		TxnTimeOut:  params.TxnTimeOut,
		PayMode:     params.PayMode,
		PayExtra:    params.PayExtra.ToJson(),
		ClientIp:    params.ClientIp,
		NotifyUrl:   sandPay.Config.NotifyUrl,
		FrontUrl:    sandPay.Config.FrontUrl,
		Extends:     params.Extends,
	}

	signDataJsonString := pay.GenerateSignString(body, header)
	sign, _ := pay.PrivateSha1SignData(signDataJsonString)
	postData := pay.GeneratePostData(signDataJsonString, sign)

	data, err := pay.PayPost(config.ApiHost+"/qr/api/order/create", postData)
	if err != nil {
		return
	}
	resp.SetData(data.Data)
	return resp, err
}

// 支付宝统一下单接口
func (sandPay *SandPay) OrderPayQrAlipay(params params.OrderPayParams) (resp response.OrderPayResponse, err error) {
	config := sandPay.Config
	timeString := time.Now().Format("20060102150405")

	header := request.Header{}
	header.SetMethod(`sandpay.trade.precreate`).SetVersion(`1.0`).SetAccessType("1")
	header.SetChannelType("07").SetMid(config.MerId).SetProductId("00000006").SetReqTime(timeString)
	body := request.OrderPayBody{
		PayTool: "0401",
		OrderCode:   params.OrderNo,
		TotalAmount: params.GetTotalAmountToString(),
		Subject:     params.Subject,
		Body:        params.Body,
		TxnTimeOut:  params.TxnTimeOut,
		PayMode:     params.PayMode,
		PayExtra:    params.PayExtra.ToJson(),
		ClientIp:    params.ClientIp,
		NotifyUrl:   sandPay.Config.NotifyUrl,
		FrontUrl:    sandPay.Config.FrontUrl,
		Extends:     params.Extends,
	}

	signDataJsonString := pay.GenerateSignString(body, header)
	sign, _ := pay.PrivateSha1SignData(signDataJsonString)
	postData := pay.GeneratePostData(signDataJsonString, sign)

	data, err := pay.PayPost(config.ApiHost+"/qr/api/order/create", postData)
	if err != nil {
		return
	}
	resp.SetData(data.Data)
	return resp, err
}

// 统一下单接口
func (sandPay *SandPay) OrderPay(params params.OrderPayParams) (resp response.OrderPayResponse, err error) {
	config := sandPay.Config
	timeString := time.Now().Format("20060102150405")

	header := request.Header{}
	header.SetMethod(`sandpay.trade.pay`).SetVersion(`1.0`).SetAccessType("1")
	header.SetChannelType("07").SetMid(config.MerId).SetProductId("00000005").SetReqTime(timeString)
	body := request.OrderPayBody{
		OrderCode:   params.OrderNo,
		TotalAmount: params.GetTotalAmountToString(),
		Subject:     params.Subject,
		Body:        params.Body,
		TxnTimeOut:  params.TxnTimeOut,
		PayMode:     params.PayMode,
		PayExtra:    params.PayExtra.ToJson(),
		ClientIp:    params.ClientIp,
		NotifyUrl:   sandPay.Config.NotifyUrl,
		FrontUrl:    sandPay.Config.FrontUrl,
		Extends:     params.Extends,
	}

	signDataJsonString := pay.GenerateSignString(body, header)
	sign, _ := pay.PrivateSha1SignData(signDataJsonString)
	postData := pay.GeneratePostData(signDataJsonString, sign)

	data, err := pay.PayPost(config.ApiHost+"/gateway/api/order/pay", postData)
	if err != nil {
		return
	}
	resp.SetData(data.Data)
	return resp, err
}

// 订单查询接口
func (sandPay *SandPay) OrderQuery(orderNo string, extend string) (resp response.OrderQueryResponse, err error) {
	config := sandPay.Config
	timeString := time.Now().Format("20060102150405")

	header := request.Header{}
	header.SetMethod(`sandpay.trade.query`).SetVersion(`1.0`).SetAccessType("1")
	header.SetChannelType("07").SetMid(config.MerId).SetProductId("00000007").SetReqTime(timeString)
	body := request.OrderQueryBody{
		OrderCode: orderNo,
		Extends:   extend,
	}

	signDataJsonString := pay.GenerateSignString(body, header)
	sign, _ := pay.PrivateSha1SignData(signDataJsonString)
	postData := pay.GeneratePostData(signDataJsonString, sign)

	data, err := pay.PayPost(config.ApiHost+"/order/query", postData)

	if err != nil {
		return
	}
	resp.SetData(data.Data)
	return resp, err
}

// 退货申请接口
func (sandPay *SandPay) OrderRefund(refundParams params.OrderRefundParams) (resp response.OrderRefundResponse, err error) {
	config := sandPay.Config
	timeString := time.Now().Format("20060102150405")

	header := request.Header{}
	header.SetMethod(`sandpay.trade.refund`).SetVersion(`1.0`).SetAccessType("1")
	header.SetChannelType("07").SetMid(config.MerId).SetProductId("00000007").SetReqTime(timeString)
	body := request.OrderRefundBody{
		OrderCode:    refundParams.OrderNo,
		OriOrderCode: refundParams.RefundNO,
		RefundAmount: refundParams.GetRefundAmount(),
		NotifyUrl:    config.NotifyUrl,
		RefundReason: refundParams.RefundReason,
		Extends:      refundParams.Extends,
	}

	signDataJsonString := pay.GenerateSignString(body, header)
	sign, _ := pay.PrivateSha1SignData(signDataJsonString)
	postData := pay.GeneratePostData(signDataJsonString, sign)

	data, err := pay.PayPost(config.ApiHost+"/order/refund", postData)
	if err != nil {
		return
	}
	resp.SetData(data.Data)
	return resp, err
}

// 退货申请接口
func (sandPay *SandPay) OrderRefunds(refundParams params.OrderRefundParams) (resp response.OrderRefundResponse, err error) {
	config := sandPay.Config
	timeString := time.Now().Format("20060102150405")

	header := request.Header{}
	header.SetMethod(`sandpay.trade.refund`).SetVersion(`1.0`).SetAccessType("1")
	header.SetChannelType("07").SetMid(config.MerId).SetProductId("00000007").SetReqTime(timeString)
	body := request.OrderRefundBody{
		OrderCode:    refundParams.OrderNo,
		OriOrderCode: refundParams.RefundNO,
		RefundAmount: refundParams.GetRefundAmount(),
		NotifyUrl:    config.NotifyUrl,
		RefundReason: refundParams.RefundReason,
		Extends:      refundParams.Extends,
	}

	signDataJsonString := pay.GenerateSignString(body, header)
	sign, _ := pay.PrivateSha1SignData(signDataJsonString)
	postData := pay.GeneratePostData(signDataJsonString, sign)

	data, err := pay.PayPost(config.ApiHost+"/order/refund", postData)
	if err != nil {
		return
	}
	resp.SetData(data.Data)
	return resp, err
}

// queryString 回调参数校验
func NotifyVerifyData(dataString string) (response response.Response, err error) {
	var fields []string
	fields = strings.Split(dataString, "&")

	vals := url.Values{}
	for _, field := range fields {
		f := strings.SplitN(field, "=", 2)
		if len(f) >= 2 {
			key, val := f[0], f[1]
			vals.Set(key, val)
		}
	}

	result, err := pay.PublicSha1Verify(vals)

	if err != nil {
		return response, err
	}
	mapInfo := result.(map[string]string)
	for key, value := range mapInfo {
		response.SetKeyValue(key, value)
	}
	return response, err
}
