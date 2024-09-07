import axios from "axios";
import { useNavigate } from "react-router-dom";
import { useStore } from "../store";

export function useError() {
  const navigate = useNavigate();
  const resetEditedTask = useStore((state) => state.resetEditedTask);
  const getCsrfToken = async () => {
    const { data } = await axios.get(`${import.meta.env.VITE_API_URL}/csrf`);
    axios.defaults.headers.common["X-CSRF-Token"] = data.csrf_token;
  };

  function switchErrorHandling(message: string) {
    switch (message) {
      case "invalid csrf token":
        getCsrfToken();
        alert("Please try again.");
        break;

      case "invalid or expired jwt":
        alert("Please log in again.");
        resetEditedTask();
        navigate("/");
        break;

      case "missing or malformed jwt":
        alert("Please log in.");
        resetEditedTask();
        navigate("/");
        break;

      case "duplicate key not allowed":
        alert("email already exist, please use another one.");
        break;

      case "crypto/bcrypto: hashedPassword is not the hash of the given password":
        alert("password is not correct.");
        break;

      case "record not found.":
        alert("email is not correct.");
        break;

      default:
        alert(message);
    }
  }

  return { switchErrorHandling };
}
