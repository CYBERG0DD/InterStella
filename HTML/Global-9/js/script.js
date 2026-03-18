


/* ───── CANVAS HOLOGRAPHIC GRID ───── */
(function () {
  const cv = document.getElementById('bg-canvas');
  const cx = cv.getContext('2d');
  let W, H;
  function resize() { W = cv.width = innerWidth; H = cv.height = innerHeight }
  addEventListener('resize', resize); resize();

  // hex grid
  function drawHex(x, y, r, a) {
    cx.beginPath();
    for (let i = 0; i < 6; i++) {
      const angle = Math.PI / 180 * (60 * i - 30);
      i === 0 ? cx.moveTo(x + r * Math.cos(angle), y + r * Math.sin(angle)) : cx.lineTo(x + r * Math.cos(angle), y + r * Math.sin(angle));
    }
    cx.closePath();
    cx.strokeStyle = `rgba(0,180,255,${a})`; cx.lineWidth = .5; cx.stroke();
  }

  const particles = Array.from({ length: 50 }, () => ({
    x: Math.random() * 2000, y: Math.random() * 1200,
    vx: (Math.random() - .5) * .2, vy: (Math.random() - .5) * .2,
    size: Math.random() * 1.2 + .4, alpha: Math.random() * .3 + .05,
  }));

  // data streams
  const streams = Array.from({ length: 28 }, () => ({
    x: Math.random() * 2000, y: Math.random() * 1200,
    len: Math.floor(Math.random() * 16 + 8),
    speed: Math.random() * 1.8 + 0.8,
    chars: [], alpha: Math.random() * .75 + 0.35,
  }));
  streams.forEach(s => {
    s.chars = Array.from({ length: s.len }, () => String.fromCharCode(0x30A0 + Math.random() * 96));
  });

  let t = 0;
  function draw() {
    cx.clearRect(0, 0, W, H);
    // gradient bg
    const g = cx.createRadialGradient(W / 2, H / 2, 0, W / 2, H / 2, Math.max(W, H) * .7);
    g.addColorStop(0, 'rgba(0,8,24,1)'); g.addColorStop(1, 'rgba(0,0,6,1)');
    cx.fillStyle = g; cx.fillRect(0, 0, W, H);

    // hex grid
    const hr = 32, hw = hr * Math.sqrt(3);
    for (let row = -1; row < H / hr / 1.5 + 1; row++) {
      for (let col = -1; col < W / hw + 1; col++) {
        const xo = col * hw + (row % 2) * hw / 2;
        const yo = row * hr * 1.5;
        const d = Math.hypot(xo - W / 2, yo - H / 2);
        const a = Math.max(0, .025 * (1 - d / Math.max(W, H)));
        drawHex(xo, yo, hr * .9, a);
      }
    }

    // streaming data columns
    t += .016;
    streams.forEach((s, i) => {
      s.y += s.speed;
      if (s.y > H + s.len * 14) s.y = -20;
      if (Math.random() < .05) s.chars[Math.floor(Math.random() * s.len)] = String.fromCharCode(0x30A0 + Math.random() * 96);
      s.chars.forEach((c, j) => {
        const frac = j / s.len;
        const a = frac < .3 ? s.alpha * frac / .3 : s.alpha * (1 - frac) * .85;
        cx.font = `${12 + j * .6}px Share Tech Mono`;
        cx.fillStyle = `rgba(0,255,${120 + j * 18},${a * (j === s.len - 1 ? 3 : 1.4)})`;
        cx.fillText(c, s.x, s.y + j * 16);
      });
    });

    // particles
    particles.forEach(p => {
      p.x += p.vx; p.y += p.vy;
      if (p.x < 0) p.x = W; if (p.x > W) p.x = 0;
      if (p.y < 0) p.y = H; if (p.y > H) p.y = 0;
      cx.beginPath(); cx.arc(p.x, p.y, p.size, 0, Math.PI * 2);
      cx.fillStyle = `rgba(0,200,255,${p.alpha})`; cx.fill();
    });

    // connections
    particles.forEach((a, i) => {
      particles.slice(i + 1, i + 5).forEach(b => {
        const d = Math.hypot(a.x - b.x, a.y - b.y);
        if (d < 120) {
          cx.beginPath(); cx.moveTo(a.x, a.y); cx.lineTo(b.x, b.y);
          cx.strokeStyle = `rgba(0,180,255,${.06 * (1 - d / 120)})`; cx.lineWidth = .4; cx.stroke();
        }
      });
    });

    requestAnimationFrame(draw);
  }
  draw();
})();

/* ───── BOOT SEQUENCE ───── */
const bootLines = [
  { t: 'dim', m: '> GLOBAL "9" CORPORATION // NEXUS NODE 7' },
  { t: 'dim', m: '> BUILD 9.0.1 // KERNEL INIT' },
  { t: '', m: '> LOADING SECURITY MODULES...' },
  { t: 'ok', m: '  [OK] BIOMETRIC ENGINE v4.2 LOADED' },
  { t: 'ok', m: '  [OK] ENCRYPTION MATRIX ACTIVE' },
  { t: 'ok', m: '  [OK] IDENTITY PROTOCOL ONLINE' },
  { t: 'warn', m: '  [!!] EXTERNAL ACCESS DETECTED — VERIFYING...' },
  { t: 'ok', m: '  [OK] CLEARANCE LEVEL 3 GRANTED' },
  { t: '', m: '> LAUNCHING REGISTRATION INTERFACE...' },
];

(function boot() {
  const bl = document.getElementById('boot-lines');
  let i = 0;
  function showNext() {
    if (i >= bootLines.length) {
      setTimeout(() => {
        const bs = document.getElementById('boot-screen');
        bs.style.transition = 'opacity .6s'; bs.style.opacity = '0';
        setTimeout(() => { bs.style.display = 'none'; const mw = document.getElementById('main-wrap'); mw.classList.add('visible'); }, 600);
      }, 500);
      return;
    }
    const line = bootLines[i++];
    const div = document.createElement('div');
    div.className = 'boot-line ' + (line.t || '');
    bl.appendChild(div);
    let j = 0; const txt = line.m;
    const typ = setInterval(() => {
      div.textContent = txt.slice(0, ++j);
      if (j >= txt.length) {
        clearInterval(typ);
        setTimeout(showNext, i === bootLines.length ? 800 : 120);
      }
    }, 18);
    bl.scrollTop = bl.scrollHeight;
  }
  showNext();
})();

/* ───── STATE ───── */
let step = 1;
let prof = 'cyber security specialist';
let clearanceLevel = '';
let workEnv = 'FIELD';
let division = 'GHOST PROTOCOL';
let oathAccepted = false;

/* ───── NAVIGATION ───── */
function goStep(n) {
  if (n > step && !validateStep(step)) return;
  document.getElementById('step-' + step).classList.remove('active');
  ['1', '2', '3', '4', '5'].forEach(i => {
    const t = document.getElementById('tab-' + i);
    const ni = parseInt(i);
    t.className = 'step-tab' + (ni === n ? ' active' : ni < n ? ' done' : '');
  });
  step = n;
  document.getElementById('step-' + n).classList.add('active');
  document.getElementById('hud-step').textContent = ('0' + n).slice(-2);
}

/* ───── VALIDATION ───── */
function validateStep(s) {
  if (s === 1) {
    const fn = document.getElementById('firstName').value.trim();
    const ln = document.getElementById('lastName').value.trim();
    if (!fn) { errField(document.getElementById('firstName'), 'h-fn', 'FIRST NAME REQUIRED'); return false }
    if (!ln) { errField(document.getElementById('lastName'), 'h-ln', 'LAST NAME REQUIRED'); return false }
    return true;
  }
  if (s === 2) {
    const u = document.getElementById('username').value.trim();
    const p = document.getElementById('phone').value.trim();
    const y = parseInt(document.getElementById('birthYear').value);
    if (u.length < 5) { errField(document.getElementById('username'), 'h-user', 'MIN 5 CHARACTERS'); return false }
    if (p.length !== 11 || isNaN(p)) { errField(document.getElementById('phone'), 'h-phone', 'MUST BE 11 DIGITS'); return false }
    const age = 2026 - y;
    if (isNaN(age) || age < 18 || age > 45) { errField(document.getElementById('birthYear'), 'h-year', `AGE ${age} // RANGE: 18–45`); return false }
    return true;
  }
  if (s === 4) {
    if (!clearanceLevel) {
      const h = document.getElementById('h-clearance'); h.textContent = '⚠ SELECT A CLEARANCE LEVEL'; h.className = 'field-msg err'; return false;
    }
    const r = document.getElementById('region').value;
    if (!r) {
      const h = document.getElementById('h-region'); h.textContent = '⚠ SELECT YOUR REGION'; h.className = 'field-msg err'; return false;
    }
    return true;
  }
  if (s === 5) {
    const nk = document.getElementById('neuralKey').value.trim();
    if (nk.length < 8) { errField(document.getElementById('neuralKey'), 'h-neural', 'NEURAL KEY MUST BE MIN 8 CHARS'); return false; }
    if (!oathAccepted) {
      const ob = document.getElementById('oath-box');
      ob.style.borderColor = 'rgba(255,56,96,0.5)';
      ob.style.boxShadow = '0 0 12px rgba(255,56,96,0.15)';
      setTimeout(() => { ob.style.borderColor = 'rgba(255,204,0,0.15)'; ob.style.boxShadow = 'none' }, 1500);
      return false;
    }
    return true;
  }
  return true;
}

function errField(inp, hid, msg) {
  inp.classList.add('invalid'); inp.classList.remove('valid');
  const h = document.getElementById(hid); h.textContent = '⚠ ' + msg; h.className = 'field-msg err';
  inp.classList.add('shake'); setTimeout(() => inp.classList.remove('shake'), 400);
  inp.focus();
}
function okField(inp, hid, indId, msg) {
  inp.classList.remove('invalid'); inp.classList.add('valid');
  const h = document.getElementById(hid); h.textContent = '✓ ' + msg; h.className = 'field-msg ok';
  if (indId) { const ind = document.getElementById(indId); ind.textContent = '◈'; ind.style.cssText = 'opacity:1;color:var(--c3)' }
}

function vName(inp, hid) {
  const v = inp.value.trim();
  inp.classList.remove('valid', 'invalid');
  if (v.length >= 2) okField(inp, hid, '', 'IDENTITY DATA LOGGED');
  else if (v.length) { inp.classList.add('invalid'); const h = document.getElementById(hid); h.textContent = '// TOO SHORT'; h.className = 'field-msg err' }
  else { document.getElementById(hid).textContent = ''; inp.classList.remove('valid', 'invalid') }
}

function vUser(inp) {
  const v = inp.value.trim();
  const bars = [1, 2, 3, 4, 5].map(i => document.getElementById('sb' + i));
  bars.forEach(b => b.className = 'sb');
  let str = 0;
  if (v.length >= 5) str = 1;
  if (v.length >= 7) str = 2;
  if (/[0-9_\-]/.test(v)) str = Math.max(str, 2);
  if (v.length >= 9 && /[0-9]/.test(v)) str = 3;
  const cls = ['', 'w', 'w', 'm', 'm', 's'];
  const labels = ['', '', 'WEAK SIGNAL', 'MEDIUM SIGNAL', 'MEDIUM SIGNAL', 'STRONG SIGNAL'];
  for (let i = 0; i <= str && i < 5; i++)bars[i].className = 'sb ' + (cls[str + 1] || 's');
  document.getElementById('str-label').textContent = '// ' + (labels[str + 1] || 'STRONG SIGNAL');
  if (v.length < 5) {
    inp.classList.add('invalid'); inp.classList.remove('valid');
    const h = document.getElementById('h-user'); h.textContent = `// ${5 - v.length} MORE CHARS NEEDED`; h.className = 'field-msg err';
    document.getElementById('ind-user').style.opacity = 0;
  } else {
    okField(inp, 'h-user', 'ind-user', `HANDLE REGISTERED`);
  }
}

function vPhone(inp) {
  const v = inp.value.replace(/\D/g, ''); inp.value = v;
  if (v.length === 11) okField(inp, 'h-phone', 'ind-phone', 'COMM LINK VERIFIED');
  else {
    inp.classList.remove('valid'); inp.classList.add('invalid');
    const h = document.getElementById('h-phone'); h.textContent = '// ' + v.length + '/11 DIGITS'; h.className = 'field-msg' + (v.length ? 'err' : '') + ' err';
    document.getElementById('ind-phone').style.opacity = 0;
  }
}

function vYear(inp) {
  const age = 2026 - parseInt(inp.value);
  if (age >= 18 && age <= 45) okField(inp, 'h-year', 'ind-year', `AGE ${age} // ELIGIBLE`);
  else if (inp.value.length === 4) {
    inp.classList.remove('valid'); inp.classList.add('invalid');
    const h = document.getElementById('h-year'); h.textContent = `// AGE ${age} — NOT ELIGIBLE`; h.className = 'field-msg err';
    document.getElementById('ind-year').style.opacity = 0;
  }
}

/* ───── ROLES DATA ───── */
const ROLES = [
  { id: 'cyber', icon: '🛡', label: 'CYBER\nSECURITY', val: 'cyber security specialist' },
  { id: 'sw', icon: '⚙', label: 'SOFTWARE\nENGINEER', val: 'software engineer' },
  { id: 'it', icon: '🔌', label: 'IT\nTECHNICIAN', val: 'it technician' },
  { id: 'net', icon: '🌐', label: 'NETWORK\nARCHITECT', val: 'network architect' },
  { id: 'ai', icon: '🤖', label: 'AI\nSPECIALIST', val: 'ai specialist' },
  { id: 'data', icon: '📊', label: 'DATA\nSCIENTIST', val: 'data scientist' },
  { id: 'cloud', icon: '☁', label: 'CLOUD\nENGINEER', val: 'cloud engineer' },
  { id: 'pentest', icon: '🔓', label: 'PENETRATION\nTESTER', val: 'penetration tester' },
  { id: 'devops', icon: '⚡', label: 'DEVOPS\nENGINEER', val: 'devops engineer' },
];

(function renderRoles() {
  const grid = document.getElementById('prof-grid');
  grid.innerHTML = ROLES.map((r, i) => `
    <div class="prof-card${i === 0 ? ' sel' : ''}" id="pc-${r.id}" onclick="selProf('${r.val}','pc-${r.id}')">
      <div class="prof-scan"></div>
      <div class="prof-icon">${r.icon}</div>
      <div class="prof-label">${r.label.replace('\n', '<br>')}</div>
      <div class="prof-status" id="ps-${r.id}">${i === 0 ? '// SELECTED' : '// AVAILABLE'}</div>
    </div>
  `).join('');
})();

/* ───── PROFESSION ───── */
function selProf(val, id) {
  prof = val;
  ROLES.forEach(r => {
    const card = document.getElementById('pc-' + r.id);
    const stat = document.getElementById('ps-' + r.id);
    if (card) { card.classList.remove('sel'); if (stat) stat.textContent = '// AVAILABLE'; }
  });
  const sel = document.getElementById(id);
  const selSt = document.getElementById('ps-' + id.replace('pc-', ''));
  if (sel) sel.classList.add('sel');
  if (selSt) selSt.textContent = '// SELECTED';
}

/* ───── CLEARANCE DATA ───── */
const CLEARANCES = [
  { level: 'LV1', name: 'RECRUIT', desc: 'Entry access' },
  { level: 'LV2', name: 'OPERATIVE', desc: 'Field missions' },
  { level: 'LV3', name: 'SPECIALIST', desc: 'Classified ops' },
  { level: 'LV4', name: 'COMMANDER', desc: 'High command' },
  { level: 'LV5', name: 'DIRECTOR', desc: 'Nexus access' },
  { level: 'LV9', name: 'PHANTOM', desc: 'Unrestricted' },
];
(function renderClearance() {
  const g = document.getElementById('clearance-grid');
  g.innerHTML = CLEARANCES.map(c => `
    <div class="cl-card" id="cl-${c.level}" onclick="selClearance('${c.level}','cl-${c.level}')">
      <div class="cl-level">${c.level}</div>
      <div class="cl-name">${c.name}</div>
      <div class="cl-desc">${c.desc}</div>
    </div>
  `).join('');
})();
function selClearance(val, id) {
  clearanceLevel = val;
  CLEARANCES.forEach(c => { const el = document.getElementById('cl-' + c.level); if (el) el.classList.remove('sel') });
  document.getElementById(id).classList.add('sel');
  const h = document.getElementById('h-clearance'); h.textContent = '✓ CLEARANCE ' + val + ' DECLARED'; h.className = 'field-msg ok';
}

/* ───── ENV DATA ───── */
const ENVS = [
  { id: 'field', icon: '🌐', name: 'FIELD' },
  { id: 'remote', icon: '📡', name: 'REMOTE' },
  { id: 'facility', icon: '🏢', name: 'FACILITY' },
];
(function renderEnv() {
  const g = document.getElementById('env-grid');
  g.innerHTML = ENVS.map((e, i) => `
    <div class="env-card${i === 0 ? ' sel' : ''}" id="ev-${e.id}" onclick="selEnv('${e.name}','ev-${e.id}')">
      <div class="env-icon">${e.icon}</div>
      <div class="env-name">${e.name}</div>
    </div>
  `).join('');
})();
function selEnv(val, id) {
  workEnv = val;
  ENVS.forEach(e => { const el = document.getElementById('ev-' + e.id); if (el) el.classList.remove('sel') });
  document.getElementById(id).classList.add('sel');
}

/* ───── DIVISION DATA ───── */
const DIVISIONS = [
  { icon: '👻', name: 'GHOST PROTOCOL', desc: 'Stealth & covert ops' },
  { icon: '🔬', name: 'NEXUS LAB', desc: 'Research & development' },
  { icon: '🛰', name: 'ORBIT WATCH', desc: 'Surveillance & intel' },
  { icon: '⚔', name: 'IRON VANGUARD', desc: 'Tactical response' },
];
(function renderDivisions() {
  const g = document.getElementById('division-grid');
  g.innerHTML = DIVISIONS.map((d, i) => `
    <div class="div-card${i === 0 ? ' sel' : ''}" id="dv-${i}" onclick="selDiv('${d.name}','dv-${i}')">
      <div class="div-icon">${d.icon}</div>
      <div class="div-info">
        <div class="div-name">${d.name}</div>
        <div class="div-desc">${d.desc}</div>
      </div>
    </div>
  `).join('');
})();
function selDiv(val, id) {
  division = val;
  DIVISIONS.forEach((_, i) => { const el = document.getElementById('dv-' + i); if (el) el.classList.remove('sel') });
  document.getElementById(id).classList.add('sel');
}

/* ───── OATH ───── */
function toggleOath() {
  oathAccepted = !oathAccepted;
  const cb = document.getElementById('oath-cb');
  cb.textContent = oathAccepted ? '◼' : '◻';
  cb.className = 'oath-checkbox' + (oathAccepted ? ' checked' : '');
  document.getElementById('oath-label').textContent = oathAccepted ? 'OATH ACCEPTED // CREDENTIALS BOUND' : 'I ACCEPT THIS OATH AND BIND MY CREDENTIALS';
}

/* ───── NEURAL KEY ───── */
function vNeural(inp) {
  const v = inp.value.trim();
  if (v.length >= 8) okField(inp, 'h-neural', 'ind-neural', 'NEURAL SIGNATURE LOCKED');
  else if (v.length) {
    inp.classList.add('invalid'); inp.classList.remove('valid');
    const h = document.getElementById('h-neural'); h.textContent = `// ${8 - v.length} MORE CHARS NEEDED`; h.className = 'field-msg err';
    document.getElementById('ind-neural').style.opacity = 0;
  }
}

/* ───── SUBMIT ───── */
const loadSequence = [
  { title: 'IDENTITY SCAN', sub: 'Parsing biometric data...', log: '> READING IDENTITY MATRIX...' },
  { title: 'CREDENTIAL CHECK', sub: 'Verifying against G9 network...', log: '> CROSS-REFERENCING AGENT DATABASE...' },
  { title: 'ENCRYPTION', sub: 'Applying AES-256 encryption...', log: '> ENCRYPTING CREDENTIALS // AES-256...' },
  { title: 'CLEARANCE CHECK', sub: 'Assigning security level...', log: '> ASSIGNING CLEARANCE LEVEL 3...' },
];

function submitForm() {
  if (!validateStep(5)) return;
  const fName = document.getElementById('firstName').value.trim();
  const lName = document.getElementById('lastName').value.trim();
  const uname = document.getElementById('username').value.trim();
  const region = document.getElementById('region').value;

  document.getElementById('form-content').style.display = 'none';
  document.getElementById('hud-bar').style.display = 'none';
  const ls = document.getElementById('loading-screen');
  ls.style.display = 'flex'; ls.style.flexDirection = 'column'; ls.style.alignItems = 'center';

  let i = 0;
  const ll = document.getElementById('load-log');
  function nextLoad() {
    if (i >= loadSequence.length) { clearInterval(timer); setTimeout(showSuccess, 400); return; }
    const lv = loadSequence[i++];
    document.getElementById('load-title').textContent = lv.title;
    document.getElementById('load-sub').textContent = lv.sub;
    const div = document.createElement('div'); div.className = 'log-line';
    div.innerHTML = `<span class="ok">[OK]</span> ${lv.log}`;
    ll.appendChild(div); ll.scrollTop = ll.scrollHeight;
  }
  nextLoad();
  const timer = setInterval(nextLoad, 620);

  function showSuccess() {
    ls.style.display = 'none';
    const ss = document.getElementById('success-screen');
    ss.style.display = 'flex'; ss.style.flexDirection = 'column'; ss.style.alignItems = 'center';
    const uid = 'G9-' + Math.random().toString(36).slice(2, 8).toUpperCase();
    document.getElementById('su-panel').innerHTML = `
      <div class="su-row"><span class="su-key">OPERATIVE</span><span class="su-val">${(fName + ' ' + lName).toUpperCase()}</span></div>
      <div class="su-row"><span class="su-key">AGENT HANDLE</span><span class="su-val">@${uname.toUpperCase()}</span></div>
      <div class="su-row"><span class="su-key">ROLE</span><span class="su-val">${prof.toUpperCase()}</span></div>
      <div class="su-row"><span class="su-key">DIVISION</span><span class="su-val">${division}</span></div>
      <div class="su-row"><span class="su-key">REGION</span><span class="su-val">${region.replace('-', ' // ')}</span></div>
      <div class="su-row"><span class="su-key">ENVIRONMENT</span><span class="su-val">${workEnv}</span></div>
      <div class="su-row"><span class="su-key">UID</span><span class="su-val gold">${uid}</span></div>
      <div class="su-row"><span class="su-key">STATUS</span><span class="su-val green">● ACTIVE</span></div>
    `;
    document.getElementById('su-stats').innerHTML = `
      <div class="su-stat"><div class="su-stat-val">G9</div><div class="su-stat-label">Network</div></div>
      <div class="su-stat"><div class="su-stat-val" style="color:var(--c3)">●</div><div class="su-stat-label">Online</div></div>
      <div class="su-stat"><div class="su-stat-val" style="color:var(--c5)">${clearanceLevel}</div><div class="su-stat-label">Clearance</div></div>
    `;
  }
}
