package main

import (
	"sync/atomic"
	"GoJob"
	"time"
	"fmt"
)

type Counter struct {
	atomicValue uint64;
}

func (counter * Counter) increment(){
	atomic.AddUint64(&counter.atomicValue,1);
}

// WorkerInterface implementation!
func (counter * Counter) Job(){
	counter.increment();
}

func newCounter() *Counter  {
	return &Counter{atomicValue:0};
}

func (counter Counter) getAtomicValue() uint64{
	return atomic.LoadUint64(&counter.atomicValue);
}


func (counter * Counter) print()  {
	fmt.Println("Value is ",counter.atomicValue);
}

var goJobManager = GoJob.SingletonJobManager();

func main() {
	newCounterPtr := newCounter();
	jobName := "incrementNumber";
	//Creates new job with given job name and allows 4 workers to run concurrently
	goJobManager.NewJob(jobName,4);
	//I want to increment 100 times so I need to add 100 times...
	for i:=0;i<100;i++ {
		//Add new task to job
		goJobManager.AddTask(jobName,newCounterPtr);
	}
	//Start task with given jobName
	goJobManager.StartTasks(jobName);
	//GoJob will run task on  goroutines, to see this I sleep main thread for 1 s.
	time.Sleep(1*time.Second)
	newCounterPtr.print();
}



