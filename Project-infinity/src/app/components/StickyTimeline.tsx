import { motion } from "motion/react";
import { useEffect, useState } from "react";

interface StickyTimelineProps {
  phases: Array<{ id: number; title: string }>;
}

export function StickyTimeline({ phases }: StickyTimelineProps) {
  const [activePhase, setActivePhase] = useState(1);

  useEffect(() => {
    const handleScroll = () => {
      const sections = phases.map((phase) =>
        document.getElementById(`phase-${phase.id}`)
      );

      const scrollPosition = window.scrollY + window.innerHeight / 2;

      for (let i = sections.length - 1; i >= 0; i--) {
        const section = sections[i];
        if (section && section.offsetTop <= scrollPosition) {
          setActivePhase(i + 1);
          break;
        }
      }
    };

    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  }, [phases]);

  return (
    <>
      {/* Desktop Timeline - Left Side */}
      <div className="hidden lg:block fixed left-8 top-1/2 -translate-y-1/2 z-50">
        <div className="relative">
          {/* Vertical Line */}
          <div className="absolute left-[7px] top-0 bottom-0 w-[2px] bg-slate-700" />

          {/* Active Progress Line */}
          <motion.div
            className="absolute left-[7px] top-0 w-[2px] bg-[#00F0FF]"
            initial={{ height: "0%" }}
            animate={{
              height: `${((activePhase - 1) / (phases.length - 1)) * 100}%`,
            }}
            transition={{ duration: 0.5 }}
          />

          {/* Phase Dots */}
          <div className="relative space-y-12">
            {phases.map((phase) => (
              <button
                key={phase.id}
                onClick={() => {
                  document.getElementById(`phase-${phase.id}`)?.scrollIntoView({
                    behavior: "smooth",
                  });
                }}
                className="flex items-center gap-4 group"
              >
                <motion.div
                  className={`w-4 h-4 rounded-full border-2 transition-all ${
                    activePhase >= phase.id
                      ? "bg-[#00F0FF] border-[#00F0FF] shadow-[0_0_15px_rgba(0,240,255,0.5)]"
                      : "bg-slate-800 border-slate-600"
                  }`}
                  whileHover={{ scale: 1.2 }}
                />
                <span
                  className={`text-sm font-mono tracking-wider transition-all ${
                    activePhase === phase.id
                      ? "text-[#00F0FF] opacity-100"
                      : "text-slate-500 opacity-0 group-hover:opacity-100"
                  }`}
                >
                  {phase.title}
                </span>
              </button>
            ))}
          </div>
        </div>
      </div>

      {/* Mobile Timeline - Top Horizontal */}
      <div className="lg:hidden fixed top-0 left-0 right-0 z-50 bg-[#050505]/95 backdrop-blur-sm border-b border-slate-800">
        <div className="px-4 py-3">
          <div className="flex items-center gap-2 overflow-x-auto scrollbar-hide">
            {phases.map((phase) => (
              <button
                key={phase.id}
                onClick={() => {
                  document.getElementById(`phase-${phase.id}`)?.scrollIntoView({
                    behavior: "smooth",
                  });
                }}
                className={`flex-shrink-0 px-3 py-1.5 rounded-full text-xs font-mono tracking-wider transition-all ${
                  activePhase === phase.id
                    ? "bg-[#00F0FF]/20 text-[#00F0FF] border border-[#00F0FF]"
                    : "bg-slate-800/50 text-slate-400 border border-slate-700"
                }`}
              >
                {String(phase.id).padStart(2, "0")}
              </button>
            ))}
          </div>

          {/* Progress Bar */}
          <div className="mt-2 h-[2px] bg-slate-800 rounded-full overflow-hidden">
            <motion.div
              className="h-full bg-[#00F0FF]"
              initial={{ width: "0%" }}
              animate={{ width: `${(activePhase / phases.length) * 100}%` }}
              transition={{ duration: 0.5 }}
            />
          </div>
        </div>
      </div>
    </>
  );
}
