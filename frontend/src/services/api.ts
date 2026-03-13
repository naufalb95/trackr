import axios from 'axios';
import type { Task, CreateTaskRequest, TaskStatus } from '@type/task';

const API_BASE_URL = 'http://localhost:8080/';

export async function getTasks(): Promise<Task[]> {
  const url = API_BASE_URL + "api/tasks";
  const { data } = await axios<{ data: Task[]}>({
    method: 'GET',
    url
  })
  return data.data;
}

export async function createTask(task: CreateTaskRequest): Promise<Task> {
  const url = API_BASE_URL + "api/tasks";
  const response = await axios<Task>({
    url,
    method: 'POST',
    data: task,
    headers: {
      'Content-Type': "application/json"
    }
  });

  if (response.status !== 201) {
    throw new Error(`Failed to create task: ${response.statusText}`);
  }

  return response.data;
}

export async function updateTaskStatus(taskId: number, task: TaskStatus): Promise<void> {
  const url = API_BASE_URL + "api/tasks/" + taskId;
  const response = await axios<Task>({
    url,
    method: 'PATCH',
    data: { status: task },
    headers: {
      'Content-Type': "application/json"
    }
  });

  if (response.status !== 200) {
    throw new Error(`Failed to update task: ${response.statusText}`);
  }
}

export async function deleteTask(taskId: number): Promise<void> {
  const url = API_BASE_URL + "api/tasks/" + taskId;
  await axios.delete(url)
}