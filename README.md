> This is a release platform for PingCAP<br>
> Author: VelocityLight<br>
> Url: https://github.com/VelocityLight/tirelease<br>

### Design
[click and jump](https://pingcap.feishu.cn/docs/doccnI803yGKKKeQsh56EdNi3Cc#UeCMnT)

### Technologies
+ backend: golang & gin
+ database: mysql
+ deployment: kubenetes
+ frontend: [create-react-app](https://github.com/facebook/create-react-app) & [material-ui/mui](https://github.com/mui-org/material-ui)

### Quick Run
```
./scripts/run.sh
```

### Project Directory
```
├── Makefile
├── README.md
├── api
├── cmd
│   └── tirelease
├── go.mod
├── go.sum
├── scripts
│   └── run.sh
└── web
    ├── README.md
    ├── build
    ├── node_modules
    ├── package.json
    ├── public
    ├── src
    └── yarn.lock
```