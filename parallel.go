package gomodulesdemo

import "sync"

func ForEach(list []interface{}, f func(val interface{})) {

	var wg sync.WaitGroup
	wg.Add(len(list))

	for _, val := range list {
		go func(v interface{}) {
			f(v)
			wg.Done()
		}(val)
	}
	wg.Wait()
}
