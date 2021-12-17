# TiRelease
[![GoDoc](https://godoc.org/github.com/kubernetes/test-infra?status.svg)](https://godoc.org/github.com/kubernetes/test-infra)
[![Build status](https://prow.k8s.io/badge.svg?jobs=post-test-infra-bazel)](https://testgrid.k8s.io/sig-testing-misc#post-bazel)

This repository is a release platform for PingCAP, Welcome bros!

## Design
[click and jump](https://pingcap.feishu.cn/docs/doccnI803yGKKKeQsh56EdNi3Cc#UeCMnT)

## Technologies
+ backend: golang & gin
+ database: mysql
+ deployments: kubenetes
+ frontend: [create-react-app](https://github.com/facebook/create-react-app) & [material-ui/mui](https://github.com/mui-org/material-ui)

## Quick Run
```
git clone https://github.com/VelocityLight/tirelease.git
cd tirelease/
make run
```
After waiting a few seconds, application is available and can be visited in the browser:[localhost:8080](http://localhost:8080/)

## Main Directory
```
.
├── go.mod
├── go.sum
├── README.md
├── .gitignore
├── Makefile # code compilation and other instructions
├── Dockerfile # build docker image contains website and server binary
├── cmd/ # main applications for this project's multiple package
│   └── tirelease/
├── scripts/  # scripts to perform various build, install, analysis, etc operations, keep the root level Makefile small and simple
├── deploy/ # profile for deployment 
│   └── kubernetes
└── website/ # ui components and pages. detail can jump to  website/README.MD
    ├── src
    ├── public
    └── package.json
```