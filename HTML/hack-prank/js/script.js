
// binary rain
const rainEl = document.getElementById('rain');
const cols = Math.floor(window.innerWidth / 18);
for (let i = 0; i < cols; i++) {
  const c = document.createElement('div'); c.className = 'rc';
  c.style.left = (i * 18) + 'px';
  c.style.animationDuration = (6 + Math.random() * 14) + 's';
  c.style.animationDelay = (-Math.random() * 20) + 's';
  let t = ''; for (let j = 0; j < 50; j++)t += Math.random() > 0.5 ? '1' : '0';
  c.textContent = t; rainEl.appendChild(c);
}

let userName = '';

// ── ENTRY FORM ──
function startVerify() {
  const fname = document.getElementById('inp-fname').value.trim();
  const lname = document.getElementById('inp-lname').value.trim();
  const device = document.getElementById('inp-device').value.trim();

  if (!fname) { document.getElementById('inp-fname').focus(); return; }

  userName = fname + (lname ? ' ' + lname : '');
  const userDevice = device || 'Unknown Device';

  // store for later
  window._userDevice = userDevice;

  // switch to verify screen
  document.getElementById('entry-screen').style.display = 'none';
  document.getElementById('verify-screen').style.display = 'block';

  runVerifySequence();
}

// allow Enter key on inputs
document.querySelectorAll('.entry-input').forEach(inp => {
  inp.addEventListener('keydown', e => { if (e.key === 'Enter') startVerify(); });
});

// ── VERIFY SEQUENCE ──
function runVerifySequence() {
  const steps = ['vs1', 'vs2', 'vs3', 'vs4'];
  const bar = document.getElementById('verify-bar');
  let pct = 0;

  steps.forEach((id, i) => {
    setTimeout(() => {
      // mark previous done
      if (i > 0) { document.getElementById(steps[i - 1]).classList.remove('active'); document.getElementById(steps[i - 1]).classList.add('done'); }
      document.getElementById(id).classList.add('active');
      pct = (i + 1) / steps.length * 100;
      bar.style.width = pct + '%';

      if (i === steps.length - 1) {
        setTimeout(() => {
          document.getElementById(id).classList.remove('active');
          document.getElementById(id).classList.add('done');
          bar.style.width = '100%';
          setTimeout(launchHackScreen, 600);
        }, 700);
      }
    }, i * 900);
  });
}

// ── LAUNCH HACK SCREEN ──
function launchHackScreen() {
  document.getElementById('verify-screen').style.display = 'none';
  document.getElementById('main').style.display = 'block';

  // personalize
  document.getElementById('hack-title').innerHTML = userName.split(' ')[0].toUpperCase() + ',<br>YOU\'VE BEEN HACKED';
  document.getElementById('hack-sub').textContent = '// ' + userName.toUpperCase() + ' — SYSTEM COMPROMISED — WIPE INITIATED';
  document.getElementById('target-val').textContent = userName.toUpperCase();
  document.getElementById('ip-val').textContent = '192.168.' + Math.floor(Math.random() * 255) + '.' + Math.floor(Math.random() * 255);
  document.getElementById('device-val').textContent = window._userDevice;
  document.getElementById('sys-id').textContent = 'SYS-ID: ' + Math.random().toString(36).substr(2, 8).toUpperCase();
  document.getElementById('timestamp').textContent = new Date().toISOString().replace('T', ' ').substr(0, 19);

  // device stats
  loadDeviceStats();

  // file bars
  initFileBars();

  // countdown
  startCountdown();

  // webcam + chat
  setTimeout(() => { document.getElementById('webcam-alert').style.display = 'block'; }, 2000);
  startChat();

  // scan line
  document.querySelector('.countdown-section').insertAdjacentHTML('beforeend',
    '<div class="scan-line"></div><div class="pulse-ring"></div><div class="pulse-ring-2"></div>');
}

// ── DEVICE STATS ──
async function loadDeviceStats() {
  let battPct = Math.floor(60 + Math.random() * 35);
  try { const b = await navigator.getBattery(); battPct = Math.floor(b.level * 100); } catch (e) { }
  setTimeout(() => {
    document.getElementById('battery-val').textContent = battPct + '%';
    document.getElementById('battery-bar').style.width = battPct + '%';
    const s = Math.floor(40 + Math.random() * 50);
    document.getElementById('storage-val').textContent = s + '% COPIED';
    document.getElementById('storage-bar').style.width = s + '%';
    const r = Math.floor(55 + Math.random() * 40);
    document.getElementById('ram-val').textContent = r + '% DUMPED';
    document.getElementById('ram-bar').style.width = r + '%';
    const cpu = Math.floor(70 + Math.random() * 25);
    document.getElementById('cpu-val').textContent = cpu + '% LOAD';
    document.getElementById('cpu-bar').style.width = cpu + '%';
  }, 1200);
}

// ── FILE BARS ──
const fakeFiles = [
  '/home/' + '{name}' + '/Documents/passwords.txt',
  '/Users/Desktop/bank_details.xlsx',
  'C:\\Users\\Photos\\private_2024\\',
  '/var/log/system_keys.json',
  'C:\\Windows\\System32\\credentials',
  '/home/.ssh/id_rsa',
];
let bars = [];

function initFileBars() {
  const el = document.getElementById('file-bars');
  const first = userName.split(' ')[0].toLowerCase();
  fakeFiles.forEach((f, i) => {
    const fname = f.replace('{name}', first);
    const row = document.createElement('div'); row.className = 'file-row';
    row.innerHTML = `<div class="file-name"><span>${fname}</span><span class="file-pct" id="pct-${i}">0%</span></div><div class="progress-track"><div class="progress-fill" id="bar-${i}"></div></div>`;
    el.appendChild(row);
    bars.push({ bar: null, pct: null, progress: 0, speed: 3 + Math.random() * 6 });
  });
  bars.forEach((_, i) => { bars[i].bar = document.getElementById('bar-' + i); bars[i].pct = document.getElementById('pct-' + i); });
}

// ── TERMINAL ──
const logs = [
  { t: '> INITIATING INTRUSION PROTOCOL...', c: 'dim' },
  { t: '> FIREWALL BYPASSED ✓', c: 'green' },
  { t: '> SCANNING FILE SYSTEM...', c: 'dim' },
  { t: '> ROOT ACCESS GRANTED ✓', c: 'green' },
  { t: '> LOCATING TARGET FILES...', c: 'dim' },
  { t: '> ENCRYPTING USER DATA...', c: 'dim' },
  { t: '> UPLOADING TO REMOTE SERVER...', c: 'dim' },
  { t: '> WIPING RECOVERY PARTITIONS...', c: 'dim' },
  { t: '> DELETION SEQUENCE ACTIVE ✓', c: 'green' },
  { t: '> WARNING: PROCESS IRREVERSIBLE', c: '' },
  { t: '> IDENTITY CONFIRMED: ' + '{name}', c: 'green' },
  { t: '> GOODBYE 👋', c: '' },
];
let logIdx = 0;
function addLog() {
  const termEl = document.getElementById('terminal');
  if (logIdx >= logs.length) return;
  const entry = logs[logIdx];
  const l = document.createElement('div');
  l.className = 'log-line ' + (entry.c || '');
  l.textContent = entry.t.replace('{name}', userName.toUpperCase());
  termEl.appendChild(l);
  termEl.scrollTop = termEl.scrollHeight;
  logIdx++;
}

// ── COUNTDOWN ──
function startCountdown() {
  let timeLeft = 12;
  const countEl = document.getElementById('countdown');
  const mainEl = document.getElementById('main');
  function animateBars() { bars.forEach(b => { if (b.progress < 100) { b.progress = Math.min(100, b.progress + b.speed); b.bar.style.width = b.progress + '%'; b.pct.textContent = Math.floor(b.progress) + '%'; } }); }
  const interval = setInterval(() => {
    timeLeft--;
    countEl.textContent = timeLeft;
    animateBars();
    addLog();
    if (timeLeft <= 4) {
      countEl.classList.add('urgent');
      bars.forEach(b => b.speed = 18);
      mainEl.classList.add('scanning');
    }
    if (timeLeft <= 0) {
      clearInterval(interval);
      bars.forEach(b => { b.bar.style.width = '100%'; b.pct.textContent = '100%'; });
      setTimeout(showGotcha, 700);
    }
  }, 1000);
}

// ── HACKER CHAT (uses name) ──
function startChat() {
  const first = userName.split(' ')[0];
  const messages = [
    { t: 'got root access 😈', delay: 3000 },
    { t: 'found ' + first + '\'s files...', delay: 5000 },
    { t: 'lol they have no idea 💀', delay: 7000 },
    { t: 'passwords.txt found ✓', delay: 9000 },
    { t: 'almost done with ' + first + ' 👀', delay: 11000 },
  ];
  setTimeout(() => {
    document.getElementById('hacker-chat').style.display = 'block';
    const chatBody = document.getElementById('chat-body');
    messages.forEach(msg => {
      setTimeout(() => {
        const m = document.createElement('div'); m.className = 'chat-msg';
        m.innerHTML = '<span class="sender">haxor_x:</span>' + msg.t;
        chatBody.appendChild(m); chatBody.scrollTop = chatBody.scrollHeight;
      }, msg.delay - 3000);
    });
  }, 3000);
}

// ── GOTCHA ──
function showGotcha() {
  const first = userName.split(' ')[0];
  document.getElementById('gotcha-name').textContent = '😂 GOT YOU, ' + first.toUpperCase() + '! 😂';
  document.getElementById('gotcha').classList.add('show');
  spawnConfetti();
}

function spawnConfetti() {
  const colors = ['#00ff41', '#ff0033', '#ffcc00', '#00e5ff', '#ff6600'];
  for (let i = 0; i < 120; i++) {
    setTimeout(() => {
      const c = document.createElement('div'); c.className = 'confetti-piece';
      c.style.left = Math.random() * 100 + 'vw'; c.style.top = '-20px';
      c.style.background = colors[Math.floor(Math.random() * colors.length)];
      c.style.borderRadius = Math.random() > 0.5 ? '50%' : '0';
      c.style.width = (6 + Math.random() * 10) + 'px'; c.style.height = (6 + Math.random() * 10) + 'px';
      c.style.animationDuration = (2 + Math.random() * 3) + 's';
      document.body.appendChild(c); setTimeout(() => c.remove(), 5000);
    }, i * 30);
  }
}

function shareLink() {
  const url = window.location.href;
  if (navigator.share) { navigator.share({ title: 'Check this out 👀', url }); }
  else if (navigator.clipboard) { navigator.clipboard.writeText(url).then(() => alert('Link copied! Send it to your friends 😂')); }
  else { prompt('Copy this link:', url); }
}

function restart() { location.reload(); }
