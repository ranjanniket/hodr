#!/bin/bash

# Update system packages
sudo apt-get update

# Install necessary dependencies
sudo apt-get install -y apt-transport-https software-properties-common

# Add the GPG key for Grafana
wget -q -O - https://packages.grafana.com/gpg.key | sudo apt-key add -

# Add Grafana repository for stable releases
echo "deb https://packages.grafana.com/oss/deb stable main" | sudo tee -a /etc/apt/sources.list.d/grafana.list

# Update package list and install Grafana
sudo apt-get update
sudo apt-get -y install grafana

# Enable and start Grafana service
sudo systemctl enable grafana-server
sudo systemctl start grafana-server

# Check Grafana status
sudo systemctl status grafana-server

echo "Grafana has been installed and is running."

