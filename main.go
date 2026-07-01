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
	//V1版本
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
	//V2版本
	{
		//ctx, cancel := context.WithTimeout(context.Background(), time.Second*2) //设置2秒超时
		//defer cancel()
		//
		//pools := V2.NewWorkerPool(ctx, 3)
		//pools.Start()
		//
		//for i := 0; i < 10; i++ {
		//	id := i
		//	err := pools.Submit(func(ctx context.Context) {
		//		select {
		//		case <-ctx.Done():
		//			fmt.Println("task", id, "已取消")
		//		case <-time.After(time.Second):
		//			fmt.Println("task", id, "完成")
		//		}
		//	})
		//	if err != nil {
		//		fmt.Println("Submit failed", id, err)
		//	}
		//}
		//
		//pools.Shutdown()
	}
	{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2) //设置2秒超时
		defer cancel()

		//定义handler
		handler := func(ctx context.Context, id int) (string, error) {
			select {
			case <-ctx.Done():
				return "", ctx.Err()
			case <-time.After(time.Second):
				return fmt.Sprintf("task %d 完成", id), nil
			}
		}

		//创建池子
		p := pool.NewWorkerPool(ctx, 3, handler)
		p.Start()

		//提交：只带参数
		for i := 0; i < 10; i++ {
			f, err := p.Submit(i)
			if err != nil {
				fmt.Println("submit faild:", i, err)
				continue
			}

			go func() {
				msg, err := f.Result(context.Background())
				if err != nil {
					fmt.Println("task", i, "err:", err)
					return
				}
				fmt.Println(msg)
			}()
		}

		//等池子跑完
		p.Shutdown()

	}
}
