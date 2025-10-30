/**
1.任务队列实现
编写一个 Go 程序，实现一个任务队列（例如处理玩家请求或游戏事件）。支持如下功能：
多个生产者 goroutine 将任务提交到队列。
多个工作者 goroutine 从队列中取任务并处理。
支持优先级任务（优先级高的任务先处理）。
支持优雅退出：在收到关闭信号后，仍处理已提交任务，然后退出。
要求：使用 channel、context、select 等语言特性实现，并注意资源释放、并发安全。

2.限流 + 并发控制
某游戏后端服务需要处理高并发请求。编写一个限流器模块，要求如下：
限制每秒最多 N 个请求进入。
当请求超过限流阈值时，直接拒绝（返回错误）。
并发控制：最多同时处理 M 个任务，其余排队或拒绝。
要求实现为通用包，能够在多个服务中复用。
要求：用 Go 实现；写单元测试验证限流逻辑。

3.键值缓存 + 过期机制
实现一个简单的线程安全缓存结构，用于游戏服务中存储玩家临时数据。要求：
支持 Get(key), Set(key,value,ttl) 操作。
ttl 是存活时间，超过后自动过期并删除。
并发安全，支持大量并发读写。
可选：最大容量限制，超过容量后按 LRU 或类似策略淘汰。
要求：使用 Go map +锁或 sync.Map + 过期检查 goroutine实现；注意资源管理。

4.服务健康检查 +监控指标暴露
在一个假定的游戏服务中，编写 Go 程序：
启动 HTTP 服务暴露 /metrics 接口，返回 JSON 格式的监控指标（例如：当前 Goroutine 数、请求总数、平均处理时间）。
在处理主业务逻辑（模拟）时更新这些指标。
支持 /healthz 接口返回 “ok” 或 “error” 表示服务健康状态。
要求：使用 context 控制超时，使用 go pprof 或 runtime 包获取指标。
可选：将监控数据每 X 秒导出日志。

5.分布式一致性：Leader 选举模拟
为游戏分布式模块（如匹配服务）模拟一个简化版本的 Leader 选举机制。要求：
多个节点（可用 goroutine 模拟）竞选成为 Leader。
Leader 每隔一段时间发送心跳给 Follower，若心跳丢失，Follower 发起新选举。
支持节点加入 / 离开。
要求用 Go 实现 Goroutine、channel、select 多路复用；注意竞态条件。
*/


/**
题目1：
  1.任务队列实现
  编写一个 Go 程序，实现一个任务队列（例如处理玩家请求或游戏事件）。支持如下功能：
  多个生产者 goroutine 将任务提交到队列。
  多个工作者 goroutine 从队列中取任务并处理。
  支持优先级任务（优先级高的任务先处理）。
  支持优雅退出：在收到关闭信号后，仍处理已提交任务，然后退出。
  要求：使用 channel、context、select 等语言特性实现，并注意资源释放、并发安全。
*/

//思路:
/**
1.使用 channel 作为任务队列。
2.使用 goroutine 模拟多个生产者和消费者。
3.用优先级队列（可以用 container/heap）存储任务，保证高优先级任务先处理。
4.优雅退出：
  使用 context.Context 控制整体任务队列生命周期。
  消费者 goroutine 在退出前处理完队列内的任务。
*/
package main

import (
    "container/heap"
    "context"
    "fmt"
    "math/rand"
    "sync"
    "time"
)

type Task struct {
    priority int
    name     string
    index    int
}

// 实现 heap.Interface
type TaskHeap []*Task

func (h TaskHeap) Len() int { return len(h) }
func (h TaskHeap) Less(i, j int) bool {
    return h[i].priority > h[j].priority // 优先级高的先出
}
func (h TaskHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
    h[i].index = i
    h[j].index = j
}
func (h *TaskHeap) Push(x interface{}) {
    n := len(*h)
    task := x.(*Task)
    task.index = n
    *h = append(*h, task)
}
func (h *TaskHeap) Pop() interface{} {
    old := *h
    n := len(old)
    task := old[n-1]
    *h = old[0 : n-1]
    return task
}

func worker(ctx context.Context, id int, tasks chan *Task, wg *sync.WaitGroup) {
    defer wg.Done()
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("worker %d exiting\n", id)
            return
        case task := <-tasks:
            fmt.Printf("worker %d processing task: %s\n", id, task.name)
            time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
        }
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())
    taskChan := make(chan *Task, 100)
    var wg sync.WaitGroup
    ctx, cancel := context.WithCancel(context.Background())

    // 启动 3 个消费者
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go worker(ctx, i, taskChan, &wg)
    }

    // 模拟生产者
    go func() {
        for i := 0; i < 20; i++ {
            taskChan <- &Task{name: fmt.Sprintf("task-%d", i), priority: rand.Intn(10)}
        }
    }()

    time.Sleep(time.Second * 2)
    cancel() // 发出退出信号
    wg.Wait()
}

