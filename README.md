# 一、模块
## 1.1 Server 模块
    Server加载目录`./ware/house/`下的 'data/' 或者 'data_num_priority/'的文件，
    并解析成 *JobInfo* 结构体存储。
    此外，提供 GetJobInfo API接口获取Job工作数据。
## 1.2 Worker 模块
    Worker对Job数据进行调度并执行，核心模块由以下两个组件组成：
    a. Actuator: Task执行管理器，执行最小粒度为Task。
    b. Scheduler: Job调度管理器，执行最小粒度为Job。
    Worker分别支持以下几种创建方式：
    1. NewBaseWorker: 创建一个无容量、忽略优先级的Worker；
    2. NewWorkerWithCapacity: 创建一个有容量、忽略优先级的Worker；
    3. NewWorkerWithSimplePriority: 创建一个有容量、优先级有效的Worker，Worker采用'期待执行时间相同的情况下，才比较优先级'；
    4. NewWorkerWithSmartPriority: 创建一个有容量、智能优先级的Worker，Worker采用'期待执行时间加权优先级比较排序'。

# 二、组件
## 2.1 Task执行管理器 Actuator
    Task执行管理器核心工作是执行满足条件的Task，核心条件为'capacity'容量。
    1. 方法 Ticking() 执行Task
    2. 方法 String() 输出图标功能

## 2.2 Job调度管理器 Scheduler
    Scheduler调度管理器基于*优先队列 Priority Queue*进行Job调度。

# 三、测试用例
    笔试的Task在测试用例 smart_worker_test.go文件中实现。
    **分别实现了 Task 1.2、Task 2.1、Task 2.2、Task 2.3和Task 3.1、Task 3.3。**
    
    0. DataHubServer的方法GetJobInfo: 实现 Task 1.1;
    1. TestNewBaseWorker: 实现Task 1.2;
    2. TestNewWorkerWithCapacity: 实现 Task 2.1;
    3. TestNewWorkerWithSimplePriority: 实现 Task 2.2;
    4. TestNewWorkerWithSmartPriority: 实现 Task 2.3;
    5. TestNewWorkerWithNumPriority: 实现 Task 3.1;
    6. TestNewWorkerWithTaskSpeed: 实现 Task 3.3。