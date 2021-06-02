package alipay

import (
	"encoding/json"
	"fmt"
	"github.com/iGoogle-ink/gopay"
)

// alipay.marketing.campaign.cash.create(创建现金活动)
//    文档地址：https://opendocs.alipay.com/apis/api_5/alipay.marketing.campaign.cash.create
func (a *Client) CampaignCashCreate(bm gopay.BodyMap) (aliRsp *CampaignCashCreateResponse, err error) {
	err = bm.CheckEmptyError("total_money", "total_num")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.marketing.campaign.cash.create"); err != nil {
		return nil, err
	}
	aliRsp = new(CampaignCashCreateResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.marketing.campaign.cash.trigger(触发现金红包活动)
//    文档地址：https://opendocs.alipay.com/apis/api_5/alipay.marketing.campaign.cash.trigger
func (a *Client) CampaignCashTrigger(bm gopay.BodyMap) (aliRsp *CampaignCashTriggerResponse, err error) {
	err = bm.CheckEmptyError("crowd_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.marketing.campaign.cash.trigger"); err != nil {
		return nil, err
	}
	aliRsp = new(CampaignCashTriggerResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.marketing.campaign.cash.detail.query(现金活动详情查询)
//    文档地址：https://opendocs.alipay.com/apis/api_5/alipay.marketing.campaign.cash.detail.query
func (a *Client) CampaignCashDetailQuery(bm gopay.BodyMap) (aliRsp *CampaignCashDetailQueryResponse, err error) {
	err = bm.CheckEmptyError("crowd_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.marketing.campaign.cash.detail.query"); err != nil {
		return nil, err
	}
	aliRsp = new(CampaignCashDetailQueryResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}