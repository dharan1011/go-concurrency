package scenarios

import "fmt"
import "sync"
import "time"

type SharedResource struct {
}

func (sr *SharedResource) GetAccess() {
	time.Sleep(time.Second * 10)
}

func main() {
	sr := SharedResource{}
	var rmu sync.RWMutex
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			rmu.RLock()
			fmt.Println("Read Lock Acquired")
			defer rmu.RUnlock()
			defer wg.Done()
			sr.GetAccess()
		}()
	}
	wg.Add(1)
	go func() {
		rmu.Lock()
		fmt.Println("Lock Acquired")
		defer rmu.Unlock()
		defer wg.Done()
		sr.GetAccess()
	}()
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			rmu.RLock()
			fmt.Println("Read Lock Acquired")
			defer rmu.RUnlock()
			defer wg.Done()
			sr.GetAccess()
		}()
	}
	wg.Wait()
}
