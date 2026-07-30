#!/usr/bin/env node

const { spawnSync } = require('child_process');
const path = require('path');
const fs = require('fs');

const binaryName = process.platform === 'win32' ? 'gwho.exe' : 'gwho';
const binaryPath = path.join(__dirname, binaryName);

if (!fs.existsSync(binaryPath)) {
  console.error(`\n❌ Error: File eksekusi 'gwho' tidak ditemukan di ${binaryPath}`);
  console.error(`Pastikan proses instalasi tidak bermasalah atau coba install ulang.\n`);
  process.exit(1);
}

const result = spawnSync(binaryPath, process.argv.slice(2), { stdio: 'inherit' });

process.exit(result.status || 0);
