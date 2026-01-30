import { useEffect, useState } from "react";
import { getTasks } from "@services/api";

import type { Task } from "@type/task";

function TaskList() {
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
  }, [])

  if (loading) {
    return <div>Loading tasks...</div>
  }

  if (error) {
    return <div>Error: {error}</div>
  }

  return (
    <div>
      <h2>Tasks</h2>
      {tasks.length === 0 ? (
        <p>No tasks yet!</p>
      ) : (
        <ul>
          <li>
            {tasks[0].id}
          </li>
          {tasks.map((task) => (
            <li key={task.id}>
              <strong>{task.title}</strong> - {task.status}
              <br/>
              <small>{task.description}</small>
            </li>
          ))}
        </ul>
      )}
    </div>
  )
}

export default TaskList;