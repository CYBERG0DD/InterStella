import { motion } from "motion/react";
import { Check, Clock } from "lucide-react";
import type { Phase } from "../data/phases";

interface PhaseSectionProps {
  phase: Phase;
}

export function PhaseSection({ phase }: PhaseSectionProps) {
  return (
    <section
      id={`phase-${phase.id}`}
      className="min-h-screen flex items-center justify-center py-20 px-6 relative"
    >
      <div className="max-w-6xl w-full">
        <div className="grid lg:grid-cols-2 gap-12 items-center">
          {/* Content */}
          <motion.div
            initial={{ opacity: 0, x: -50 }}
            whileInView={{ opacity: 1, x: 0 }}
            transition={{ duration: 0.8 }}
            viewport={{ once: true }}
          >
            {/* Phase Identifier */}
            <div className="inline-block mb-4 px-4 py-1 bg-[#00F0FF]/10 border border-[#00F0FF]/30 rounded-full">
              <p className="text-[#00F0FF] tracking-[0.3em] text-xs font-mono">
                {phase.identifier}
              </p>
            </div>

            {/* Phase Title */}
            <h2 className="text-5xl md:text-6xl font-bold tracking-wider mb-6 text-white">
              {phase.title}
            </h2>

            {/* Thesis */}
            <blockquote className="border-l-4 border-[#00F0FF] pl-6 mb-8">
              <p className="text-xl text-slate-300 italic leading-relaxed">
                "{phase.thesis}"
              </p>
            </blockquote>

            {/* Context */}
            <div className="mb-6">
              <h3 className="text-sm font-mono tracking-wider text-[#00F0FF] mb-2 uppercase">
                Context
              </h3>
              <p className="text-slate-400 leading-relaxed">{phase.context}</p>
            </div>

            {/* Power Shift */}
            <div className="mb-8 p-4 bg-gradient-to-r from-slate-800/50 to-transparent rounded-lg border-l-4 border-[#00F0FF]">
              <h3 className="text-sm font-mono tracking-wider text-[#00F0FF] mb-2 uppercase">
                Power Shift
              </h3>
              <p className="text-2xl font-bold text-white tracking-wide">
                {phase.powerShift}
              </p>
            </div>

            {/* Execution Blueprint */}
            <div className="mb-8 p-6 bg-slate-900/30 backdrop-blur-sm rounded-xl border border-slate-700/50 shadow-xl">
              <h3 className="text-sm font-mono tracking-wider text-[#00F0FF] mb-4 uppercase">
                Execution Blueprint
              </h3>
              <ul className="space-y-3">
                {phase.blueprint.map((item, index) => (
                  <motion.li
                    key={index}
                    className="flex items-start gap-3"
                    initial={{ opacity: 0, x: -20 }}
                    whileInView={{ opacity: 1, x: 0 }}
                    transition={{ duration: 0.5, delay: index * 0.1 }}
                    viewport={{ once: true }}
                  >
                    {/* Completed items (first 2) vs In Progress */}
                    {index < 2 ? (
                      <Check className="w-5 h-5 text-green-400 flex-shrink-0 mt-0.5" />
                    ) : (
                      <motion.div
                        animate={{ opacity: [0.5, 1, 0.5] }}
                        transition={{ duration: 2, repeat: Infinity }}
                      >
                        <Clock className="w-5 h-5 text-yellow-400 flex-shrink-0 mt-0.5" />
                      </motion.div>
                    )}
                    <span className="text-slate-300 leading-relaxed">
                      {item}
                    </span>
                  </motion.li>
                ))}
              </ul>
            </div>

            {/* Tension */}
            <div className="mb-6 p-4 bg-red-950/20 border border-red-900/30 rounded-lg">
              <h3 className="text-sm font-mono tracking-wider text-red-400 mb-2 uppercase">
                Tension
              </h3>
              <p className="text-slate-300 leading-relaxed">{phase.tension}</p>
            </div>

            {/* Transition */}
            <div className="pt-6 border-t border-slate-700">
              <p className="text-slate-400 italic leading-relaxed">
                {phase.transition}
              </p>
            </div>
          </motion.div>

          {/* Media */}
          <motion.div
            initial={{ opacity: 0, x: 50 }}
            whileInView={{ opacity: 1, x: 0 }}
            transition={{ duration: 0.8 }}
            viewport={{ once: true }}
            className="space-y-6"
          >
            {/* Video Placeholder */}
            <div className="relative aspect-video rounded-2xl overflow-hidden bg-slate-900 border border-slate-700 shadow-2xl group">
              <div className="absolute inset-0 flex items-center justify-center">
                <div className="text-center">
                  <div className="w-16 h-16 mx-auto mb-4 rounded-full bg-[#00F0FF]/20 flex items-center justify-center">
                    <div className="w-0 h-0 border-l-[12px] border-l-[#00F0FF] border-y-[8px] border-y-transparent ml-1" />
                  </div>
                  <p className="text-slate-500 font-mono text-sm">
                    VIDEO: {phase.identifier}
                  </p>
                  <p className="text-slate-600 text-xs mt-1">
                    {phase.videoUrl}
                  </p>
                </div>
              </div>
              {/* Uncomment when video URLs are available */}
              {/* <video
                src={phase.videoUrl}
                autoPlay
                muted
                loop
                playsInline
                className="w-full h-full object-cover"
              /> */}

              {/* Glow Effect */}
              <div className="absolute inset-0 opacity-0 group-hover:opacity-100 transition-opacity duration-300">
                <div className="absolute inset-0 bg-gradient-to-t from-[#00F0FF]/20 to-transparent" />
              </div>
            </div>

            {/* Image Placeholder */}
            <div className="relative aspect-[4/3] rounded-2xl overflow-hidden bg-slate-900 border border-slate-700 shadow-2xl group">
              <div className="absolute inset-0 flex items-center justify-center">
                <div className="text-center">
                  <div className="w-16 h-16 mx-auto mb-4 rounded-lg bg-[#00F0FF]/20 flex items-center justify-center">
                    <div className="w-10 h-10 border-2 border-[#00F0FF] rounded" />
                  </div>
                  <p className="text-slate-500 font-mono text-sm">
                    IMAGE: {phase.identifier}
                  </p>
                  <p className="text-slate-600 text-xs mt-1">
                    {phase.imageUrl}
                  </p>
                </div>
              </div>
              {/* Uncomment when image URLs are available */}
              {/* <img
                src={phase.imageUrl}
                alt={`${phase.title} visual`}
                className="w-full h-full object-cover"
              /> */}

              {/* Glow Effect */}
              <div className="absolute inset-0 opacity-0 group-hover:opacity-100 transition-opacity duration-300">
                <div className="absolute inset-0 bg-gradient-to-t from-[#00F0FF]/20 to-transparent" />
              </div>
            </div>
          </motion.div>
        </div>
      </div>
    </section>
  );
}
