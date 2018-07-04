package util

import (
	"github.com/ServiceComb/go-chassis/core/lager"
	"sync"
)

//ThreadGroupWait realise a thread group wait
type ThreadGroupWait struct {
	count int
	mtx   sync.Mutex
	cond  *sync.Cond
}

//NewThreadGroupWait is a function which initializes value for threadgroupwait struct and returns it
func NewThreadGroupWait() *ThreadGroupWait {
	tmp := new(ThreadGroupWait)
	tmp.count = 1
	tmp.cond = sync.NewCond(&tmp.mtx)
	return tmp
}

//Add is a method to add a thread waitgroup
func (this *ThreadGroupWait) Add(count int) {
	this.mtx.Lock()
	defer this.mtx.Unlock()
	this.count++
}

//Done is a method to say that waitgroup is done
func (this *ThreadGroupWait) Done() {
	this.mtx.Lock()
	defer this.mtx.Unlock()
	this.count--
	if this.count < 0 {
		this.cond.Broadcast()
	}
}

//Wait is a method which waits until done function is called
func (this *ThreadGroupWait) Wait() {
	this.mtx.Lock()
	defer this.mtx.Unlock()
	this.cond.Wait()
}

//RoutineTask interface
type RoutineTask interface {
	Svc(agrs interface{}) interface{}
}

//RoutineManager is a struct
type RoutineManager struct {
	wg *ThreadGroupWait
}

//NewRoutineManager is a fucntion which initializes value for routine manager struct
func NewRoutineManager() *RoutineManager {
	tmp := new(RoutineManager)
	tmp.wg = NewThreadGroupWait()
	return tmp
}

//Wait is method which waits for until done function is called
func (this *RoutineManager) Wait() {
	this.wg.Wait()
}

//Spawn is a method which spawns new routine
func (this *RoutineManager) Spawn(task RoutineTask, agrs interface{}, routineName string) {
	lager.Logger.Info("Routine spawn:" + routineName)
	this.wg.Add(1)
	go this.spawn(task, agrs, routineName)
}

func (this *RoutineManager) spawn(task RoutineTask, agrs interface{}, routineName string) {
	task.Svc(agrs)
	this.wg.Done()
	lager.Logger.Info("Routine exit:" + routineName)
}

//Done is a method which tells waitgroup that it's done waiting
func (this *RoutineManager) Done() {
	this.wg.Done()
}
