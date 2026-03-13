export interface Task {
  id: number;
  title: string;
  description: string;
  status: TaskStatus;
  createdAt: Date;
  updatedAt: Date;
}

export type TaskStatus = 'todo' | 'in_progress' | 'done';

export type CreateTaskRequest = Omit<Task, 'id' | 'createdAt' | 'updatedAt'>;

export type UpdateTaskRequest = Partial<Omit<Task, 'id' | 'createdAt' | 'updatedAt'>>;