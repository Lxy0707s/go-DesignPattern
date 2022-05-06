package adapter

type RebateInfo struct {
	userId  string // 用户ID
	bizId   string // 业务ID
	bizTime string // 业务时间
	desc    string // 业务描述
}

var RebateInfoInstance *RebateInfo

func NewRebateInfo() *RebateInfo {
	if RebateInfoInstance == nil {
		RebateInfoInstance = &RebateInfo{
			userId:  "0",
			bizId:   "001",
			bizTime: "0-0-0-1",
			desc:    "test",
		}
	}
	return RebateInfoInstance
}

func (r *RebateInfo) SetUserId(userId string) {
	r.userId = userId
}

func (r *RebateInfo) SetBizId(bizId string) {
	r.bizId = bizId
}

func (r *RebateInfo) SetBizTime(bizTime string) {
	r.bizTime = bizTime
}

func (r *RebateInfo) SetDesc(desc string) {
	r.desc = desc
}

func (r *RebateInfo) GetUserId() string {
	return r.userId
}

func (r *RebateInfo) GetBizId() string {
	return r.bizId
}

func (r *RebateInfo) GetBizTime() string {
	return r.bizTime
}

func (r *RebateInfo) GetDesc() string {
	return r.desc
}
