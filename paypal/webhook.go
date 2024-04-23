package paypal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ldpool/gopay"
)

// 验证 Webhook 签名
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/webhooks/v1/#verify-webhook-signature_post
func (c *Client) VerifyWebhookSignature(ctx context.Context, bm gopay.BodyMap) (ppRsp *WebHookRsp, err error) {
	res, bs, err := c.doPayPalPost(ctx, bm, verifyWebhookSignature)
	if err != nil {
		return nil, err
	}
	ppRsp = &WebHookRsp{Code: Success}
	ppRsp.Response = new(WebHookVerify)
	if err = json.Unmarshal(bs, ppRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusCreated {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}
