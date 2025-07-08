import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query"
import api, { HttpApiError } from "@/api"

const keys = {
    profile: () => ["profile"],
}

export function useSignup() {
    const { mutate: signup, isPending, error } = useMutation({
        mutationFn: async (data: api.ModelsSignupRequest) => {
            const res = await api.signup(data)
            return res.data as api.ModelsSignupResponse
        },
        onSuccess: (data) => {
            setToken(data.token)
        }
    })

    return { signup, isPending, error }
}

export function useUser() {
  const qc = useQueryClient();
  const token = getToken();
  if (!token) {
    return {
      isAuthenticated: false,
      profile: null,
      isLoading: false,
      isError: false,
      error: null,
      logout: () => {},
      refresh: () => {},
    }
  }

  const logout = () => {
    clearToken();
    qc.removeQueries({ queryKey: keys.profile() });
  };

  const query = useQuery({
    queryKey: keys.profile(),
    queryFn: async () => {
        try {
            const res = await api.getProfile()
            return res.data as api.ModelsProfile
        } catch (err) {
            if (err instanceof HttpApiError && err.status === 401) {
                logout()
            }
            throw err
        }
    },
    enabled: !!token,
    staleTime: 1000 * 60 * 10,
  });

  const refresh = () => qc.invalidateQueries({ queryKey: keys.profile() });

  return {
    isAuthenticated: !!token,
    profile: query.data,
    isLoading: query.isLoading,
    isError: query.isError,
    error: query.error as Error | null,
    logout,
    refresh,
  }
}

export function useSaveProfile() {
    const queryClient = useQueryClient()
    const { mutate: saveProfile, isPending, error } = useMutation({
        mutationFn: async (data: api.ModelsProfileUpdateRequest) => {
            const res = await api.saveProfile(data)
            return res.data as api.ModelsProfile
        },
        onSuccess: () => {
            queryClient.invalidateQueries({
                queryKey: keys.profile(),
            })
        },
    })

    return { saveProfile, isPending, error }
}

const TOKEN_KEY = "token";

function getToken()    { return localStorage.getItem(TOKEN_KEY) }
function setToken(t: string) { localStorage.setItem(TOKEN_KEY, t) }
function clearToken() { localStorage.removeItem(TOKEN_KEY) }
