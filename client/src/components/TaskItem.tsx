import { memo } from "react";
import { PencilIcon, TrashIcon } from "@heroicons/react/24/solid";

import { Task } from "../types";
import { useStore } from "../store";
import { useTaskMutation } from "../hooks/useTask";

const MemoComponent = ({ task }: { task: Task }) => {
  const updateTask = useStore((state) => state.updateEditedTask);
  const { deleteTaskMutation } = useTaskMutation();
  return (
    <li className="my-3">
      <span className="font-bold">{task.title}</span>
      <div className="flex float-right ml-20">
        <PencilIcon
          className="h-5 w-5 mx-1 text-blue-500 cursor-pointer"
          onClick={() => {
            updateTask(task);
          }}
        />
        <TrashIcon
          className="h-5 w-5 text-blue-500 cursor-pointer"
          onClick={() => {
            deleteTaskMutation.mutate(task);
          }}
        />
      </div>
    </li>
  );
};
export const TaskItem = memo(MemoComponent);
