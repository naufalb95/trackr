import { useState } from "react";

import './App.css'
import TaskList from "@components/TaskList"
import TaskForm from "@components/TaskForm"

function App() {
  const [refreshKey, setRefreshKey] = useState(0)

  const handleTaskCreated = () => {
    setRefreshKey(prev => prev + 1)
  }

  return (
    <div className="App" style={{ maxWidth: '800px', margin: '0 auto', padding: '2rem'}}>
      <h1>Trackr - Task Manager</h1>
      <TaskList refreshTrigger={refreshKey} />
      <TaskForm onTaskCreated={handleTaskCreated}/>
    </div>
  )
}

export default App
