/**
  @author:
  @date:
  @note:
**/

package main

import (
	"context"
	"ekit/pool"
	"fmt"
	"time"
)

func main() {

	{
		////启动同时执行3个任务的线程池
		//pools, wg := pool.StartPool(3)
		//
		//for i := 0; i < 10; i++ {
		//	//把值拷贝给新变量，避免闭包问题
		//	val := i
		//	wg.Add(1)
		//	//向channel里面传入数据，worker里面的for循环会收到数据
		//	pools <- func() {
		//		time.Sleep(time.Second)
		//		fmt.Println("task", val)
		//	}
		//}
		//
		//wg.Wait()
		//close(pools)
		//fmt.Println("all done")
	}

	{
		//pools := V1.NewWorkerPool(3)
		//pools.Start()
		//
		//for i := 0; i < 10; i++ {
		//	i := i
		//	pools.Submit(func() {
		//		time.Sleep(time.Second)
		//		fmt.Println("V1 task", i)
		//	})
		//}
		//
		//pools.Shutdown()
		//fmt.Println("pools V1 shutdown")
	}

	{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) //设置2秒超时
		defer cancel()

		pools := pool.NewWorkerPool(ctx, 3)
		pools.Start()

		for i := 0; i < 10; i++ {
			id := i
			err := pools.Submit(func(ctx context.Context) {
				select {
				case <-ctx.Done():
					fmt.Println("task", id, "已取消")
				case <-time.After(time.Second):
					fmt.Println("task", id, "完成")
				}
			})
			if err != nil {
				fmt.Println("Submit failed", id, err)
			}
		}

		pools.Shutdown()
	}
}
