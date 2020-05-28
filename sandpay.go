package sandpay

import (
	"fmt"
	"github.com/shop-r1/sandpay/pay"
	"github.com/shop-r1/sandpay/pay/params"
	"github.com/shop-r1/sandpay/pay/request"
	"github.com/shop-r1/sandpay/pay/response"
	"net/url"
	"strings"
	"time"
)

var Client SandPay

type SandPay struct {
	Config pay.Config
}

// 微信统一下单接口
func (sandPay *SandPay) OrderPayWx(params params.OrderPayParams) (resp response.OrderPayResponse, err error) {
	config := sandPay.Config
	timeString := time.Now().Format("20060102150405")

	header := request.Header{}
	header.SetMethod(`sandpay.trade.pay`).SetVersion(`1.0`).SetAccessType("1")
	header.SetChannelType("07").SetMid(config.MerId).SetProductId("00000005").SetReqTime(timeString)
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

// 聚合统一下单接口
func (sandPay *SandPay) OrderPayQr(params params.OrderPayParams) (resp response.OrderPayResponse, err error) {
	config := sandPay.Config
	timeString := time.Now().Format("20060102150405")

	header := request.Header{}
	header.SetMethod(`sandpay.trade.precreate`).SetVersion(`1.0`).SetAccessType("1")
	header.SetChannelType("07").SetMid(config.MerId).SetProductId("00000012").SetReqTime(timeString)
	body := request.OrderPayBody{
		PayTool:     "0401",
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
	fmt.Println(signDataJsonString)
	sign, _ := pay.PrivateSha1SignData(signDataJsonString)
	postData := pay.GeneratePostData(signDataJsonString, sign)

	fmt.Println(postData)
	data, err := pay.PayPost(config.ApiHost+"/qr/api/order/create", postData)
	if err != nil {
		return
	}
	resp.SetData(data.Data)
	return resp, err
}

func (sandPay *SandPay) OrderPayWechat(params params.OrderPayParams) (resp response.OrderPayResponse, err error) {
	config := sandPay.Config
	timeString := time.Now().Format("20060102150405")

	header := request.Header{}
	header.SetMethod(`sandpay.trade.pay`).SetVersion(`1.0`).SetAccessType("1")
	header.SetChannelType("07").SetMid(config.MerId).SetProductId("00002020").SetReqTime(timeString)
	body := request.OrderPayBody{
		//PayTool:     "0401",
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
	fmt.Println(signDataJsonString)
	sign, _ := pay.PrivateSha1SignData(signDataJsonString)
	postData := pay.GeneratePostData(signDataJsonString, sign)

	fmt.Println(postData)
	data, err := pay.PayPost(config.ApiHost+"/qr/api/order/create", postData)
	if err != nil {
		return
	}
	resp.SetData(data.Data)
	return resp, err
}

// 聚合统一下单接口
func (sandPay *SandPay) OrderPayH5(params params.OrderPayParams) (resp response.OrderPayResponse, err error) {
	config := sandPay.Config
	timeString := time.Now().Format("20060102150405")

	header := request.Header{}
	header.SetMethod(`sandpay.trade.orderCreate`).SetVersion(`1.0`).SetAccessType("1")
	header.SetChannelType("07").SetMid(config.MerId).SetProductId("00002000").SetReqTime(timeString)
	body := request.OrderPayBody{
		//PayTool:     "0401",
		OrderCode:   params.OrderNo,
		TotalAmount: params.GetTotalAmountToString(),
		Subject:     params.Subject,
		Body:        params.Body,
		TxnTimeOut:  params.TxnTimeOut,
		//PayMode:     params.PayMode,
		//PayExtra:    params.PayExtra.ToJson(),
		//ClientIp:    params.ClientIp,
		NotifyUrl: sandPay.Config.NotifyUrl,
		FrontUrl:  sandPay.Config.FrontUrl,
		Extends:   params.Extends,
	}

	signDataJsonString := pay.GenerateSignString(body, header)

	sign, _ := pay.PrivateSha1SignData(signDataJsonString)

	data, err := pay.PayPostRedirect(config.ApiHost+"/gw/web/order/create", signDataJsonString, sign)
	if err != nil {
		return
	}
	resp.Body.QrCode = data.Data
	return resp, err
}

// 聚合统一下单接口
func (sandPay *SandPay) OrderPayH5K(params params.OrderPayParams) (resp response.OrderPayResponse, err error) {
	config := sandPay.Config
	timeString := time.Now().Format("20060102150405")

	header := request.Header{}
	header.SetMethod(`sandpay.trade.precreate`).SetVersion(`1.0`).SetAccessType("1")
	header.SetChannelType("07").SetMid(config.MerId).SetProductId("00000016").SetReqTime(timeString)
	body := request.OrderPayBody{
		PayTool:     "0401",
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
	fmt.Println(signDataJsonString)
	sign, _ := pay.PrivateSha1SignData(signDataJsonString)
	postData := pay.GeneratePostData(signDataJsonString, sign)

	fmt.Println(postData)
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
		PayTool:     "0401",
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
