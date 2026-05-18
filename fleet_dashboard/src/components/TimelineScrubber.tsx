import { useFilterStore } from '../stores/useFilterStore';

type Props = {
  minYear?: number;
  maxYear?: number;
};

export function TimelineScrubber({ minYear = 2000, maxYear = 2030 }: Props) {
  const year = useFilterStore((state) => state.year);
  const setYear = useFilterStore((state) => state.setYear);

  return (
    <section className="panel">
      <div className="panel-header">
        <h2>Timeline Scrubber</h2>
        <strong>{year}</strong>
      </div>
      <input
        type="range"
        min={minYear}
        max={maxYear}
        value={year}
        onChange={(event) => setYear(Number(event.target.value))}
        className="slider"
      />
      <div className="timeline-labels">
        <span>{minYear}</span>
        <span>{Math.round((minYear + maxYear) / 2)}</span>
        <span>{maxYear}</span>
      </div>
    </section>
  );
}
