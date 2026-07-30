import axios from 'axios';

export function getErrorMessage(
  error: unknown,
  fallback = 'Something went wrong. Please try again.'
): string {
  if (axios.isAxiosError(error)) {
    const data = error.response?.data as { errorMessage?: string; message?: string } | undefined;
    const backendError = data?.errorMessage || data?.message;

    if (backendError) return backendError;

    if (!error.response) {
      return "We couldn't connect to Nexora. Check your connection and try again.";
    }

    switch (error.response.status) {
      case 403:
        return "You don't have permission to perform this action.";
      case 404:
        return 'This resource could not be found or is no longer available.';
      case 409:
        return 'The operation conflicts with the current state of the resource.';
      default:
        if (error.response.status >= 500) {
          return 'Something went wrong while processing your request.';
        }
    }
  }

  if (error instanceof Error) {
    return error.message;
  }

  return fallback;
}
