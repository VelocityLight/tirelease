# TiRelease
[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)

This repository is a release platform for PingCAP, Welcome bros!

## Design
[click and jump](https://pingcap.feishu.cn/docs/doccnI803yGKKKeQsh56EdNi3Cc#UeCMnT)

## Technologies
+ backend: golang & gin
+ database: mysql
+ deployments: kubenetes
+ frontend: [create-react-app](https://github.com/facebook/create-react-app) & [material-ui/mui](https://github.com/mui-org/material-ui)

## Quick Start
```
git clone https://github.com/VelocityLight/tirelease.git
cd tirelease/
make run
```
After waiting a few seconds, application is available and can be visited in the browser:[localhost:8080](http://localhost:8080/)

## File Structure
```
tirelease
├── .gitignore
├── README.md
├── go.mod             # Golang environment configuration
├── go.sum
├── Makefile           # Code compilation and other instructions
├── Dockerfile         # Build docker image contains website and server binary
└── scripts/           # Scripts to perform various build, install, analysis, etc operations, keep the root level Makefile small and simple
└── api/               # Restful api registry
    ├── routers.go
└── configs/
    ├── config.go      # Load configuration under profiles/
    └── profiles/      # Globle configuration for whole project
└── cmd/               # Main applications for this project's multiple package
    └── tirelease/
└── deploy/            # Profile of deployment 
    └── kubernetes/
└── commons/           # Common utils pkg
    └── database/      # Database connectors
    └── httpclient/    # Http util
    └── github/        # Github client, reference: https://github.com/google/go-github
└── website/           # UI components and pages. detail can jump to  website/README.MD
    ├── yarn.lock      # React environment configuration for machines
    ├── package.json   # React environment configuration for people
    └── src/           # JS/CSS...
    └── public/        # HomePage: index.html and icons...

```