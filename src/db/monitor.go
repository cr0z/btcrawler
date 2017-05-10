/**
 * Created by 93201 on 2017/5/10.
 */
package db

import "github.com/x-croz/log"

var monitor *TaskMonitor

type TaskMonitor struct {
	tasks map[string]task
}

func (m *TaskMonitor) init() {
	m.tasks = make(map[string]task)
	for _, v := range m.tasks {
		v.init()
	}
}

func (m *TaskMonitor) start() {
	for _, v := range m.tasks {
		go v.start()
	}
}

func (m *TaskMonitor) addTask(name string, t task) {
	m.tasks[name] = t
}

func (m *TaskMonitor) Send(name string, d interface{}) {
	if m.tasks[name] != nil {
		m.tasks[name].onCall(d)
	}
}
func (m *TaskMonitor) Stop() {
	for _, v := range m.tasks {
		v.stop()
	}
}
func Monitor() *TaskMonitor {
	if monitor == nil {
		log.Panic("db task monitor not initailize")
	}
	return monitor
}
