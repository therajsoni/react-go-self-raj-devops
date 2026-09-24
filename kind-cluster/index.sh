#!/bin/bash

sudo apt update && sudo apt upgrade -y
sudo apt install -y docker.io
sudo systemctl enable --now docker
docker --version
sudo usermod -aG docker $USER
newgrp docker