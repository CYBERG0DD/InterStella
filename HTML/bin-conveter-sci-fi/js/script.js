
/* ── CANVAS GRID BG ── */
let drawGrid;
(function () {
  const canvas = document.getElementById('bg-canvas');
  const ctx = canvas.getContext('2d');
  let W, H;
  function resize() {
    W = canvas.width = window.innerWidth;
    H = canvas.height = window.innerHeight;
  }
  resize();
  window.addEventListener('resize', () => { resize(); drawGrid(); });

  drawGrid = function (isHex) {
    const c = isHex ? '191,95,255' : '0,229,255';
    ctx.clearRect(0, 0, W, H);

    const lines = 16;
    const horizon = H * 0.42;
    const vp = { x: W / 2, y: horizon };

    ctx.strokeStyle = `rgba(${c},0.04)`;
    ctx.lineWidth = 1;

    for (let i = 0; i <= lines; i++) {
      const t = i / lines;
      const y = horizon + (H - horizon) * (t * t);
      ctx.beginPath(); ctx.moveTo(0, y); ctx.lineTo(W, y); ctx.stroke();
    }

    const spread = W * 1.2;
    for (let i = 0; i <= lines; i++) {
      const x = (W / 2 - spread / 2) + (spread / lines) * i;
      ctx.beginPath(); ctx.moveTo(vp.x, vp.y); ctx.lineTo(x, H); ctx.stroke();
    }

    const grad = ctx.createLinearGradient(0, horizon - 20, 0, horizon + 20);
    grad.addColorStop(0, 'transparent');
    grad.addColorStop(0.5, `rgba(${c},0.07)`);
    grad.addColorStop(1, 'transparent');
    ctx.fillStyle = grad;
    ctx.fillRect(0, horizon - 20, W, 40);

    const rg = ctx.createRadialGradient(W / 2, 0, 0, W / 2, 0, H * .5);
    rg.addColorStop(0, `rgba(${c},0.03)`);
    rg.addColorStop(1, 'transparent');
    ctx.fillStyle = rg;
    ctx.fillRect(0, 0, W, H);
  };

  drawGrid(false);
})();

/* ── BINARY RAIN ── */
const rainContainer = document.getElementById('binary-rain');
const cols = Math.floor(window.innerWidth / 18);
for (let i = 0; i < cols; i++) {
  const col = document.createElement('div');
  col.className = 'rain-col';
  col.style.left = (i * 18) + 'px';
  col.style.animationDuration = (10 + Math.random() * 18) + 's';
  col.style.animationDelay = (-Math.random() * 24) + 's';
  col.style.fontSize = (9 + Math.random() * 4) + 'px';
  let bits = '';
  for (let j = 0; j < 60; j++) bits += (Math.random() > .5 ? '1' : '0');
  col.textContent = bits;
  rainContainer.appendChild(col);
}

/* ── STATE ── */
let mode = 'encode';
let outputText = '';
let animTimeout = null;

/* ── ATOM CANVAS ── */
(function () {
  const canvas = document.getElementById('atom-canvas');
  const ctx = canvas.getContext('2d');
  const W = canvas.width, H = canvas.height;
  const cx = W / 2, cy = H / 2;

  // 3 orbits: tiltX, tiltY, speed, color-source
  const orbits = [
    { tiltX: 20, tiltY: 0, tiltZ: 0, speed: 0.018, angle: 0 },
    { tiltX: 20, tiltY: 0, tiltZ: 60, speed: 0.013, angle: 2.09 },
    { tiltX: 20, tiltY: 0, tiltZ: 120, speed: 0.022, angle: 4.19 },
  ];

  const RX = 100, RY = 28; // orbit radii — wide ellipse tilted

  // project 3D point to 2D with perspective
  function project(x, y, z) {
    const fov = 320;
    const dist = fov + z;
    return {
      x: cx + (x * fov) / dist,
      y: cy + (y * fov) / dist,
      scale: fov / dist
    };
  }

  // rotate point around Z axis
  function rotZ(x, y, a) {
    return { x: x * Math.cos(a) - y * Math.sin(a), y: x * Math.sin(a) + y * Math.cos(a) };
  }
  // rotate point around X axis
  function rotX(y, z, a) {
    return { y: y * Math.cos(a) - z * Math.sin(a), z: y * Math.sin(a) + z * Math.cos(a) };
  }

  // get orbit 3D points for an ellipse rotated by tiltZ
  function getOrbitPoints(tiltZdeg, steps) {
    const tz = tiltZdeg * Math.PI / 180;
    const tilt = 68 * Math.PI / 180; // how much to tilt toward viewer
    const pts = [];
    for (let i = 0; i <= steps; i++) {
      const a = (i / steps) * Math.PI * 2;
      let x = RX * Math.cos(a);
      let y = RY * Math.sin(a);
      let z = 0;
      // rotate in plane by tiltZ
      const rz = rotZ(x, y, tz);
      x = rz.x; y = rz.y;
      // tilt toward viewer (X rotation)
      const rx = rotX(y, z, tilt);
      y = rx.y; z = rx.z;
      pts.push({ x, y, z });
    }
    return pts;
  }

  // electron position on orbit at angle
  function getElectronPos(tiltZdeg, angle) {
    const tz = tiltZdeg * Math.PI / 180;
    const tilt = 68 * Math.PI / 180;
    let x = RX * Math.cos(angle);
    let y = RY * Math.sin(angle);
    let z = 0;
    const rz = rotZ(x, y, tz);
    x = rz.x; y = rz.y;
    const rx = rotX(y, z, tilt);
    y = rx.y; z = rx.z;
    return { x, y, z };
  }

  let activeColor = '#00e5ff';
  let glowColor = 'rgba(0,229,255,0.5)';

  function hexToRgb(hex) {
    const r = parseInt(hex.slice(1, 3), 16);
    const g = parseInt(hex.slice(3, 5), 16);
    const b = parseInt(hex.slice(5, 7), 16);
    return { r, g, b };
  }

  function drawAtom() {
    ctx.clearRect(0, 0, W, H);
    const rgb = hexToRgb(activeColor);
    const col = `rgb(${rgb.r},${rgb.g},${rgb.b})`;
    const colA = (a) => `rgba(${rgb.r},${rgb.g},${rgb.b},${a})`;

    // draw each orbit as projected ellipse path
    orbits.forEach((orb, oi) => {
      const pts = getOrbitPoints(orb.tiltZ, 120);
      ctx.beginPath();
      pts.forEach((p, i) => {
        const proj = project(p.x, p.y, p.z);
        if (i === 0) ctx.moveTo(proj.x, proj.y);
        else ctx.lineTo(proj.x, proj.y);
      });
      ctx.closePath();
      ctx.strokeStyle = colA(0.35);
      ctx.lineWidth = 1;
      ctx.stroke();
    });

    // draw electrons on top
    orbits.forEach((orb) => {
      const ep = getElectronPos(orb.tiltZ, orb.angle);
      const proj = project(ep.x, ep.y, ep.z);
      const r = 5 * proj.scale;

      // glow
      const eg = ctx.createRadialGradient(proj.x, proj.y, 0, proj.x, proj.y, r * 4);
      eg.addColorStop(0, colA(0.9));
      eg.addColorStop(0.4, colA(0.4));
      eg.addColorStop(1, colA(0));
      ctx.beginPath();
      ctx.arc(proj.x, proj.y, r * 4, 0, Math.PI * 2);
      ctx.fillStyle = eg;
      ctx.fill();

      // electron core with sphere shading
      const esg = ctx.createRadialGradient(proj.x - r * 0.3, proj.y - r * 0.3, r * 0.1, proj.x, proj.y, r);
      esg.addColorStop(0, 'white');
      esg.addColorStop(0.3, col);
      esg.addColorStop(1, colA(0.6));
      ctx.beginPath();
      ctx.arc(proj.x, proj.y, Math.max(r, 3), 0, Math.PI * 2);
      ctx.fillStyle = esg;
      ctx.fill();
    });
  }

  function animate() {
    orbits.forEach(o => o.angle += o.speed);
    drawAtom();
    requestAnimationFrame(animate);
  }
  animate();

  // expose color updater
  window.setAtomColor = function (hex) {
    activeColor = hex;
    const rgb = hexToRgb(hex);
    glowColor = `rgba(${rgb.r},${rgb.g},${rgb.b},0.5)`;
    document.getElementById('atom-canvas').style.filter = `drop-shadow(0 0 18px ${hex})`;
  };
})();


function setMode(m) {
  mode = m;
  const isHex = m === 'hexencode' || m === 'hexdecode';

  ['encode', 'decode', 'hexencode', 'hexdecode'].forEach(id => {
    const el = document.getElementById('btn-' + id);
    el.classList.remove('active', 'hex-active');
  });
  document.getElementById('btn-' + m).classList.add(isHex ? 'hex-active' : 'active');

  // Switch ALL active CSS variables globally
  const root = document.documentElement;
  if (isHex) {
    root.style.setProperty('--active-color', '#bf5fff');
    root.style.setProperty('--active-dim', 'rgba(191,95,255,0.12)');
    root.style.setProperty('--active-glow', 'rgba(191,95,255,0.3)');
    root.style.setProperty('--active-border', 'rgba(191,95,255,0.25)');
    if (window.setAtomColor) window.setAtomColor('#bf5fff');
  } else {
    root.style.setProperty('--active-color', '#00e5ff');
    root.style.setProperty('--active-dim', 'rgba(0,229,255,0.12)');
    root.style.setProperty('--active-glow', 'rgba(0,229,255,0.3)');
    root.style.setProperty('--active-border', 'rgba(0,229,255,0.25)');
    if (window.setAtomColor) window.setAtomColor('#00e5ff');
  }

  // keep panel class for legacy fallback
  document.getElementById('main-panel').classList.toggle('hex-mode', isHex);

  // recolor rain
  const rainColor = isHex ? '#bf5fff' : '#00e5ff';
  document.querySelectorAll('.rain-col').forEach(col => col.style.color = rainColor);

  // redraw canvas grid with new color
  drawGrid(isHex);

  // recolor HUD dots
  document.querySelectorAll('.hud-dot').forEach(d => {
    d.style.background = rainColor;
    d.style.boxShadow = `0 0 6px ${rainColor}`;
  });

  // recolor logo ring
  const logoRing = document.querySelector('.logo-ring');
  if (logoRing) logoRing.style.borderColor = isHex ? 'rgba(191,95,255,0.2)' : 'rgba(0,229,255,0.15)';

  // recolor h1 glow
  const h1 = document.querySelector('h1');
  if (h1) h1.style.textShadow = isHex
    ? '0 0 30px rgba(191,95,255,.9),0 0 80px rgba(191,95,255,.4),0 0 160px rgba(191,95,255,.15)'
    : '0 0 30px rgba(0,229,255,.9),0 0 80px rgba(0,229,255,.4),0 0 160px rgba(0,229,255,.15)';
  if (h1) h1.style.color = rainColor;

  const titles = {
    encode: ['ENCODE // TEXT TO BINARY', 'INPUT TEXT', 'Type something to encode...', 'BASE-2'],
    decode: ['DECODE // BINARY TO TEXT', 'INPUT BINARY (8-BIT SPACE-SEP)', '01001000 01100101 01101100 01101111...', 'BASE-2'],
    hexencode: ['ENCODE // TEXT TO HEX', 'INPUT TEXT', 'Type something to convert to hex...', 'BASE-16'],
    hexdecode: ['DECODE // HEX TO TEXT', 'INPUT HEX (SPACE-SEPARATED)', '48 65 6c 6c 6f 20 57 6f 72 6c 64...', 'BASE-16'],
  };

  const [title, label, ph, base] = titles[m];
  document.getElementById('panel-title').textContent = title;
  document.getElementById('input-label-text').textContent = label;
  document.getElementById('input-area').placeholder = ph;
  document.getElementById('ph-right').textContent = base;

  clearAll();
}

/* ── VALIDATION ── */
function isValidText(t) { return t.length > 0; }
function isValidBinary(b) { return /^[01 ]+$/.test(b); }
function isValidHex(h) { return /^[0-9a-fA-F ]+$/.test(h.trim()); }

/* ── ENCODE / DECODE ── */
function encodeBinary(text) {
  return text.split('').map(c => c.charCodeAt(0).toString(2).padStart(8, '0')).join(' ');
}
function decodeBinary(binary) {
  return binary.trim().split(/\s+/).map(w => {
    const code = parseInt(w, 2);
    if (isNaN(code)) throw new Error('Invalid binary: ' + w);
    return String.fromCharCode(code);
  }).join('');
}
function encodeHex(text) {
  return text.split('').map(c => c.charCodeAt(0).toString(16).padStart(2, '0').toUpperCase()).join(' ');
}
function decodeHex(hex) {
  return hex.trim().split(/\s+/).map(b => {
    const code = parseInt(b, 16);
    if (isNaN(code)) throw new Error('Invalid hex: ' + b);
    return String.fromCharCode(code);
  }).join('');
}

/* ── INPUT HANDLER ── */
function handleInput() {
  const val = document.getElementById('input-area').value;
  const cc = document.getElementById('char-count');
  cc.textContent = val.length + ' CHARS';
  cc.classList.toggle('warn', val.length > 200);
  document.getElementById('input-area').classList.remove('error');
  document.getElementById('error-msg').textContent = '';
}

function handleKeydown(e) {
  if (e.key === 'Enter' && (e.ctrlKey || e.metaKey)) {
    e.preventDefault(); runTranslate();
  }
}

/* ── TRANSLATE ── */
function runTranslate() {
  const input = document.getElementById('input-area').value.trim();
  const errorEl = document.getElementById('error-msg');
  const inputEl = document.getElementById('input-area');

  inputEl.classList.remove('error');
  errorEl.textContent = '';

  if (!input) {
    errorEl.textContent = '⚠ INPUT CANNOT BE EMPTY';
    inputEl.classList.add('error'); shake(inputEl); return;
  }

  try {
    if (mode === 'encode') {
      outputText = encodeBinary(input);
      showOutput(outputText);
      updateStats(input.length, outputText.replace(/ /g, '').length, input.length, 'BIN ENC');
    } else if (mode === 'decode') {
      if (!isValidBinary(input)) {
        errorEl.textContent = '⚠ ONLY 0, 1 AND SPACES ALLOWED';
        inputEl.classList.add('error'); shake(inputEl); return;
      }
      outputText = decodeBinary(input);
      showOutput(outputText);
      updateStats(outputText.length, input.replace(/ /g, '').length, input.trim().split(/\s+/).length, 'BIN DEC');
    } else if (mode === 'hexencode') {
      outputText = encodeHex(input);
      showOutput(outputText);
      updateStats(input.length, outputText.replace(/ /g, '').length, input.length, 'HEX ENC');
    } else if (mode === 'hexdecode') {
      if (!isValidHex(input)) {
        errorEl.textContent = '⚠ ONLY HEX CHARS (0-9, A-F) AND SPACES';
        inputEl.classList.add('error'); shake(inputEl); return;
      }
      outputText = decodeHex(input);
      showOutput(outputText);
      updateStats(outputText.length, input.replace(/ /g, '').length, input.trim().split(/\s+/).length, 'HEX DEC');
    }
  } catch (e) {
    errorEl.textContent = '⚠ ' + e.message.toUpperCase();
    inputEl.classList.add('error');
  }
}

/* ── SHOW OUTPUT ── */
function showOutput(text) {
  const box = document.getElementById('output-box');
  box.classList.remove('empty');
  box.classList.add('has-result', 'glitch');
  setTimeout(() => box.classList.remove('glitch'), 280);

  if (animTimeout) clearTimeout(animTimeout);
  box.textContent = '';

  const delay = Math.min(18, 700 / text.length);
  let i = 0;
  function addChar() {
    if (i < text.length) {
      box.textContent += text[i++];
      animTimeout = setTimeout(addChar, delay);
    }
  }
  addChar();
  document.getElementById('copy-btn').style.display = 'block';
}

/* ── STATS ── */
function updateStats(chars, bits, bytes, modeLabel) {
  document.getElementById('stat-chars').textContent = chars;
  document.getElementById('stat-bits').textContent = bits;
  document.getElementById('stat-bytes').textContent = bytes;
  document.getElementById('stat-mode').textContent = modeLabel;
  document.getElementById('stats-bar').classList.add('visible');
}

/* ── COPY ── */
function copyOutput() {
  if (!outputText) return;
  const btn = document.getElementById('copy-btn');
  const doSuccess = () => {
    btn.textContent = 'COPIED!';
    btn.classList.add('copied');
    setTimeout(() => { btn.textContent = 'COPY'; btn.classList.remove('copied'); }, 1800);
  };
  if (navigator.clipboard && navigator.clipboard.writeText) {
    navigator.clipboard.writeText(outputText).then(doSuccess).catch(() => fallbackCopy(doSuccess));
  } else { fallbackCopy(doSuccess); }
}
function fallbackCopy(onSuccess) {
  const ta = document.createElement('textarea');
  ta.value = outputText;
  ta.style.cssText = 'position:fixed;top:-9999px;left:-9999px;opacity:0';
  document.body.appendChild(ta);
  ta.focus(); ta.select();
  try { document.execCommand('copy'); onSuccess(); }
  catch (e) { alert('Copy failed — please select and copy manually.'); }
  document.body.removeChild(ta);
}

/* ── CLEAR ── */
function clearAll() {
  document.getElementById('input-area').value = '';
  document.getElementById('input-area').classList.remove('error');
  document.getElementById('error-msg').textContent = '';
  document.getElementById('char-count').textContent = '0 CHARS';
  document.getElementById('char-count').classList.remove('warn');
  document.getElementById('output-box').textContent = 'Awaiting input...';
  document.getElementById('output-box').classList.add('empty');
  document.getElementById('output-box').classList.remove('has-result');
  document.getElementById('stats-bar').classList.remove('visible');
  outputText = '';
  if (animTimeout) clearTimeout(animTimeout);
}

/* ── SHAKE ── */
function shake(el) {
  el.style.animation = 'none';
  requestAnimationFrame(() => {
    el.style.animation = 'glitch .3s ease';
    setTimeout(() => el.style.animation = '', 300);
  });
}

/* ── SERVICE WORKER ── */
if ('serviceWorker' in navigator) {
  window.addEventListener('load', () => {
    navigator.serviceWorker.register('sw.js')
      .then(() => console.log('BASEX SW registered'))
      .catch(e => console.log('SW error:', e));
  });
}