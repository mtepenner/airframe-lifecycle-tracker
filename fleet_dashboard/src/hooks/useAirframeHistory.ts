import { useQuery } from '@tanstack/react-query';
import { useFilterStore } from '../stores/useFilterStore';

type AirframeSnapshot = {
  tailNumber: string;
  model: string;
  operator: string;
  livery: string;
  seatingCapacity: number;
  validFrom: string;
  validTo: string;
};

const API_BASE = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';

export function useAirframeHistory() {
  const tailNumber = useFilterStore((state) => state.tailNumber);

  return useQuery({
    queryKey: ['airframe-history', tailNumber],
    queryFn: async (): Promise<AirframeSnapshot[]> => {
      const response = await fetch(`${API_BASE}/api/v1/fleet/${tailNumber}/history`);
      if (!response.ok) {
        throw new Error('Failed to fetch airframe history');
      }
      return response.json();
    },
    staleTime: 60_000,
    enabled: Boolean(tailNumber)
  });
}
