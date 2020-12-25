package engine

type EventLoop struct {
	messageQueue []Command

	finish    chan int
	finishing bool
}

func (el *EventLoop) run() {
	for {
		if el.finishing && len(el.messageQueue) == 0 {
			el.finish <- 0
			return
		}
		if len(el.messageQueue) != 0 {
			el.messageQueue[0].Execute(el)
			el.messageQueue = el.messageQueue[1:]
		}
	}
}

func (el *EventLoop) Start() {
	el.finish = make(chan int)
	go el.run()
}

func (el *EventLoop) Post(cmd Command) {
	el.messageQueue = append(el.messageQueue, cmd)
}

func (el *EventLoop) AwaitFinish() {
	el.finishing = true
	<-el.finish
}
