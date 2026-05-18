import { useQuery } from '@tanstack/react-query';
import { useFilterStore } from '../stores/useFilterStore';

type DowntimePoint = {
  date: string;
  hours: number;
  tailNumber: string;
  model: string;
};

type DowntimeResponse = {
  from: string;
  to: string;
  points: DowntimePoint[];
  totalHours: number;
};

const API_BASE = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';

function formatDate(date: Date) {
  return date.toISOString().slice(0, 10);
}

export function useDowntimeAnalytics() {
  const { year, model } = useFilterStore();

  return useQuery({
    queryKey: ['downtime-analytics', year, model],
    queryFn: async (): Promise<DowntimeResponse> => {
      const from = new Date(Date.UTC(year, 0, 1));
      const to = new Date(Date.UTC(year, 11, 31));
      const search = new URLSearchParams({
        from: formatDate(from),
        to: formatDate(to)
      });
      if (model) {
        search.set('model', model);
      }

      const response = await fetch(`${API_BASE}/api/v1/analytics/downtime?${search.toString()}`);
      if (!response.ok) {
        throw new Error('Failed to fetch downtime analytics');
      }
      return response.json();
    },
    staleTime: 60_000
  });
}
