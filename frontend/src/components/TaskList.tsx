import { useEffect, useState } from "react";
import { getTasks } from "@services/api";

import type { Task } from "@type/task";

function TaskList({ refreshTrigger }: { refreshTrigger?: number }) {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    async function fetchTasks() {
      try {
        setLoading(true);
        const data = await getTasks();
        setTasks(data);
      } catch (err) {
        setError(err instanceof Error ? err.message : 'An error occurred.')
      } finally {
        setLoading(false);
      }
    }

    fetchTasks()
  }, [refreshTrigger])

  if (loading) {
    return <div>Loading tasks...</div>
  }

  if (error) {
    return <div>Error: {error}</div>
  }

  return (
    <div>
      <h2>Tasks ({tasks.length})</h2>
      {tasks.length === 0 ? (
        <p>No tasks yet! Create one above.</p>
      ) : (
        <ul style={{ listStyle: 'none', padding: 0 }}>
          {tasks.map((task) => {
            console.log('📝 Rendering task:', task)
            return (
              <li 
                key={task.id}
                style={{
                  padding: '1rem',
                  marginBottom: '0.5rem',
                  border: '1px solid #ddd',
                  borderRadius: '4px',
                  backgroundColor: '#f9f9f9',
                }}
              >
                <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                  <strong style={{ fontSize: '1.1rem' }}>{task.title}</strong>
                  <span 
                    style={{
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
                  >
                    {task.status === 'in_progress' ? 'In Progress' : 
                     task.status === 'todo' ? 'To Do' : 'Done'}
                  </span>
                </div>
                {task.description && (
                  <p style={{ margin: '0.5rem 0 0 0', color: '#666' }}>
                    {task.description}
                  </p>
                )}
              </li>
            )
          })}
        </ul>
      )}
    </div>
  )
}

export default TaskList;