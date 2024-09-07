import axios, { AxiosError } from "axios";
import { useStore } from "../store";
import { useError } from "./useError";
import { Credentials } from "../types";
import { useMutation } from "@tanstack/react-query";
import { useNavigate } from "react-router-dom";

export function useAuth() {
  const navigate = useNavigate();
  const resetEditedTask = useStore((state) => state.resetEditedTask);
  const { switchErrorHandling } = useError();

  const loginMutation = useMutation({
    mutationFn: (credentials: Credentials) =>
      axios.post(`${import.meta.env.VITE_API_URL}/login`, credentials),

    onSuccess: () => {
      navigate("/todo");
    },
    onError: (error: AxiosError<{ message: string }>) => {
      if (error.response?.data.message) {
        switchErrorHandling(error.response.data.message);
      } else {
        alert("An error occurred. Please try again.");
      }
    },
  });

  const registerMutation = useMutation({
    mutationFn: (credentials: Credentials) =>
      axios.post(`${import.meta.env.VITE_API_URL}/signup`, credentials),

    onSuccess: () => {
      navigate("/");
    },
    onError: (error: AxiosError<{ message: string }>) => {
      if (error.response?.data.message) {
        switchErrorHandling(error.response.data.message);
      } else {
        alert("An error occurred. Please try again.");
      }
    },
  });

  const logoutMutation = useMutation({
    mutationFn: () => axios.post(`${import.meta.env.VITE_API_URL}/logout`),

    onSuccess: () => {
      resetEditedTask();
      navigate("/");
    },
    onError: (error: AxiosError<{ message: string }>) => {
      if (error.response?.data.message) {
        switchErrorHandling(error.response.data.message);
      } else {
        alert("An error occurred. Please try again.");
      }
    },
  });

  return {
    loginMutation,
    registerMutation,
    logoutMutation,
  };
}
