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
├── config.yaml        # Global configuration
├── Makefile           # Code compilation and other instructions
└── scripts/           # Scripts to perform various operations, keep the root level Makefile small and simple
└── cmd/               # Main application starters of Golang
└── api/               # REST API registry & Static file router
    ├── api.go
└── deployments/
    └── docker/        # Build docker image contains website and server binary
    └── kubernetes/    # Deployment yaml for k8s
└── commons/           # Common utils for whole project
    └── configs/       # Global configuration reader
    └── database/      # Database connectors
    └── httpclient/    # Http client utils
    └── github/        # Github client, reference: https://github.com/google/go-github
└── internal/          # Business code & function
    └── entity/        # Object entity
    └── repository/    # Function operator
    └── service/       # Http handler
└── website/           # UI components and pages. detail can jump to  website/README.MD
    ├── yarn.lock      # React environment configuration for machines
    ├── package.json   # React environment configuration for people
    └── src/           # JavaScript/CSS...
    └── public/        # HomePage: index.html and icons
```