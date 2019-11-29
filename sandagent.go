package sandpay

import (
	"github.com/shop-r1/sandpay/agent"
	"github.com/shop-r1/sandpay/agent/params"
	"github.com/shop-r1/sandpay/agent/request"
	"github.com/shop-r1/sandpay/agent/response"
	"log"
	"time"
)

var AgentPayClient PaymentAgent

func init() {
	//TODO 改为配置文件
	AgentPayClient.Config = agent.Config{
		Version:         "1.0",
		MerId:           "S6887745",
		PrivatePath:     "/Users/linwenxiang/Downloads/certs/server.key", //私钥文件
		CertPath:        "/Users/linwenxiang/Downloads/certs/server.crt", //公钥文件
		EncryptCertPath: "/Users/linwenxiang/Downloads/certs/sand.cer",   //导出的公钥文件
		//ApiHost:         "https://caspay.sandpay.com.cn/agent-main/openapi",
		ApiHost:         "https://cashier.sandpay.com.cn",
		//ApiHost:   "http://61.129.71.103:8003/agent-main/openapi",
		NotifyUrl: "https://www.baidu.com",
		FrontUrl:  "https://www.baidu.com",
	}
	agent.LoadCertInfo(&AgentPayClient.Config)
}

type PaymentAgent struct {
	Config agent.Config
}

// 实时代付
func (sandPay *PaymentAgent) AgentPay(params params.AgentPayParams) (resp response.AgentPayResponse, err error) {
	timeString := time.Now().Format("20060102150405")

	body := request.AgentPayBody{
		Version:      agent.VERSION,
		OrderCode:    params.OrderCode,
		ProductId:    agent.PRODUCTID_AGENTPAY_TOC,
		CurrencyCode: agent.CURRENCY_CODE,
		TranTime:     timeString,
		TimeOut:      params.TimeOut,
		TranAmt:      params.TranAmt,
		AccAttr:      params.AccAttr,
		AccNo:        params.AccNo,
		AccType:      params.AccType,
		AccName:      params.AccName,
		BankName:     params.BankName,
		Remark:       params.Remark,
		PayMode:      params.PayMode,
		ChannelType:  params.ChannelType,
	}

	signDataJsonString := agent.GenerateSignString(body)
	//log.Println("请求提中 encryptData 原始数据  ", signDataJsonString)
	encryptData, sign, encryptKey, err := agent.PrivateSha1SignData(signDataJsonString)
	postData := agent.GeneratePostData(encryptData, encryptKey, agent.AGENT_PAY, sandPay.Config.MerId, sign)
	//log.Println("发送的请求体body内容", postData)
	data, err := agent.PayPost(sandPay.Config.ApiHost+"/agentpay", postData)
	//fmt.Println("过滤 \" 解析内容", data, err)
	if err != nil {
		return
	}
	resp.SetData(data.Data)
	return resp, err
}

// 订单查询
func (sandPay *PaymentAgent) AgentPayQuery(tranTime string, orderCode string) (resp response.QueryOrderResponse, err error) {

	body := request.QueryOrderBody{
		Version:   agent.VERSION,
		ProductId: agent.PRODUCTID_AGENTPAY_TOC,
		TranTime:  tranTime,
		OrderCode: orderCode,
		Extend:    "",
	}

	signDataJsonString := agent.GenerateSignString(body)
	log.Println(signDataJsonString)
	encryptData, sign, encryptKey, err := agent.PrivateSha1SignData(signDataJsonString)
	postData := agent.GeneratePostData(encryptData, encryptKey, agent.ORDER_QUERY, sandPay.Config.MerId, sign)

	data, err := agent.PayPost(sandPay.Config.ApiHost+"/queryOrder", postData)
	if err != nil {
		return
	}
	resp.SetData(data.Data)
	return resp, err
}
