package queue

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Task 任务
type Task struct {
	ID        string      `json:"id"`
	Type      string      `json:"type"`
	Payload   interface{} `json:"payload"`
	Status    string      `json:"status"` // pending, running, completed, failed
	Progress  int         `json:"progress"`
	Result    interface{} `json:"result"`
	Error     string      `json:"error"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

// TaskQueue 任务队列
type TaskQueue struct {
	tasks map[string]*Task
	mu    sync.RWMutex
}

// NewTaskQueue 创建任务队列
func NewTaskQueue() *TaskQueue {
	return &TaskQueue{
		tasks: make(map[string]*Task),
	}
}

// Add 添加任务
func (q *TaskQueue) Add(taskType string, payload interface{}) string {
	q.mu.Lock()
	defer q.mu.Unlock()

	taskID := fmt.Sprintf("task_%d", time.Now().UnixNano())

	task := &Task{
		ID:        taskID,
		Type:      taskType,
		Payload:   payload,
		Status:    "pending",
		Progress:  0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	q.tasks[taskID] = task

	return taskID
}

// Get 获取任务
func (q *TaskQueue) Get(taskID string) (*Task, error) {
	q.mu.RLock()
	defer q.mu.RUnlock()

	task, exists := q.tasks[taskID]
	if !exists {
		return nil, fmt.Errorf("任务不存在: %s", taskID)
	}

	return task, nil
}

// Update 更新任务状态
func (q *TaskQueue) Update(taskID string, status string, progress int, result interface{}, errMsg string) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	task, exists := q.tasks[taskID]
	if !exists {
		return fmt.Errorf("任务不存在: %s", taskID)
	}

	task.Status = status
	task.Progress = progress
	task.Result = result
	task.Error = errMsg
	task.UpdatedAt = time.Now()

	return nil
}

// List 列出任务
func (q *TaskQueue) List(status string) []*Task {
	q.mu.RLock()
	defer q.mu.RUnlock()

	var tasks []*Task
	for _, task := range q.tasks {
		if status == "" || task.Status == status {
			tasks = append(tasks, task)
		}
	}

	return tasks
}

// Delete 删除任务
func (q *TaskQueue) Delete(taskID string) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	if _, exists := q.tasks[taskID]; !exists {
		return fmt.Errorf("任务不存在: %s", taskID)
	}

	delete(q.tasks, taskID)

	return nil
}

// Process 处理任务（示例）
func (q *TaskQueue) Process(ctx context.Context, taskID string, handler func(payload interface{}) (interface{}, error)) error {
	task, err := q.Get(taskID)
	if err != nil {
		return err
	}

	// 更新状态为运行中
	q.Update(taskID, "running", 0, nil, "")

	// 执行任务
	result, err := handler(task.Payload)
	if err != nil {
		q.Update(taskID, "failed", 0, nil, err.Error())
		return err
	}

	// 更新状态为完成
	q.Update(taskID, "completed", 100, result, "")

	return nil
}
