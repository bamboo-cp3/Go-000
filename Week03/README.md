### 题目
基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出

分析：
1. 启动http server 和 linux signal 信号注册 分别启动goroutine去处理
2. http server 退出 linux signal同时退出 反之亦然
3. 通过context 携带取消信号，做到同时取消
