package main

//func main() {
//	var group sync.WaitGroup
//	ch := make(chan int, 5)
//	group.Add(1)
//	go func() {
//		for i := 0; i < 5; i++ {
//			ch <- i
//		}
//		close(ch)
//	}()
//	go func() {
//		rand.Int()
//		for v:=range ch{
//			fmt.Println(v)
//		}
//		group.Done()
//	}()
//	group.Wait()
//
//}
