import axios from "axios";

export function getErrorMessage(error: unknown): string {
  if (axios.isAxiosError(error)) {
    // Tu backend devuelve errorMessage dentro de data
    const backendError = (error.response?.data as any)?.errorMessage;
    return (
      backendError ||
      error.message ||
      "Unexpected error"
    );
  }

  if (error instanceof Error) {
    return error.message;
  }

  return "Unexpected error";
}
