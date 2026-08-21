import { motion } from "motion/react";
import { ArrowRight } from "lucide-react";

export function Footer() {
  return (
    <footer className="relative min-h-screen flex items-center justify-center bg-gradient-to-b from-[#050505] to-slate-950 overflow-hidden">
      {/* Background Effects */}
      <div className="absolute inset-0">
        <motion.div
          className="absolute inset-0 opacity-20"
          animate={{
            background: [
              "radial-gradient(circle at 50% 50%, #00F0FF 0%, transparent 70%)",
              "radial-gradient(circle at 50% 50%, #00F0FF 0%, transparent 90%)",
              "radial-gradient(circle at 50% 50%, #00F0FF 0%, transparent 70%)",
            ],
          }}
          transition={{
            duration: 8,
            repeat: Infinity,
            ease: "easeInOut",
          }}
        />
      </div>

      <div className="relative z-10 text-center px-6 max-w-5xl">
        <motion.div
          initial={{ opacity: 0, y: 30 }}
          whileInView={{ opacity: 1, y: 0 }}
          transition={{ duration: 1 }}
          viewport={{ once: true }}
        >
          <div className="inline-block mb-6 px-6 py-2 border border-[#00F0FF]/30 rounded-full">
            <p className="text-[#00F0FF] tracking-[0.3em] text-xs font-mono">
              LEGACY SECTION
            </p>
          </div>

          <h2 className="text-6xl md:text-8xl font-bold tracking-wider mb-8 text-white">
            THE VISION
          </h2>

          <div className="space-y-8 mb-12">
            <div className="max-w-3xl mx-auto">
              <h3 className="text-xl text-[#00F0FF] font-mono tracking-wider mb-4 uppercase">
                Mission
              </h3>
              <p className="text-2xl text-slate-300 leading-relaxed">
                To build systems that don't just solve problems—
                <span className="text-white font-bold"> they define industries.</span>
              </p>
            </div>

            <div className="max-w-3xl mx-auto">
              <h3 className="text-xl text-[#00F0FF] font-mono tracking-wider mb-4 uppercase">
                Vision
              </h3>
              <p className="text-2xl text-slate-300 leading-relaxed">
                A world where technology moves from reactive to predictive,
                from fragmented to unified,
                <span className="text-white font-bold"> from product to platform.</span>
              </p>
            </div>
          </div>

          {/* CTA */}
          <motion.div
            initial={{ opacity: 0, scale: 0.9 }}
            whileInView={{ opacity: 1, scale: 1 }}
            transition={{ duration: 0.8, delay: 0.3 }}
            viewport={{ once: true }}
          >
            <motion.button
              className="group relative px-12 py-5 bg-[#00F0FF] text-black font-bold text-lg tracking-wider rounded-full overflow-hidden transition-all hover:shadow-[0_0_40px_rgba(0,240,255,0.6)]"
              whileHover={{ scale: 1.05 }}
              whileTap={{ scale: 0.95 }}
            >
              <span className="relative z-10 flex items-center gap-3">
                JOIN THE VISION
                <ArrowRight className="w-6 h-6 group-hover:translate-x-1 transition-transform" />
              </span>

              {/* Animated Background */}
              <motion.div
                className="absolute inset-0 bg-gradient-to-r from-[#00F0FF] via-cyan-300 to-[#00F0FF]"
                animate={{
                  x: ["-100%", "100%"],
                }}
                transition={{
                  duration: 3,
                  repeat: Infinity,
                  ease: "linear",
                }}
              />
            </motion.button>
          </motion.div>

          {/* Bottom Text */}
          <motion.div
            className="mt-20 pt-12 border-t border-slate-800"
            initial={{ opacity: 0 }}
            whileInView={{ opacity: 1 }}
            transition={{ duration: 1, delay: 0.5 }}
            viewport={{ once: true }}
          >
            <p className="text-slate-500 text-sm font-mono tracking-wider">
              THIS IS NO LONGER A ROADMAP.
            </p>
            <p className="text-[#00F0FF] text-xl font-bold tracking-wider mt-2">
              IT'S A SYSTEM IN MOTION.
            </p>
          </motion.div>
        </motion.div>
      </div>
    </footer>
  );
}
