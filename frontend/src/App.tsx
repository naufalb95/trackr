import { useState, useEffect, useEffectEvent } from "react"
import { getTasks } from "@services/api"

import './App.css'
import TaskList from "@components/TaskList"
import TaskForm from "@components/TaskForm"
import type { Task } from "@type/task"


function App() {
  const [tasks, setTasks] = useState<Task[]>([])
  const [refreshKey, setRefreshKey] = useState(0)
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null)

  const fetchTasks = useEffectEvent(async () => {
    try {
      setLoading(true)
      setError(null)
      const data = await getTasks()
      setTasks(data)
    } catch (err) {
      const errMessage = err instanceof Error ? err.message : "Unknown error";
      console.error("Failed to fetch tasks:", errMessage)
      setError(errMessage);
    } finally {
      setLoading(false)
    }
  })

  useEffect(() => {
    fetchTasks()
  }, [refreshKey])

  const handleTaskCreated = () => {
    setRefreshKey(prev => prev + 1)
  }

  return (
    <div className="App" style={{ maxWidth: '800px', margin: '0 auto', padding: '2rem'}}>
      <h1>Trackr - Task Manager</h1>
      <TaskList tasks={tasks} loading={loading} error={error} />
      <TaskForm onTaskCreated={handleTaskCreated}/>
    </div>
  )
}

export default App
