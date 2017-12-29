package Singleton

import "sync"

//Singleton 單體模式
type Singleton struct{}

var singleton *Singleton

var singletonLock sync.Mutex

//GetInstance 取得實體
func GetInstance() *Singleton {

	if singleton == nil {
		singletonLock.Lock()
		if singleton == nil {
			singleton = &Singleton{}
		}
	}

	return singleton
}

//ClearInstance 清除實體
func ClearInstance() {
	singleton = nil
}
