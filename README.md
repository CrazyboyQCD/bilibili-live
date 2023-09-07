# bilibili-live
哔哩哔哩录播器

## 功能
- 定时录制
- 尽可能的开箱即用

## 依赖
- [ffmpeg](https://www.gyan.dev/ffmpeg/builds/), 请将ffmpeg放在环境变量里
- [streamlink](https://streamlink.github.io/)，建议用[Chocolatey](https://chocolatey.org/packages/streamlink)安装: ```choco install streamlink```

## 运行
```
git clone https://github.com/CrazyboyQCD/bilibili-live.git.git
修改config.yml
go build ./main.go
./main.exe
```