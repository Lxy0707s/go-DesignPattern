package store

import "sync"

type ICommodity interface {
	SendCommodity(uId, commodityId, bizId string, extMap map[string]string)
}

// Manager 需要思考改造
type Manager struct {
	lock       sync.RWMutex
	taskServer []ICommodity
}

func (m *Manager) run() {

}
