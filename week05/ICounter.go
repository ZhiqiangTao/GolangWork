package week05

import (
	"sync"
	"time"
)

// ICounter 计数器接口
type ICounter interface {
	// Update 计数器具备更新值的能力
	Update(newVal float64)
}

// Counter 计数器模型
type Counter struct {
	//多个g上报数据的时候，设计到数据安全的问题，需要用锁在保证临界区的互斥性
	Locker *sync.RWMutex
	//计数器桶，key：时间戳 value:该时间下计数的值(存的是对象指针，方便修改)
	Buckets map[int64]*BucketValue
}

// BucketValue 把每个桶的值以struct保存
type BucketValue struct {
	Value float64
}

// Increment 计数器模型Counter提供了增加计数器桶值的能力
func (c *Counter) Increment(newValue float64) {
	if newValue <= 0 {
		return //0没有意义，直接pass
	}

	//上写锁，确保临界区数据的互斥性
	c.Locker.Lock()
	defer c.Locker.Unlock()

	//获取最新的当前时间戳下的桶内的value
	cur := c.getCurrentBucketValue()
	cur.Value += newValue
	c.removeOldBucketValue()
}

// 获取当前时间戳下的，桶内的value
func (c *Counter) getCurrentBucketValue() *BucketValue {
	//获取当前时间戳
	now := time.Now().Unix()
	var bv *BucketValue
	var exist bool

	if bv, exist = c.Buckets[now]; !exist {
		//如果当前时间戳，在桶中并没有记录，那么就新建一个
		bv = &BucketValue{Value: 0} //初始为0
		c.Buckets[now] = bv
	}
	return bv
}

// 移除时间窗口10秒以外的桶数据, 实现随着时间的推移，保证滑动窗口一直固定在10s
func (c *Counter) removeOldBucketValue() {
	now := time.Now().Unix() - 10
	for unix := range c.Buckets {
		if unix <= now {
			delete(c.Buckets, unix)
		}
	}
}
