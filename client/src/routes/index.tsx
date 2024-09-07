import { BrowserRouter, Route, Routes as DomRoutes } from "react-router-dom";
import { Auth } from "../components/Auth";
import { Todo } from "../components/Todo";
import { useEffect } from "react";
import axios from "axios";
import { CsrfToken } from "../types";

export const Routes = () => {
  useEffect(() => {
    axios.defaults.withCredentials = true;
    const getCsrfToken = async () => {
      const { data } = await axios.get<CsrfToken>(
        `${import.meta.env.VITE_API_URL}/csrf`
      );
      axios.defaults.headers.common["X-CSRF-Token"] = data.csrf_token;
    };
    getCsrfToken();
  }, []);

  return (
    <BrowserRouter>
      <DomRoutes>
        <Route path="/" element={<Auth />} />
        <Route path="/todo" element={<Todo />} />
      </DomRoutes>
    </BrowserRouter>
  );
};
