// electron.js
const { app, BrowserWindow } = require('electron')

function createWindow() {
    // 创建浏览器窗口
    const win = new BrowserWindow({
        width: 800,
        height: 600,
        webPreferences: {
            nodeIntegration: true,
            contextIsolation: false // 如果使用了较新版本的Electron和Vue CLI插件，可能需要设置为true，并适当调整preload脚本
        }
    })

    // 加载Vue应用的index.html。假设`npm run build`已经生成了dist目录
    // 对于开发环境，您可能需要指向一个开发服务器的地址，如 http://localhost:8080
    win.loadFile('dist/index.html')

    // 打开开发者工具
    win.webContents.openDevTools()
}

app.whenReady().then(createWindow)

app.on('window-all-closed', () => {
    if (process.platform !== 'darwin') {
        app.quit()
    }
})

app.on('activate', () => {
    if (BrowserWindow.getAllWindows().length === 0) {
        createWindow()
    }
})
