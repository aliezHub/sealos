#!/bin/bash
export readonly ARCH=${1:-amd64}
mkdir -p tars

sealos pull --policy=always --platform=linux/"${ARCH}" ghcr.io/labring/sealos-cloud-user-controller:latest
sealos pull --policy=always --platform=linux/"${ARCH}" ghcr.io/labring/sealos-cloud-terminal-controller:latest
sealos pull --policy=always --platform=linux/"${ARCH}" ghcr.io/labring/sealos-cloud-app-controller:latest
sealos pull --policy=always --platform=linux/"${ARCH}" ghcr.io/labring/sealos-cloud-resources-controller:latest
sealos pull --policy=always --platform=linux/"${ARCH}" ghcr.io/labring/sealos-cloud-account-controller:latest
sealos pull --policy=always --platform=linux/"${ARCH}" ghcr.io/labring/sealos-cloud-license-controller:latest

sealos pull --policy=always --platform=linux/"${ARCH}" ghcr.io/labring/sealos-cloud-desktop-frontend:latest
sealos pull --policy=always --platform=linux/"${ARCH}" ghcr.io/labring/sealos-cloud-terminal-frontend:latest
sealos pull --policy=always --platform=linux/"${ARCH}" ghcr.io/labring/sealos-cloud-applaunchpad-frontend:latest
sealos pull --policy=always --platform=linux/"${ARCH}" ghcr.io/labring/sealos-cloud-dbprovider-frontend:latest
sealos pull --policy=always --platform=linux/"${ARCH}" ghcr.io/labring/sealos-cloud-costcenter-frontend:latest
sealos pull --policy=always --platform=linux/"${ARCH}" ghcr.io/labring/sealos-cloud-template-frontend:latest

sealos pull --policy=always --platform=linux/"${ARCH}" ghcr.io/labring/sealos-cloud-database-service:latest
sealos pull --policy=always --platform=linux/"${ARCH}" ghcr.io/labring/sealos-cloud-job-init-controller:latest

sealos save -o tars/user.tar ghcr.io/labring/sealos-cloud-user-controller:latest
sealos save -o tars/terminal.tar ghcr.io/labring/sealos-cloud-terminal-controller:latest
sealos save -o tars/app.tar ghcr.io/labring/sealos-cloud-app-controller:latest
sealos save -o tars/monitoring.tar ghcr.io/labring/sealos-cloud-resources-controller:latest
sealos save -o tars/account.tar ghcr.io/labring/sealos-cloud-account-controller:latest
sealos save -o tars/license.tar ghcr.io/labring/sealos-cloud-license-controller:latest

sealos save -o tars/frontend-desktop.tar  ghcr.io/labring/sealos-cloud-desktop-frontend:latest
sealos save -o tars/frontend-terminal.tar  ghcr.io/labring/sealos-cloud-terminal-frontend:latest
sealos save -o tars/frontend-dbprovider.tar ghcr.io/labring/sealos-cloud-dbprovider-frontend:latest
sealos save -o tars/frontend-costcenter.tar ghcr.io/labring/sealos-cloud-costcenter-frontend:latest
sealos save -o tars/frontend-applaunchpad.tar ghcr.io/labring/sealos-cloud-applaunchpad-frontend:latest
sealos save -o tars/frontend-template.tar ghcr.io/labring/sealos-cloud-template-frontend:latest

sealos save -o tars/database-service.tar ghcr.io/labring/sealos-cloud-database-service:latest
sealos save -o tars/job-init.tar ghcr.io/labring/sealos-cloud-job-init-controller:latest
