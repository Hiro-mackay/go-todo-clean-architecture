import axios from "axios";
import { Task } from "../types";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useStore } from "../store";
import { useError } from "./useError";

export const TasksQueryKey = "tasks";

export const useTasksQuery = () => {
  const getTasks = async () => {
    const { data } = await axios.get<Task[]>(
      `${import.meta.env.VITE_API_URL}/tasks`,
      {
        withCredentials: true,
      }
    );
    return data;
  };

  return useQuery<Task[], Error>({
    queryKey: [TasksQueryKey],
    queryFn: getTasks,
    staleTime: Infinity,
  });
};

export const useTaskMutation = () => {
  const queryClient = useQueryClient();
  const { switchErrorHandling } = useError();
  const resetEditedTask = useStore((state) => state.resetEditedTask);

  const createTaskMutation = useMutation({
    onSuccess: async (newTask: Task) => {
      await queryClient.cancelQueries({
        queryKey: [TasksQueryKey],
      });

      queryClient.setQueryData<Task[]>([TasksQueryKey], (old) => [
        ...(old || []),
        newTask,
      ]);
      resetEditedTask();
    },
    onError: (error, newTask, context: any) => {
      queryClient.setQueryData([TasksQueryKey], context.previousTasks);
      switchErrorHandling(error?.message);
    },
    mutationFn: (newTask: Task) =>
      axios
        .post(`${import.meta.env.VITE_API_URL}/tasks`, newTask, {
          withCredentials: true,
        })
        .then((res) => res.data),
  });

  const updateTaskMutation = useMutation({
    onSuccess: async (updatedTask: Task) => {
      await queryClient.cancelQueries({
        queryKey: [TasksQueryKey],
      });
      const previousTasks = queryClient.getQueryData<Task[]>([TasksQueryKey]);
      queryClient.setQueryData<Task[]>([TasksQueryKey], (old) =>
        old?.map((task) => (task.id === updatedTask.id ? updatedTask : task))
      );
      resetEditedTask();
      return { previousTasks };
    },
    onError: (error, updatedTask, context: any) => {
      queryClient.setQueryData([TasksQueryKey], context.previousTasks);
      switchErrorHandling(error?.message);
    },
    mutationFn: (updatedTask: Task) =>
      axios
        .put(
          `${import.meta.env.VITE_API_URL}/tasks/${updatedTask.id}`,
          updatedTask,
          {
            withCredentials: true,
          }
        )
        .then((res) => res.data),
  });

  const deleteTaskMutation = useMutation({
    onSuccess: async (deletedTask: Task) => {
      await queryClient.cancelQueries({
        queryKey: [TasksQueryKey],
      });
      const previousTasks = queryClient.getQueryData<Task[]>([TasksQueryKey]);
      queryClient.setQueryData<Task[]>([TasksQueryKey], (old) =>
        old?.filter((task) => task.id !== deletedTask.id)
      );
      return { previousTasks };
    },
    onError: (error, deletedTask, context: any) => {
      queryClient.setQueryData([TasksQueryKey], context.previousTasks);
      switchErrorHandling(error?.message);
    },
    mutationFn: (deletedTask: Task) =>
      axios
        .delete(`${import.meta.env.VITE_API_URL}/tasks/${deletedTask.id}`, {
          withCredentials: true,
        })
        .then(() => deletedTask),
  });

  return {
    createTaskMutation,
    updateTaskMutation,
    deleteTaskMutation,
  };
};
