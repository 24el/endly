package model

import (
	"sync/atomic"
	"sync"
)

//Process represents a workflow execution process.
type Process struct {
	*Workflow
	Pipeline *Pipeline
	*Activities
	Terminated int32
	Scheduled  *Task
	*ExecutionError
}

//Terminate flags current workflow as terminated
func (p *Process) Terminate() {
	atomic.StoreInt32(&p.Terminated, 1)
}

//CanRun returns true if current workflow can run
func (p *Process) CanRun() bool {
	return !(p.IsTerminated() || p.Scheduled != nil)
}

//IsTerminated returns true if current workflow has been terminated
func (p *Process) IsTerminated() bool {
	return atomic.LoadInt32(&p.Terminated) == 1
}

//Push adds supplied activity
func (p *Process) Push(activity *Activity) {
	if p.Workflow != nil {
		activity.Caller = p.Workflow.Name
	}
	if p.Pipeline != nil {
		activity.Caller = p.Pipeline.Name
	}
	p.Activities.Push(activity)
}

//NewProcess creates a new workflow, pipeline process
func NewProcess(workflow *Workflow, pipeline *Pipeline) *Process {
	return &Process{
		ExecutionError:      &ExecutionError{},
		Workflow:   workflow,
		Pipeline:   pipeline,
		Activities: NewActivities(),
	}
}

//processes  represents running workflow/pipe process stack.
type Processes struct {
	mux       *sync.RWMutex
	processes []*Process
}

//Push adds a workflow to the workflow stack.
func (p *Processes) Push(process *Process) {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.processes = append(p.processes, process)
}

//Pop removes the first workflow from the workflow stack.
func (p *Processes) Pop() *Process {
	p.mux.Lock()
	defer p.mux.Unlock()
	if len(p.processes) == 0 {
		return nil
	}
	var result = (p.processes)[len(p.processes)-1]
	p.processes = p.processes[0: len(p.processes)-1]
	return result
}

//LastRun returns the last workflow process.
func (p *Processes) Last() *Process {
	p.mux.RLock()
	defer p.mux.RUnlock()
	for i := len(p.processes) - 1; i >= 0; i-- {
		if p.processes[i].Workflow != nil {
			return p.processes[i]
		}
	}
	return nil
}

//LastRun returns the first workflow process.
func (p *Processes) First() *Process {
	p.mux.RLock()
	defer p.mux.RUnlock()
	for i := 0; i < len(p.processes); i++ {
		if p.processes[i].Workflow != nil {
			return p.processes[i]
		}
	}
	return nil
}

//NewProcesses creates a new processes
func NewProcesses() *Processes {
	return &Processes{
		processes: make([]*Process, 0),
		mux:       &sync.RWMutex{},
	}
}

//Error represent workflow error
type ExecutionError struct {
	Error    string
	Caller   string
	TaskName string
}