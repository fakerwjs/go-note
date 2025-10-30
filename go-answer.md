### 1️⃣ `new` 和 `make` 的区别

**答案**：

| 特性   | `new`             | `make`                         |
| ---- | ----------------- | ------------------------------ |
| 用途   | 分配内存，返回指向类型的指针    | 初始化 slice、map、channel，返回对应类型   |
| 返回值  | `*T`              | `T` (引用类型，如 slice/map/channel) |
| 初始化  | 返回零值指针，未初始化内部数据结构 | 初始化内部数据结构，可直接使用                |
| 使用场景 | 普通变量/结构体          | slice/map/channel 等引用类型        |

**注意**：

* `new(T)` → `*T`
* `make([]int, 5)` → `[]int`，可直接 append

---

### 2️⃣ Go 源码编译到可执行二进制流程

**答案**：

1. **解析 (Parser)**：将源码解析成 AST（抽象语法树）。
2. **类型检查 (Type Check)**：检查类型一致性。
3. **编译 (Compiler)**：

   * 将 AST 转为中间表示 SSA (Static Single Assignment)。
   * SSA → 汇编代码
4. **链接 (Linker)**：

   * 链接标准库、依赖库 → 生成 ELF/PE/Mach-O 可执行文件。
5. **输出二进制**：

   * 可执行文件包含 Go runtime（调度器、GC、栈管理等）。

---

### 3️⃣ Go 调度模型（G、P、M）

**答案**：

* **G (Goroutine)**：用户态轻量级线程。
* **M (Machine)**：绑定操作系统线程的实体。
* **P (Processor)**：逻辑处理器，管理可运行的 G 队列。

**工作流程**：

* 每个 M 绑定一个 P，P 调度 G 运行。
* M 空闲 → 从全局 G 队列抢 G。
* 多 P 存在 → 负载均衡 G。
* 支持 M 切换、work stealing。

---

### 4️⃣ sync.Mutex 实现及 CPU 指令

**答案**：

* 底层依赖 **原子操作**（atomic.CompareAndSwapInt32）实现。
* Linux/AMD64 常用 **`LOCK CMPXCHG`** 指令。
* 内部状态：

  * 0：未锁
  * 1：锁定
  * 2：有等待队列（fast path / slow path）
* 避免内核态切换（自旋 + 阻塞）。

---

### 5️⃣ Go map vs C++ map

* Go map → **哈希表**，查找 O(1) 平均
* C++ map → **红黑树**，查找 O(log n)
* Go map 遍历无序，C++ map 有序
* Go map 支持动态扩容，但并发读写需加锁

---

### 6️⃣ context 包作用

* **用途**：控制 goroutine 生命周期、取消操作、传递请求范围信息（如超时、TraceID）
* 常见用法：

```go
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()
```

* **注意**：未调用 `cancel()` → goroutine 可能泄漏。

---

### 7️⃣ Go 内存泄漏场景

* 未释放的 channel 或 goroutine
* `context` 未 cancel
* 全局 map/slice 持有大量对象引用
* 闭包引用外部对象导致无法 GC

---

### 8️⃣ goroutine 完美退出

* 使用 **context** + **select**

```go
func worker(ctx context.Context) {
  for {
    select {
    case <-ctx.Done():
      return
    default:
      // do work
    }
  }
}
```

* 可避免阻塞和泄漏

---

### 9️⃣ slice、map、array 区别与底层

| 类型    | 底层                      | 备注         |
| ----- | ----------------------- | ---------- |
| array | 连续内存                    | 值类型        |
| slice | 结构体 {ptr, len, cap}     | 引用类型，可动态扩容 |
| map   | 哈希表 + bucket + overflow | 引用类型，随机访问  |

---

### 10️⃣ map 并发安全

* **不安全**，并发写会 panic
* Go 不设计锁的原因：

  * 提醒开发者使用 channel 或 sync.Map
  * 避免隐藏 bug

---

### 11️⃣ 协程池/控制 goroutine 数量

* 使用 **channel 限流**

```go
sem := make(chan struct{}, 10) // 限制10个并发
for _, job := range jobs {
  sem <- struct{}{}
  go func(j Job) {
    defer func(){ <-sem }()
    do(j)
  }(job)
}
```

* 或第三方库（ants、tunny）

---

### 12️⃣ channel vs 锁

* **channel**：消息传递，适合 **数据共享**或流水线
* **锁**：数据互斥，适合 **共享状态**
* 优先使用 channel 避免死锁

---

### 13️⃣ panic 与 recover

* **recover** 只能捕获同一个 goroutine 的 panic
* defer 中使用：

```go
defer func() {
  if r := recover(); r != nil {
    fmt.Println("Recovered:", r)
  }
}()
```

* 不是所有 panic 都能 recover（如 runtime.Goexit）

---

### 14️⃣ GC 机制

* **三色标记 + 增量标记**
* 优化方式：

  * 减少临时对象
  * reuse slice/map
  * 避免全局引用

---

### 15️⃣ interface 底层

* **eface**：非泛型 interface，存放 type 和 data
* **iface**：非空 interface，包含指向 type descriptor 和 data pointer
* **原理**：动态分发，支持多态

---

### 16️⃣ slice 扩容策略

* 初次扩容 len<1024 → cap*2
* len>=1024 → cap*1.25
* **性能关键点**：

  * 避免频繁 append
  * 尽量预分配 cap

---

### 17️⃣ defer 执行顺序

* **后进先出 (LIFO)**
* 开销：少量，避免循环大量 defer
* 场景：

  * 释放资源
  * 错误处理
  * 日志输出

---

### 18️⃣ reflect 使用

* 动态获取类型/值
* 慎用原因：

  * 低性能
  * 破坏类型安全
* 使用场景：

  * 序列化（JSON）
  * ORM
  * 泛型兼容

---

### 19️⃣ Go 在游戏后端优势

* 高并发（goroutine + channel）
* 内存占用低
* 快速编译
* 内置 GC 和 runtime
* 对比：

  * **C++**：高性能，但开发复杂
  * **Java**：GC 停顿长，线程重

---

### 20️⃣ 游戏/高并发服务设计要点

* **高并发连接**：使用 epoll + goroutine +协程池
* **性能瓶颈**：

  * GC 停顿
  * map 并发访问
  * 频繁网络 IO
* **日志/监控**：

  * Zap/Logrus + Prometheus + Grafana
* **微服务**：

  * gRPC / HTTP
  * 限流：token bucket
  * 熔断：Hystrix 风格
* **sync.Map vs map+锁**：

  * sync.Map 适合读多写少
  * map+锁 更灵活，写多性能好


你希望我直接生成这个文档吗？
