package adapter

type MQAdapter struct {
}

/*func (m *MQAdapter) filter(strJson string, link map[string]string) func(strJson string, link map[string]string) {
 return m.filter(strJson, link)
}*/

func (m *MQAdapter) filter(obj map[string]string, link map[string]string) *RebateInfo {
	rebateInfo := NewRebateInfo()

	for key, l := range link {
		switch key {
		case "userId":
			rebateInfo.SetUserId(l)
		case "bizId":
			rebateInfo.SetBizId(l)
		case "bizTime":
			rebateInfo.SetBizTime(l)
		case "desc":
			rebateInfo.SetDesc(l)
		}
	}
	return rebateInfo
}
