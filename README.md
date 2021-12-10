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

### Main Directory
```
.
├── go.mod
├── go.sum
├── README.md
├── .gitignore
├── Makefile  # code compilation and other instructions
├── api  # restful api registry & http router
├── cmd  # main applications for this project's multiple package
│   └── tirelease
├── scripts  # scripts to perform various build, install, analysis, etc operations, keep the root level Makefile small and simple
│   └── run.sh
└── web  # web components and pages. detail can jump to  web/README.MD
    ├── package.json
    ├── src
    └── public
```