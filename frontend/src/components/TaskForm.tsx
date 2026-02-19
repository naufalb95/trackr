import { useState } from "react";
import { createTask } from "@services/api";

import type { SubmitEvent } from "react";
import type { CreateTaskDTO, TaskStatus } from "@/types/task";

interface TaskFormProps {
  onTaskCreated: () => void
}

function TaskForm({ onTaskCreated }: TaskFormProps) {
  const [title, setTitle] = useState('')
  const [description, setDescription] = useState('')
  const [status, setStatus] = useState<TaskStatus>('todo')
  const [isSubmitting, setIsSubmitting] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const handleSubmit = async (e: SubmitEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (!title.trim()) {
      setError("Title is required")
      return
    }

    try {
      setIsSubmitting(true)
      setError(null)

      const taskData: CreateTaskDTO = {
        title: title.trim(),
        description: description.trim(),
        status
      }

      await createTask(taskData)

      setTitle('')
      setDescription('')
      setStatus('todo')

      onTaskCreated()
    } catch (err) {
      console.error("Error when creating task:", err)
      setError(err instanceof Error ? err.message : "Failed to create task.")
    } finally {
      setIsSubmitting(false);
    }
  }

  return (
    <div
      style={{
        marginBottom: '2rem',
        padding: '1rem',
        border: '1px solid #ccc',
        borderRadius: '8px'
      }}
    >
      <h3>Create New Task</h3>

      <form onSubmit={handleSubmit}>
        <div style={{ marginBottom: '1rem'}}>
          <label htmlFor='title' style={{
            display: 'block',
            marginBottom: '0.5rem'
          }}>
            Title *
          </label>
          <input 
            id="title"
            type="text"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            placeholder="Enter task title"
            disabled={isSubmitting}
            style={{ width: '100%', padding: '0.5rem', fontSize: '1rem' }}
          />
        </div>

        <div style={{ marginBottom: '1rem'}}>
          <label htmlFor='description' style={{
            display: 'block',
            marginBottom: '0.5rem'
          }}>
            Description
          </label>
          <textarea 
            id="description"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            placeholder="Enter task description"
            disabled={isSubmitting}
            rows={3}
            style={{ width: '100%', padding: '0.5rem', fontSize: '1rem' }}
          />
        </div>

        <div style={{ marginBottom: '1rem' }}>
          <label htmlFor="status" style={{ display: 'block', marginBottom: '0.5rem' }}>
            Status
          </label>
          <select
            id="status"
            value={status}
            onChange={(e) => setStatus(e.target.value as TaskStatus)}
            disabled={isSubmitting}
            style={{ width: '100%', padding: '0.5rem', fontSize: '1rem' }}
          >
            <option value="todo">To Do</option>
            <option value="in_progress">In Progress</option>
            <option value="done">Done</option>
          </select>
        </div>

        {error && (
          <div style={{ color: 'red', marginBottom: '1rem' }}>{error}</div>
        )}

        <button
          type="submit"
          disabled={isSubmitting}
          style={{
            padding: '0.75rem 1.5rem',
            fontSize: '1rem',
            backgroundColor: isSubmitting ? '#ccc' : '#007bff',
            color: 'white',
            border: 'none',
            borderRadius: '4px',
            cursor: isSubmitting ? 'not-allowed' : 'pointer',
          }}
        >
          {isSubmitting ? 'Creating...' : 'Create Task'}
        </button>
      </form>
    </div>
  )
}

export default TaskForm