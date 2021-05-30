package calculator

type Calculator struct {
	Input  <-chan int
	Output chan<- int
}

func (c *Calculator) Start() {
	//running anonymous go routine with the inner loop
	go func() { 
		//loop will read from channel c.Input till it's closed by sender
		for i := range c.Input {
			c.Output <- i*i
		}
		//close the c.Output when after the c.Input was closed
		close(c.Output)
	}()
}
