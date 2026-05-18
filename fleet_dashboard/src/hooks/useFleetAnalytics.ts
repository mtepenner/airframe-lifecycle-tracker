import { useQuery } from '@tanstack/react-query';
import { useFilterStore } from '../stores/useFilterStore';

type FleetCubeRow = {
  model: string;
  operator: string;
  year: number;
  totalDowntimeHours: number;
  maintenanceEvents: number;
};

type FleetAnalyticsResponse = {
  at: string;
  rows: FleetCubeRow[];
  totalEvents: number;
};

const API_BASE = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';

export function useFleetAnalytics() {
  const { year, model, operator } = useFilterStore();

  return useQuery({
    queryKey: ['fleet-analytics', year, model, operator],
    queryFn: async (): Promise<FleetAnalyticsResponse> => {
      const search = new URLSearchParams();
      search.set('year', String(year));
      if (model) {
        search.set('model', model);
      }
      if (operator) {
        search.set('operator', operator);
      }

      const response = await fetch(`${API_BASE}/api/v1/fleet?${search.toString()}`);
      if (!response.ok) {
        throw new Error('Failed to fetch fleet analytics');
      }
      return response.json();
    },
    staleTime: 60_000
  });
}
