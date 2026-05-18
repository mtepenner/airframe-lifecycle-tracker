import { create } from 'zustand';

type FilterState = {
  year: number;
  model: string;
  operator: string;
  tailNumber: string;
  setYear: (year: number) => void;
  setModel: (model: string) => void;
  setOperator: (operator: string) => void;
  setTailNumber: (tailNumber: string) => void;
};

export const useFilterStore = create<FilterState>((set) => ({
  year: new Date().getUTCFullYear(),
  model: '',
  operator: '',
  tailNumber: 'N711MD',
  setYear: (year) => set({ year }),
  setModel: (model) => set({ model }),
  setOperator: (operator) => set({ operator }),
  setTailNumber: (tailNumber) => set({ tailNumber })
}));
