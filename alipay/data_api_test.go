package alipay

import (
	"testing"

	pay "github.com/whrwsoftware/payment"
	"github.com/whrwsoftware/payment/pkg/xlog"
)

func TestClient_DataBillDownloadUrlQuery(t *testing.T) {
	bm := make(pay.BodyMap)
	bm.Set("bill_type", "trade").
		Set("bill_date", "2016-04-05")
	rsp, err := client.DataBillDownloadUrlQuery(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("rsp:", rsp)
}
