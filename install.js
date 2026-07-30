const fs = require('fs');
const path = require('path');
const { pipeline } = require('stream');
const { promisify } = require('util');
const axios = require('axios');
const tar = require('tar');

const streamPipeline = promisify(pipeline);

const VERSION = require('./package.json').version;
const REPO = 'Marchiz3D/gwho';

const osMap = {
  win32: 'windows',
  darwin: 'darwin',
  linux: 'linux'
};

const archMap = {
  x64: 'amd64',
  arm64: 'arm64'
};

const os = osMap[process.platform];
const arch = archMap[process.arch];

if (!os || !arch) {
  console.error(`❌ Sistem operasi atau arsitektur tidak didukung: ${process.platform} ${process.arch}`);
  process.exit(1);
}

const binDir = path.join(__dirname, 'bin');
if (!fs.existsSync(binDir)) {
  fs.mkdirSync(binDir);
}
const fileName = `gwho_${os}_${arch}.tar.gz`;
const url = `https://github.com/${REPO}/releases/download/v${VERSION}/${fileName}`;
const tarPath = path.join(__dirname, fileName);

async function downloadAndExtract() {
  console.log(`\n⏳ Mengunduh gwho v${VERSION} untuk ${os}-${arch}...`);
  console.log(`🔗 URL: ${url}`);

  try {
    const response = await axios({
      url: url,
      method: 'GET',
      responseType: 'stream'
    });

    await streamPipeline(response.data, fs.createWriteStream(tarPath));

    console.log('📦 Mengekstrak file binary...');
    await tar.x({
      file: tarPath,
      cwd: binDir,
    });

    fs.unlinkSync(tarPath);

    const binaryName = os === 'windows' ? 'gwho.exe' : 'gwho';
    const binaryPath = path.join(binDir, binaryName);
    if (fs.existsSync(binaryPath) && os !== 'windows') {
      fs.chmodSync(binaryPath, '755');
    }

    console.log('✅ Instalasi gwho via NPM berhasil!\n');
  } catch (err) {
    console.error('\n❌ Gagal mendownload atau mengekstrak file binary.');
    console.error('Pastikan kamu sudah merilis (membuat Release) versi ini di GitHub.');
    console.error(`Pesan Error: ${err.message}\n`);
  }
}

downloadAndExtract();
