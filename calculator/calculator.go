package calculator

type Calculator struct {
	Input  <-chan int
	Output chan<- int
}

func (c *Calculator) Start() {
	go func() {
		for i := range c.Input {
			c.Output <- i*i
		}
		close(c.Output)
	}()
}
