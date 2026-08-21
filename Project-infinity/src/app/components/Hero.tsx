import { motion } from "motion/react";
import { ChevronDown } from "lucide-react";

export function Hero() {
  return (
    <section className="relative min-h-screen flex items-center justify-center overflow-hidden bg-[#050505]">
      {/* Animated Background */}
      <div className="absolute inset-0">
        <motion.div
          className="absolute inset-0 opacity-30"
          animate={{
            background: [
              "radial-gradient(circle at 20% 50%, #00F0FF 0%, transparent 50%)",
              "radial-gradient(circle at 80% 50%, #00F0FF 0%, transparent 50%)",
              "radial-gradient(circle at 50% 80%, #00F0FF 0%, transparent 50%)",
              "radial-gradient(circle at 20% 50%, #00F0FF 0%, transparent 50%)",
            ],
          }}
          transition={{
            duration: 20,
            repeat: Infinity,
            ease: "linear",
          }}
        />

        {/* Particle Grid */}
        <div className="absolute inset-0 opacity-20">
          <div className="absolute inset-0" style={{
            backgroundImage: `radial-gradient(circle, #00F0FF 1px, transparent 1px)`,
            backgroundSize: "50px 50px",
          }} />
        </div>
      </div>

      {/* Content */}
      <div className="relative z-10 text-center px-6">
        <motion.div
          initial={{ opacity: 0, y: 30 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 1, delay: 0.2 }}
        >
          <div className="inline-block mb-6 px-6 py-2 border border-[#00F0FF]/30 rounded-full">
            <p className="text-[#00F0FF] tracking-[0.3em] text-xs font-mono">
              INFINITE ROADMAP
            </p>
          </div>
        </motion.div>

        <motion.h1
          className="text-7xl md:text-9xl font-bold tracking-wider mb-8 text-white"
          initial={{ opacity: 0, y: 30 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 1, delay: 0.4 }}
        >
          ___
        </motion.h1>

        <motion.p
          className="text-xl md:text-2xl text-slate-300 max-w-3xl mx-auto tracking-wide leading-relaxed"
          initial={{ opacity: 0, y: 30 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 1, delay: 0.6 }}
        >
          NOT A ROADMAP. A SYSTEM IN MOTION.
          <br />
          <span className="text-[#00F0FF]">SEVEN PHASES. INFINITE IMPACT.</span>
        </motion.p>

        <motion.div
          className="mt-16"
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          transition={{ duration: 1, delay: 1 }}
        >
          <motion.div
            animate={{ y: [0, 10, 0] }}
            transition={{ duration: 2, repeat: Infinity }}
          >
            <ChevronDown className="w-8 h-8 text-[#00F0FF] mx-auto" />
          </motion.div>
        </motion.div>
      </div>

      {/* Bottom Gradient Fade */}
      <div className="absolute bottom-0 left-0 right-0 h-32 bg-gradient-to-t from-[#050505] to-transparent" />
    </section>
  );
}
