import { FormEvent, useEffect } from "react";
import { useQueryClient } from "@tanstack/react-query";
import {
  ArrowRightOnRectangleIcon,
  ShieldCheckIcon,
} from "@heroicons/react/24/solid";

import { TaskItem } from "./TaskItem";
import {
  TasksQueryKey,
  useTaskMutation,
  useTasksQuery,
} from "../hooks/useTask";
import { useAuth } from "../hooks/useAuth";
import { useStore } from "../store";

export const Todo = () => {
  const queryClient = useQueryClient();
  const { editedTask } = useStore();
  const updateTask = useStore((state) => state.updateEditedTask);
  const { data, isLoading } = useTasksQuery();
  const { createTaskMutation, updateTaskMutation } = useTaskMutation();
  const { logoutMutation } = useAuth();
  const submitTaskHandler = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (editedTask.id === 0) createTaskMutation.mutate(editedTask);
    else {
      updateTaskMutation.mutate(editedTask);
    }
  };
  const logout = async () => {
    await logoutMutation.mutateAsync();
    queryClient.removeQueries({
      queryKey: [TasksQueryKey],
    });
  };

  return (
    <div className="flex justify-center items-center flex-col min-h-screen text-gray-600 font-mono">
      <div className="flex items-center my-3">
        <ShieldCheckIcon className="h-8 w-8 mr-3 text-indigo-500 cursor-pointer" />
        <span className="text-center text-3xl font-extrabold">
          Task Manager
        </span>
      </div>
      <ArrowRightOnRectangleIcon
        onClick={logout}
        className="h-6 w-6 my-6 text-blue-500 cursor-pointer"
      />
      <form onSubmit={submitTaskHandler}>
        <input
          className="mb-3 mr-3 px-3 py-2 border border-gray-300"
          placeholder="title ?"
          type="text"
          onChange={(e) => updateTask({ ...editedTask, title: e.target.value })}
          value={editedTask.title || ""}
        />
        <button
          className="disabled:opacity-40 mx-3 py-2 px-3 text-white bg-indigo-600 rounded"
          disabled={!editedTask.title}
        >
          {editedTask.id === 0 ? "Create" : "Update"}
        </button>
      </form>
      {isLoading ? (
        <p>Loading...</p>
      ) : (
        <ul className="my-5">
          {data?.map((task) => (
            <TaskItem key={task.id} task={task} />
          ))}
        </ul>
      )}
    </div>
  );
};
