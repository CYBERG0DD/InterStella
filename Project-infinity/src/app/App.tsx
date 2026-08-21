import { useEffect } from "react";
import Lenis from "@studio-freight/lenis";
import { Hero } from "./components/Hero";
import { StickyTimeline } from "./components/StickyTimeline";
import { PhaseSection } from "./components/PhaseSection";
import { Footer } from "./components/Footer";
import { phases } from "./data/phases";

export default function App() {
  // Initialize Lenis Smooth Scroll
  useEffect(() => {
    const lenis = new Lenis({
      duration: 1.2,
      easing: (t) => Math.min(1, 1.001 - Math.pow(2, -10 * t)),
      smooth: true,
    });

    function raf(time: number) {
      lenis.raf(time);
      requestAnimationFrame(raf);
    }

    requestAnimationFrame(raf);

    return () => {
      lenis.destroy();
    };
  }, []);

  return (
    <div className="bg-[#050505] text-white min-h-screen">
      {/* Sticky Timeline Navigation */}
      <StickyTimeline
        phases={phases.map((p) => ({ id: p.id, title: p.title }))}
      />

      {/* Hero Section */}
      <Hero />

      {/* Phase Sections */}
      {phases.map((phase) => (
        <PhaseSection key={phase.id} phase={phase} />
      ))}

      {/* Footer */}
      <Footer />
    </div>
  );
}