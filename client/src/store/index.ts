import { create } from "zustand";
import { Task } from "../types";

type State = {
  editedTask: Task;
  updateEditedTask: (payload: Task) => void;
  resetEditedTask: () => void;
};

export const useStore = create<State>((set) => ({
  editedTask: {
    id: 0,
    title: "",
    createdAt: new Date(),
    updatedAt: new Date(),
  },
  updateEditedTask: (payload) => set({ editedTask: payload }),
  resetEditedTask: () =>
    set({
      editedTask: {
        id: 0,
        title: "",
        createdAt: new Date(),
        updatedAt: new Date(),
      },
    }),
}));
