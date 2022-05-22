
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
    1. 

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