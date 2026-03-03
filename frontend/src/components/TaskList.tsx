import type { Task, TaskStatus } from "@type/task";
import type { ChangeEvent, MouseEvent } from "react";

interface TaskListParams {
  tasks: Task[],
  loading: boolean,
  error: string | null,
  onTaskDeleted: (taskId: number) => void
  onTaskStatusUpdated: (taskId: number, taskStatus: TaskStatus) => void
}

function TaskList(params: TaskListParams) {
  const { tasks, loading, error, onTaskDeleted, onTaskStatusUpdated } = params

  if (loading) {
    return <div>Loading tasks...</div>
  }

  if (error) {
    return <div>Error: {error}</div>
  }

  const handleTaskDelete = (e: MouseEvent, taskId: number) => {
    e.stopPropagation()
    onTaskDeleted(taskId)
  }

  const handleTaskStatusUpdate = (e: ChangeEvent, taskId: number, taskStatus: string) => {
    e.stopPropagation()
    let validTaskStatus: TaskStatus

    switch (taskStatus) {
      case 'todo':
      case 'in_progress':
      case 'done':
        validTaskStatus = taskStatus
        break
      default:
        throw new Error("Invalid task status")
    }

    onTaskStatusUpdated(taskId, validTaskStatus)
  }

  return (
    <div>
      <h2>Tasks ({tasks.length})</h2>
      {tasks.length === 0 ? (
        <p>No tasks yet! Create one above.</p>
      ) : (
        <ul style={{ listStyle: 'none', padding: 0 }}>
          {tasks.map((task) => {
            return (
              <li 
                key={task.id}
                style={{
                  padding: '1rem',
                  marginBottom: '0.5rem',
                  border: '1px solid #ddd',
                  borderRadius: '4px',
                  backgroundColor: '#242424',
                }}
              >
                <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                  <strong style={{ fontSize: '1.1rem' }}>{task.title}</strong>
                  <select style={{
                      width: '25%',
                      padding: '0.25rem 0.75rem',
                      borderRadius: '12px',
                      fontSize: '0.85rem',
                      backgroundColor: 
                        task.status === 'done' ? '#d4edda' :
                        task.status === 'in_progress' ? '#fff3cd' :
                        '#d1ecf1',
                      color:
                        task.status === 'done' ? '#155724' :
                        task.status === 'in_progress' ? '#856404' :
                        '#0c5460',
                    }}
                      onChange={(e) => handleTaskStatusUpdate(e, task.id, e.target.value)}>
                    <option value='todo'>To Do</option>
                    <option value='in_progress'>In Progress</option>
                    <option value='done'>Done</option>
                  </select>
                </div>
                <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                  {task.description && (
                    <p style={{ margin: '0.5rem 0 0 0', color: '#efefef', textAlign: 'justify' }}>
                      {task.description}
                    </p>
                  )}
                  <button
                    style={{
                      width: '25%',
                      padding: '0.25rem 0.75rem',
                      borderRadius: '12px',
                      fontSize: '0.85rem',
                      backgroundColor: '#ff0000',
                      color: '#efefef',
                    }}
                    onClick={(e) => handleTaskDelete(e, task.id)}
                  >
                    Delete
                  </button>
                </div>
              </li>
            )
          })}
        </ul>
      )}
    </div>
  )
}

export default TaskList;