export interface Phase {
  id: number;
  identifier: string;
  title: string;
  thesis: string;
  context: string;
  powerShift: string;
  blueprint: string[];
  tension: string;
  transition: string;
  videoUrl: string;
  imageUrl: string;
}

export const phases: Phase[] = [
  {
    id: 1,
    identifier: "PHASE 01",
    title: "IGNITION",
    thesis: "We are not starting a project. We are activating a system.",
    context: "Ideas are scattered. Execution is undefined.",
    powerShift: "Chaos → Direction",
    blueprint: [
      "Define core vision & architecture",
      "Establish branding + design system",
      "Setup Vite + Tailwind project",
      "Create roadmap structure"
    ],
    tension: "Without this, everything becomes inconsistent.",
    transition: "With direction locked, we lay the foundation.",
    videoUrl: "/media/phase-01.mp4",
    imageUrl: "/media/phase-01.jpg"
  },
  {
    id: 2,
    identifier: "PHASE 02",
    title: "FOUNDATION",
    thesis: "We are not building features. We are building infrastructure.",
    context: "Scaling requires stability.",
    powerShift: "Concept → Structure",
    blueprint: [
      "Component system",
      "Auth system",
      "Database schema",
      "API architecture"
    ],
    tension: "Skipping this creates long-term technical debt.",
    transition: "With structure in place, we bring the system to life.",
    videoUrl: "/media/phase-02.mp4",
    imageUrl: "/media/phase-02.jpg"
  },
  {
    id: 3,
    identifier: "PHASE 03",
    title: "CORE SYSTEM",
    thesis: "We are not prototyping. We are operationalizing.",
    context: "System must now function.",
    powerShift: "Static → Functional",
    blueprint: [
      "Core features implementation",
      "Data flow integration",
      "Dashboard logic",
      "User interaction systems"
    ],
    tension: "Without this, it's just UI without value.",
    transition: "With functionality alive, we unify the experience.",
    videoUrl: "/media/phase-03.mp4",
    imageUrl: "/media/phase-03.jpg"
  },
  {
    id: 4,
    identifier: "PHASE 04",
    title: "CONVERGENCE",
    thesis: "We are not adding features. We are eliminating friction.",
    context: "Disconnected systems create poor UX.",
    powerShift: "Fragmented → Unified",
    blueprint: [
      "Feature integration",
      "UX optimization",
      "Performance tuning",
      "Interaction consistency"
    ],
    tension: "Growth becomes chaos without this.",
    transition: "With everything unified, we prepare to scale.",
    videoUrl: "/media/phase-04.mp4",
    imageUrl: "/media/phase-04.jpg"
  },
  {
    id: 5,
    identifier: "PHASE 05",
    title: "EXPANSION",
    thesis: "We are not growing slowly. We are scaling deliberately.",
    context: "System must handle real-world demand.",
    powerShift: "Local → Scalable",
    blueprint: [
      "Cloud deployment",
      "Load handling",
      "External integrations",
      "Performance optimization"
    ],
    tension: "Success will break the system if unprepared.",
    transition: "With scale unlocked, intelligence becomes the advantage.",
    videoUrl: "/media/phase-05.mp4",
    imageUrl: "/media/phase-05.jpg"
  },
  {
    id: 6,
    identifier: "PHASE 06",
    title: "INTELLIGENCE",
    thesis: "We are not reacting. We are predicting.",
    context: "Data must become actionable.",
    powerShift: "Reactive → Predictive",
    blueprint: [
      "Analytics engine",
      "AI automation",
      "Recommendation systems",
      "Behavior tracking"
    ],
    tension: "Without intelligence, competitors catch up.",
    transition: "With intelligence embedded, control becomes absolute.",
    videoUrl: "/media/phase-06.mp4",
    imageUrl: "/media/phase-06.jpg"
  },
  {
    id: 7,
    identifier: "PHASE 07",
    title: "DOMINION",
    thesis: "We are not participating. We are defining the space.",
    context: "Transition from product → ecosystem.",
    powerShift: "Product → Platform",
    blueprint: [
      "Open APIs",
      "Developer ecosystem",
      "Monetization systems",
      "Global scaling"
    ],
    tension: "Without this, the system remains replaceable.",
    transition: "This is no longer a roadmap. It's a system in motion.",
    videoUrl: "/media/phase-07.mp4",
    imageUrl: "/media/phase-07.jpg"
  }
];
