import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query"
import api from "@/api"

const keys = {
    profile: () => ["profile"],
}

export function useSignup() {
    return useMutation({
        mutationFn: async (data: api.ModelsSignupRequest) => {
            const res = await api.signup(data)
            if (res.status > 299) {
                throw new Error(JSON.stringify(res.data))
            }
            return res
        },
        onSuccess: (res) => {
            const body = res.data as api.ModelsSignupResponse
            setToken(body.token)
        }
    })
}

export function useUser() {
  const qc = useQueryClient();
  const token = getToken();

  const query = useQuery({
    queryKey: keys.profile(),
    queryFn: async () => {
        const res = await api.getProfile()
        if (res.status > 299) {
            throw new Error(JSON.stringify(res.data))
        }
        return res.data as api.ModelsProfile
    },
    enabled: !!token,
    staleTime: 1000 * 60 * 10,
  });

  const logout = () => {
    clearToken();
    qc.removeQueries({ queryKey: keys.profile() });
  };

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
    return useMutation({
        mutationFn: (data: api.ModelsProfileUpdateRequest) =>
            api.saveProfile(data),
        onSuccess: () => {
            queryClient.invalidateQueries({
                queryKey: keys.profile(),
            })
        },
    })
}

const TOKEN_KEY = "token";

function getToken()    { return localStorage.getItem(TOKEN_KEY) }
function setToken(t: string) { localStorage.setItem(TOKEN_KEY, t) }
function clearToken() { localStorage.removeItem(TOKEN_KEY) }
