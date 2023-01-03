package concurrent

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type resource struct {
	mutex sync.Mutex
	count int
}

func (resource *resource) CountIncrement() {
	resource.mutex.Lock()
	defer resource.mutex.Unlock()
	resource.count++
}

type acc struct {
	wg  sync.WaitGroup
	sum int
}

func Example() {
	// lock
	var lock = resource{count: 1}
	for i := 0; i < runtime.GOMAXPROCS(4); i++ {
		// goroutine协程，gmp模型，g代表goroutine,m代表machine(类似cpu）,p代表processor(类似线程)。
		// 采用m:n线程模型，m个g，n个p。p有个本地队列和global队列.
		// 调度主要由sysmon的p来监控。
		// 调度方法:抢占式调度、工作量窃取、
		go lock.CountIncrement()
	}
	time.Sleep(100)
	fmt.Println("count:", lock.count)
	// syc.Map channel
	var sm sync.Map = sync.Map{}
	var chs []chan string = make([]chan string, runtime.GOMAXPROCS(4))
	sm.Store("counter", -1)
	for k := 0; k < runtime.GOMAXPROCS(4); k++ {
		var ch chan string = make(chan string)
		chs[k] = ch
		go func(index int) {
			fmt.Println("counter set", index)
			sm.Store("counter", index)
			fmt.Println("ch send msg:", "test"+strconv.Itoa(index))
			ch <- "test" + strconv.Itoa(index)
			defer close(ch)
		}(k)
	}
	for j := 0; j < runtime.GOMAXPROCS(4); j++ {
		// 通道是栈的形式
		chMsg := <-chs[j]
		var counter, ok = sm.Load("counter")
		fmt.Println("chMsg:", chMsg)
		if ok {
			fmt.Println("counter:", counter)
		}
	}
	// sync.pool 缓存,先进先出双向队列，队列元素采用一个环形数组，headTail公用一个字段用，零填充一个缓存行解决伪共享问题,分为private和shared，shared可以工作量窃取。
	// 同时采用victim_cache机制来保证gc后的性能平滑过度(类似于cpu的victim_cache机制)
	// 采用pin和unpin来防止p和对象的绑定过程的并发问题,类似于disruptor
	pool := sync.Pool{}
	m := pool.Get()
	fmt.Println(m)
	var (
		lilyMap  = map[string]int{"lily": 22}
		frankMap = map[string]int{"frank": 44}
	)
	pool.Put(&lilyMap)
	pool.Put(frankMap)
	fmt.Println(pool.Get())
	if frankAge, ok := frankMap["frank"]; ok {
		fmt.Println(frankAge)
	}
	// sync.waitGroup 类似于 cyclicBarrier
	acc := acc{wg: sync.WaitGroup{}, sum: 4}
	pCount := 5
	for i := 0; i < pCount; i++ {
		acc.wg.Add(1)
		go func(index int) {
			acc.sum += index
			acc.wg.Done()
		}(i)
	}
	acc.wg.Wait()
	fmt.Println("sum:", acc.sum)
	// sync.cond
	r := resource{mutex: sync.Mutex{}, count: 4}
	cond := sync.NewCond(&r.mutex)
	for i := 0; i < 5; i++ {
		go func() {
			r.mutex.Lock()
			defer r.mutex.Unlock()
			r.count--
			fmt.Println("count:", r.count)
			if r.count == 0 {
				fmt.Println("wait count:", r.count)
				cond.Wait()
				fmt.Println("signal count:", r.count)
			}
		}()
	}
	go func() {
		for {
			time.Sleep(500)
			if r.count == -1 {
				fmt.Println("broadcast")
				cond.Broadcast()
				break
			}
		}
	}()
	time.Sleep(500000)
	// sync.once
	wg := &acc.wg
	wg.Add(20)
	once := sync.Once{}
	for i := 0; i < 20; i++ {
		go func(k int) {
			fmt.Println("do count", k)
			once.Do(func() {
				fmt.Println("do once")
			})
			wg.Done()
		}(i)
	}
	wg.Wait()
	// atomic
	var num1 int64 = 5
	p := &num1
	// 先找到指针地址，然后再取出指针指向的地址的内容，不是原子操作
	ptr := atomic.LoadInt64(p)
	fmt.Println("ptr", ptr)
}
