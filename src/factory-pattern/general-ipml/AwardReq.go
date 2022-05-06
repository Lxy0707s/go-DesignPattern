package general_ipml

import "sync"

type AwardReq struct {
	UId         string            // 用户唯一ID
	AwardType   int               // 奖品类型(可以用枚举定义)；1优惠券、2实物商品、3第三方兑换卡(爱奇艺)
	AwardNumber string            // 奖品编号；sku、couponNumber、cardId
	BizId       string            // 业务ID，防重复
	extMap      map[string]string // 扩展信息
}

var awardReqInstance *AwardReq
var once sync.Once

func NewAwardReq() *AwardReq {
	if awardReqInstance == nil {
		awardReqInstance = &AwardReq{
			UId:         "",
			AwardType:   1,
			AwardNumber: "",
			BizId:       "",
			extMap:      make(map[string]string),
		}
		return awardReqInstance
	}
	return awardReqInstance
}

func (a *AwardReq) getUId() string {
	return a.UId
}

func (a *AwardReq) setUId(uId string) {
	a.UId = uId
}

func (a *AwardReq) getAwardType() int {
	return a.AwardType
}

func (a *AwardReq) setAwardType(awardType int) {
	a.AwardType = awardType
}

func (a *AwardReq) getAwardNumber() string {
	return a.AwardNumber
}

func (a *AwardReq) setAwardNumber(awardNumber string) {
	a.AwardNumber = awardNumber
}

func (a *AwardReq) getBizId() string {
	return a.BizId
}

func (a *AwardReq) setBizId(bizId string) {
	a.BizId = bizId
}

func (a *AwardReq) getExtMap() map[string]string {
	return a.extMap
}

func (a *AwardReq) setExtMap(extMap map[string]string) {
	a.extMap = extMap
}
