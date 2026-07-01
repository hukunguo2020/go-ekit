package V2

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

/*
	自己封装的线程池
*/

// 定义任务
type Task func(ctx context.Context)

// 定义协程池对象
type WorkerPool struct {
	//任务列表
	tasks chan func(ctx context.Context)
	//任务等待
	wg     sync.WaitGroup
	taskWg sync.WaitGroup
	//同时执行的任务数
	workerNum int
	//全局上下本（用于取消控制）
	ctx context.Context
	//用于关闭整个池子
	cancel context.CancelFunc
}

func NewWorkerPool(parent context.Context, workerNum int) *WorkerPool {
	ctx, cancel := context.WithCancel(parent)
	return &WorkerPool{
		//给 channel 加缓冲（1024），避免 Submit 阻塞调用方
		tasks:     make(chan func(ctx context.Context), 1024),
		workerNum: workerNum,
		ctx:       ctx,
		cancel:    cancel,
	}
}

// worker
// 增加了取消
// 增加了捕获错误
func (p *WorkerPool) worker(id int) {
	for {
		select {
		case <-p.ctx.Done(): //收到取消信号
			fmt.Println("worker", id, "exit")
			p.drainTasks() //把没执行的任务 Done 掉
			return
		case task, ok := <-p.tasks:
			if !ok {
				//通道关闭
				return
			}

			p.runTask(task)
		}
	}
}

// 运行任务+捕获错误
func (p *WorkerPool) runTask(task Task) {
	defer p.taskWg.Done()

	//捕获错误
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("worker panic:", err)
		}
	}()

	task(p.ctx)
}

// worker 退出前清空队列
func (p *WorkerPool) drainTasks() {
	for {
		select {
		case task, ok := <-p.tasks:
			if !ok {
				return
			}
			//池子已关闭，不再执行，把计数减掉
			p.taskWg.Done()
			_ = task //不执行
		default:
			return //队列空了
		}
	}
}

// 启动worker
func (p *WorkerPool) Start() {
	for i := 0; i < p.workerNum; i++ {
		//val:=i
		p.wg.Add(1)
		go func(id int) {
			defer p.wg.Done()
			p.worker(id)
		}(i)
	}
}

// 提交任务
func (p *WorkerPool) Submit(task func(ctx context.Context)) error {
	// 1. 池子关了就不接任务
	select {
	case <-p.ctx.Done():
		// 线程池已关闭
		return errors.New("worker pool closed")
	default:
	}

	// 2.登记新任务要执行
	p.taskWg.Add(1)

	// 3.再入队；入队失败要回滚
	select {
	case p.tasks <- task:
		return nil
	case <-p.ctx.Done():
		// 线程池已关闭

		p.taskWg.Done() //回滚，否则 Shutdown 死锁
		return errors.New("worker pool closed")
	}
}

// 优雅关闭
func (p *WorkerPool) Shutdown() {
	// 1. 不再接任务（Submit 要检测 closed）
	close(p.tasks)
	// 2. 等已取出的任务跑完
	p.taskWg.Wait()

	// 3. 让 worker 退出
	p.cancel()
	p.wg.Wait()
}
