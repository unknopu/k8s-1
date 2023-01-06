#!/bin/bash

sudo apt-get update
sudo apt-get -y install openjdk-8-jdk openjdk-8-jre

chmod +x *.sh

if (! docker --version >/dev/null 2>&1); then
    echo "[!] docker did not install yet"
    echo "[/] installing docker"
    sudo snap install docker --classic
else
    echo "[*] docker is ready to work"
fi

if (! kubectl version >/dev/null 2>&1); then
    echo "[!] kubectl did not install yet"
    echo "[/] installing kubectl"
    sudo snap install kubectl --classic
else
    echo "[*] kubectl is ready to work"
fi