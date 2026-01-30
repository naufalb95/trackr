import axios from 'axios';
import type { Task, CreateTaskDTO } from '@type/task';

const API_BASE_URL = 'https://api.example.com/tasks';

export async function getTasks(): Promise<Task[]> {
  const { data } = await axios.get<Task[]>(API_BASE_URL);
  return data;
}

export async function createTask(task: CreateTaskDTO): Promise<Task> {
  const response = await axios({
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