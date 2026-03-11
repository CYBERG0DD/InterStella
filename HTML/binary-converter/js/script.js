
let mode = 'encode';
let outputText = '';
let animTimeout = null;

// === Binary Rain ===
const rainContainer = document.getElementById('binary-rain');
const cols = Math.floor(window.innerWidth / 18);
for (let i = 0; i < cols; i++) {
    const col = document.createElement('div');
    col.className = 'rain-col';
    col.style.left = (i * 18) + 'px';
    col.style.animationDuration = (8 + Math.random() * 16) + 's';
    col.style.animationDelay = (-Math.random() * 20) + 's';
    col.style.fontSize = (9 + Math.random() * 4) + 'px';
    let bits = '';
    for (let j = 0; j < 60; j++) bits += (Math.random() > 0.5 ? '1' : '0');
    col.textContent = bits;
    rainContainer.appendChild(col);
}

// === Mode Toggle ===
function setMode(m) {
    mode = m;
    const isHex = m === 'hexencode' || m === 'hexdecode';

    // button states
    ['encode', 'decode', 'hexencode', 'hexdecode'].forEach(id => {
        const el = document.getElementById('btn-' + id);
        el.classList.remove('active', 'hex-active');
    });
    const activeBtn = document.getElementById('btn-' + m);
    activeBtn.classList.add(isHex ? 'hex-active' : 'active');

    // panel hex styling
    const panel = document.querySelector('.panel');
    panel.classList.toggle('hex-mode', isHex);

    const panelTitle = document.getElementById('panel-title');
    const inputLabel = document.getElementById('input-label-text');
    const inputArea = document.getElementById('input-area');

    if (m === 'encode') {
        panelTitle.textContent = 'ENCODE // TEXT TO BINARY';
        inputLabel.textContent = 'INPUT TEXT';
        inputArea.placeholder = 'Type something to encode...';
    } else if (m === 'decode') {
        panelTitle.textContent = 'DECODE // BINARY TO TEXT';
        inputLabel.textContent = 'INPUT BINARY (SPACE-SEPARATED 8-BIT)';
        inputArea.placeholder = '01001000 01100101 01101100 01101100 01101111...';
    } else if (m === 'hexencode') {
        panelTitle.textContent = 'ENCODE // TEXT TO HEX';
        inputLabel.textContent = 'INPUT TEXT';
        inputArea.placeholder = 'Type something to convert to hex...';
    } else if (m === 'hexdecode') {
        panelTitle.textContent = 'DECODE // HEX TO TEXT';
        inputLabel.textContent = 'INPUT HEX (SPACE-SEPARATED)';
        inputArea.placeholder = '48 65 6c 6c 6f 20 57 6f 72 6c 64...';
    }

    clearAll();
}

// === Validation ===
function isValidText(text) {
    return text.length > 0;
}

function isValidBinary(bin) {
    return /^[01 ]+$/.test(bin);
}

function isValidHex(hex) {
    return /^[0-9a-fA-F ]+$/.test(hex.trim());
}

// === Binary Encode/Decode ===
function encodeBinary(text) {
    return text.split('').map(c =>
        c.charCodeAt(0).toString(2).padStart(8, '0')
    ).join(' ');
}

function decodeBinary(binary) {
    const words = binary.trim().split(/\s+/);
    return words.map(w => {
        const code = parseInt(w, 2);
        if (isNaN(code)) throw new Error(`Invalid binary: ${w}`);
        return String.fromCharCode(code);
    }).join('');
}

// === Hex Encode/Decode ===
function encodeHex(text) {
    return text.split('').map(c =>
        c.charCodeAt(0).toString(16).padStart(2, '0').toUpperCase()
    ).join(' ');
}

function decodeHex(hex) {
    const bytes = hex.trim().split(/\s+/);
    return bytes.map(b => {
        const code = parseInt(b, 16);
        if (isNaN(code)) throw new Error(`Invalid hex byte: ${b}`);
        return String.fromCharCode(code);
    }).join('');
}

// === Input handler ===
function handleInput() {
    const val = document.getElementById('input-area').value;
    const count = val.length;
    const cc = document.getElementById('char-count');
    cc.textContent = count + ' chars';
    cc.classList.toggle('warn', count > 200);

    document.getElementById('input-area').classList.remove('error');
    document.getElementById('error-msg').textContent = '';
}

function handleKeydown(e) {
    if (e.key === 'Enter' && (e.ctrlKey || e.metaKey)) {
        e.preventDefault();
        runTranslate();
    }
}

// === Translate ===
function runTranslate() {
    const input = document.getElementById('input-area').value.trim();
    const errorEl = document.getElementById('error-msg');
    const inputEl = document.getElementById('input-area');

    inputEl.classList.remove('error');
    errorEl.textContent = '';

    if (!input) {
        errorEl.textContent = '⚠ INPUT CANNOT BE EMPTY';
        inputEl.classList.add('error');
        shake(inputEl);
        return;
    }

    if (mode === 'encode') {
        try {
            outputText = encodeBinary(input);
            showOutput(outputText, 'BINARY ENCODE');
            updateStats(input.length, outputText.replace(/ /g, '').length, input.length, 'BIN ENC');
        } catch (e) {
            errorEl.textContent = '⚠ ENCODING FAILED: ' + e.message;
            inputEl.classList.add('error');
        }

    } else if (mode === 'decode') {
        if (!isValidBinary(input)) {
            errorEl.textContent = '⚠ ONLY 0, 1, AND SPACES ARE ALLOWED';
            inputEl.classList.add('error');
            shake(inputEl);
            return;
        }
        try {
            outputText = decodeBinary(input);
            const groups = input.trim().split(/\s+/);
            showOutput(outputText, 'BINARY DECODE');
            updateStats(outputText.length, input.replace(/ /g, '').length, groups.length, 'BIN DEC');
        } catch (e) {
            errorEl.textContent = '⚠ DECODE FAILED: ' + e.message;
            inputEl.classList.add('error');
        }

    } else if (mode === 'hexencode') {
        try {
            outputText = encodeHex(input);
            showOutput(outputText, 'HEX ENCODE');
            updateStats(input.length, outputText.replace(/ /g, '').length, input.length, 'HEX ENC');
        } catch (e) {
            errorEl.textContent = '⚠ HEX ENCODING FAILED: ' + e.message;
            inputEl.classList.add('error');
        }

    } else if (mode === 'hexdecode') {
        if (!isValidHex(input)) {
            errorEl.textContent = '⚠ ONLY HEX CHARACTERS (0-9, A-F) AND SPACES ALLOWED';
            inputEl.classList.add('error');
            shake(inputEl);
            return;
        }
        try {
            outputText = decodeHex(input);
            const groups = input.trim().split(/\s+/);
            showOutput(outputText, 'HEX DECODE');
            updateStats(outputText.length, input.replace(/ /g, '').length, groups.length, 'HEX DEC');
        } catch (e) {
            errorEl.textContent = '⚠ HEX DECODE FAILED: ' + e.message;
            inputEl.classList.add('error');
        }
    }
}

// === Animated Output ===
function showOutput(text, modeLabel) {
    const box = document.getElementById('output-box');
    box.classList.remove('empty');
    box.classList.add('has-result');
    box.classList.add('glitch');
    setTimeout(() => box.classList.remove('glitch'), 300);

    if (animTimeout) clearTimeout(animTimeout);
    box.textContent = '';

    const delay = Math.min(20, 800 / text.length);
    let i = 0;
    function addChar() {
        if (i < text.length) {
            box.textContent += text[i];
            i++;
            animTimeout = setTimeout(addChar, delay);
        }
    }
    addChar();

    document.getElementById('copy-btn').style.display = 'block';
}

// === Stats ===
function updateStats(chars, bits, bytes, modeLabel) {
    document.getElementById('stat-chars').textContent = chars;
    document.getElementById('stat-bits').textContent = bits;
    document.getElementById('stat-bytes').textContent = bytes;
    document.getElementById('stat-mode').textContent = modeLabel;

    const bar = document.getElementById('stats-bar');
    bar.classList.add('visible');
}

// === Copy ===
function copyOutput() {
    if (!outputText) return;
    const btn = document.getElementById('copy-btn');

    // Try modern clipboard API first, fall back to execCommand
    const doSuccess = () => {
        btn.textContent = 'COPIED!';
        btn.classList.add('copied');
        setTimeout(() => {
            btn.textContent = 'COPY';
            btn.classList.remove('copied');
        }, 1800);
    };

    if (navigator.clipboard && navigator.clipboard.writeText) {
        navigator.clipboard.writeText(outputText).then(doSuccess).catch(() => fallbackCopy(doSuccess));
    } else {
        fallbackCopy(doSuccess);
    }
}

function fallbackCopy(onSuccess) {
    const ta = document.createElement('textarea');
    ta.value = outputText;
    ta.style.cssText = 'position:fixed;top:-9999px;left:-9999px;opacity:0';
    document.body.appendChild(ta);
    ta.focus();
    ta.select();
    try {
        document.execCommand('copy');
        onSuccess();
    } catch (e) {
        alert('Copy failed. Please select and copy the output manually.');
    }
    document.body.removeChild(ta);
}

// === Clear ===
function clearAll() {
    document.getElementById('input-area').value = '';
    document.getElementById('input-area').classList.remove('error');
    document.getElementById('error-msg').textContent = '';
    document.getElementById('char-count').textContent = '0 chars';
    document.getElementById('char-count').classList.remove('warn');
    document.getElementById('output-box').textContent = 'Awaiting input...';
    document.getElementById('output-box').classList.add('empty');
    document.getElementById('output-box').classList.remove('has-result');
    document.getElementById('stats-bar').classList.remove('visible');
    outputText = '';
    if (animTimeout) clearTimeout(animTimeout);
}

// === Shake animation ===
function shake(el) {
    el.style.animation = 'none';
    requestAnimationFrame(() => {
        el.style.animation = 'glitch 0.3s ease';
        setTimeout(() => el.style.animation = '', 300);
    });
}

// === Keyboard shortcut hint ===
const textarea = document.getElementById('input-area');
const btn = document.getElementById('translate-btn');
const isMac = navigator.platform.toLowerCase().includes('mac');
btn.title = `Translate (${isMac ? '⌘' : 'Ctrl'}+Enter)`;
