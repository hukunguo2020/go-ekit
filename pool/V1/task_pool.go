package V1

import (
	"sync"
)

/*
	自己封装的线程池
*/

// 定义任务
type Task func()

// 定义协程池对象
type WorkerPool struct {
	//任务列表
	tasks chan Task
	//任务等待
	wg sync.WaitGroup
	//同时执行的任务数
	workerNum int
}

func NewWorkerPool(workerNum int) *WorkerPool {
	return &WorkerPool{
		tasks:     make(chan Task),
		workerNum: workerNum,
	}
}

// 启动worker
func (p *WorkerPool) Start() {
	for i := 0; i < p.workerNum; i++ {
		//val:=i
		go func() {
			for task := range p.tasks {
				task()
				p.wg.Done()
			}
		}()
	}
}

// 提交任务
func (p *WorkerPool) Submit(task Task) {
	//wg计数器+1
	p.wg.Add(1)
	//传入任务
	p.tasks <- task
}

// 优雅关闭
func (p *WorkerPool) Shutdown() {
	close(p.tasks)
	p.wg.Wait()
}
