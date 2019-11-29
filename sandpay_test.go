package sandpay

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/shop-r1/sandpay/pay/params"
)

func TestSandPay_OrderPay(t *testing.T) {
	type args struct {
		params params.OrderPayParams
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"test01",
			args{
				params: params.OrderPayParams{
					OrderNo:     strconv.Itoa(int(time.Now().UnixNano())),
					TotalAmount: 1,
					Subject:     "话费充值",
					Body:        "用户购买话费0.12",
					TxnTimeOut:  "20191130000000",
					ClientIp:    "103.30.99.197",
					PayMode:     params.PayModWeiXinMp,
					PayExtra: params.PayExtraWeiChat{
						SubAppId: "xxxx",
						OpenId: "xxxx",
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sandPay := &SandPay{
				Config: Client.Config,
			}
			gotResp, err := sandPay.OrderPay(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("SandPay.OrderPay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(gotResp)
		})
	}
}

func TestSandPay_OrderPayQrAlipay(t *testing.T) {
	type args struct {
		params params.OrderPayParams
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"test01",
			args{
				params: params.OrderPayParams{
					OrderNo:     strconv.Itoa(int(time.Now().UnixNano())),
					TotalAmount: 1,
					Subject:     "话费充值",
					Body:        "用户购买话费0.12",
					TxnTimeOut:  "20191130000000",
					ClientIp:    "103.30.99.197",
					PayMode:     params.PayModWeiXinMp,
					PayExtra: params.PayExtraWeiChat{
						SubAppId: "wx29c4121dda03c5c2",
						OpenId: "oU48j5vOQjDBwOfWzkoSfXNSKZUY",
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sandPay := &SandPay{
				Config: Client.Config,
			}
			gotResp, err := sandPay.OrderPayQrAlipay(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("SandPay.OrderPay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(gotResp)
		})
	}
}

func TestSandPay_OrderPayQrWx(t *testing.T) {
	type args struct {
		params params.OrderPayParams
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"test01",
			args{
				params: params.OrderPayParams{
					OrderNo:     strconv.Itoa(int(time.Now().UnixNano())),
					TotalAmount: 1,
					Subject:     "话费充值",
					Body:        "用户购买话费0.12",
					TxnTimeOut:  "20191130000000",
					ClientIp:    "103.30.99.197",
					PayMode:     params.PayModWeiXinMp,
					PayExtra: params.PayExtraWeiChat{
						SubAppId: "wx29c4121dda03c5c2",
						OpenId: "oU48j5vOQjDBwOfWzkoSfXNSKZUY",
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sandPay := &SandPay{
				Config: Client.Config,
			}
			gotResp, err := sandPay.OrderPayQrWx(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("SandPay.OrderPay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(gotResp)
		})
	}
}

func TestSandPay_OrderPayWx(t *testing.T) {
	type args struct {
		params params.OrderPayParams
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"test01",
			args{
				params: params.OrderPayParams{
					OrderNo:     strconv.Itoa(int(time.Now().UnixNano())),
					TotalAmount: 1,
					Subject:     "话费充值",
					Body:        "用户购买话费0.12",
					TxnTimeOut:  "20191130000000",
					ClientIp:    "103.30.99.197",
					PayMode:     params.PayModWeiXinMp,
					PayExtra: params.PayExtraWeiChat{
						SubAppId: "wx29c4121dda03c5c2",
						OpenId: "oU48j5vOQjDBwOfWzkoSfXNSKZUY",
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sandPay := &SandPay{
				Config: Client.Config,
			}
			gotResp, err := sandPay.OrderPayWx(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("SandPay.OrderPay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(gotResp)
		})
	}
}
