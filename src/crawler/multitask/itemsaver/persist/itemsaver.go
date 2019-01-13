package persist

import "log"

func ItemSaver() chan interface{} {
	// ItemSaver 使用 out chan 和 Engine 通信
	out := make(chan interface{})
	// real save logic
	go func() {
		itemCount := 0
		for {
			item := <- out
			itemCount ++
			log.Printf("ItemSaver : got item #%d : %v \n",itemCount,item)
		}
	}()
	return out
}