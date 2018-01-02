package Singleton

import (
	"sync"
	"time"
)

//Singleton 單體模式
var singleton *string

var singletonLock sync.Mutex

//GetInstance 取得實體
func GetInstance() *string {

	if singleton == nil {
		singletonLock.Lock()
		if singleton == nil {
			var now = time.Now().Format("2006-01-02 15:04:05")
			singleton = &now
		}
	}

	return singleton
}

//ClearInstance 清除實體
func ClearInstance() {
	singleton = nil
}
