package main

// 导入所需包
import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

//===========================
// 思路是先定义滑动窗口的两个实体，桶还有窗口
//====================
type baseBucket struct {
	success   int64
	fail      int64
	timeout   int64
	rejection int64
}

// 定义桶 参考https://www.jianshu.com/p/249e4f22fb84
type bucket struct {
	baseBucket  baseBucket
	windowStart int32
}

type SlidingWindow struct {
	buckets   []*bucket    // 他所涵盖的桶，受width约束
	width     int32        // 滑动窗口的宽度
	buckWidth int32        // 桶所统计的时间片长度单位是秒
	tail      int32        // 指向窗口尾部
	mux       sync.RWMutex // 读写锁
}

func NewSlidingWindow(width, buckWidth int32) *SlidingWindow {
	if width < 1 {
		width = 1
	}
	if buckWidth < 1 {
		buckWidth = 1
	}
	return &SlidingWindow{
		width:     width,
		buckWidth: buckWidth,
		buckets:   make([]*bucket, width),
		tail:      0,
	}
}

func (sldwindow *SlidingWindow) getCurrentBucket() *bucket {
	sldwindow.mux.Lock()
	defer sldwindow.mux.Unlock()
	currentSecondTime := time.Now().Unix()
	if len(sldwindow.buckets) == 0 {
		// 确保该tail是0
		sldwindow.tail = 0
		sldwindow.buckets[sldwindow.tail] = &bucket{
			baseBucket: baseBucket{},
			windowStart: int32(currentSecondTime),
		}
		return sldwindow.buckets[sldwindow.tail]
	}
	tail := sldwindow.buckets[sldwindow.tail]
	// 桶对象用来保存[windowStart, windowStart + buckWidth)时间段内的统计信息。
	// 如果当前的时间被该桶囊括，则统计信息落入该桶。
	if int64(tail.windowStart+sldwindow.buckWidth) > currentSecondTime {
		return tail
	}

	for i := int32(0); i < sldwindow.width; i++ {
		tail := sldwindow.buckets[sldwindow.tail]
		if int64(tail.windowStart+sldwindow.buckWidth) > currentSecondTime {
			return tail
		} else if (currentSecondTime - int64((tail.windowStart + sldwindow.buckWidth))) > int64(sldwindow.width*sldwindow.buckWidth) {
			// 如果是长时间没放完，又有了新的访问。导致这些窗口内的桶都过期了。
			// 因此就需要重置一下桶。
			sldwindow.tail = 0
			sldwindow.buckets = make([]*bucket, sldwindow.width)
			return sldwindow.buckets[sldwindow.tail]
		} else {
			sldwindow.tail++
			bucket := &bucket{
				baseBucket:  baseBucket{},
				windowStart: tail.windowStart + sldwindow.buckWidth,
			}
			if sldwindow.tail >= sldwindow.width {
				copy(sldwindow.buckets[:], sldwindow.buckets[1:])
				sldwindow.tail--
			}
			sldwindow.buckets[sldwindow.tail] = bucket
		}
	}
	return sldwindow.buckets[sldwindow.tail]
}

/**
 * @desc 模拟请求成功的情景
 */
func (sldwindow *SlidingWindow) incrSuccess() {
	bucket := sldwindow.getCurrentBucket()
	atomic.AddInt64(&bucket.baseBucket.success, 1)
}

/**
 * @desc 模拟请求失败的情景
 */
func (sldwindow *SlidingWindow) incFail() {
	bucket := sldwindow.getCurrentBucket()
	atomic.AddInt64(&bucket.baseBucket.fail, 1)
}

/**
 * @desc 模拟请求超时的情景
 */
func (sldwindow *SlidingWindow) incrTimeOut() {
	bucket := sldwindow.getCurrentBucket()
	atomic.AddInt64(&bucket.baseBucket.timeout, 1)
}

/**
 * @desc 模拟请求拒绝的情景 429
 */
func (sldwindow *SlidingWindow) incrReject() {
	bucket := sldwindow.getCurrentBucket()
	atomic.AddInt64(&bucket.baseBucket.rejection, 1)
}

func main() {
	// 滑动窗口计数器
	// 用于给熔断器提供数据依据
	rw := NewSlidingWindow(2, 1)
	fmt.Println(time.Now().Unix())

	//test 2
	rw.incrSuccess()
	time.Sleep(time.Second * 3)
	rw.incrSuccess()
	fmt.Printf("%+v,%+v\n", rw.buckets[0], rw.buckets[1])

}
