package engine

type EventLoop struct {
	messageQueue []Command

	ready     bool
	waiting   chan int
	finish    chan int
	finishing bool
}

func (el *EventLoop) run() {
	for {
		select {
		case <-el.waiting:
			el.messageQueue[0].Execute(el)
			el.messageQueue = el.messageQueue[1:]
		default:
			if el.finishing && len(el.messageQueue) == 0 {
				el.finish <- 0
				return
			}
			if len(el.messageQueue) == 0 {
				el.ready = true
			} else {
				el.messageQueue[0].Execute(el)
				el.messageQueue = el.messageQueue[1:]
			}
		}
	}
}

func (el *EventLoop) Start() {
	el.waiting = make(chan int)
	el.finish = make(chan int)
	go el.run()
}

func (el *EventLoop) Post(cmd Command) {
	el.messageQueue = append(el.messageQueue, cmd)
	if el.ready {
		el.ready = false
		el.waiting <- 0
	}
}

func (el *EventLoop) AwaitFinish() {
	el.finishing = true
	<-el.finish
}
