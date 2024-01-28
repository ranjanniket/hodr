#!/bin/bash

# Install required packages
sudo apt-get install -y wget apt-transport-https gnupg lsb-release

# Add Trivy repository key
wget -qO - https://aquasecurity.github.io/trivy-repo/deb/public.key | sudo apt-key add -

# Add Trivy repository to sources list
echo "deb https://aquasecurity.github.io/trivy-repo/deb $(lsb_release -sc) main" | sudo tee -a /etc/apt/sources.list.d/trivy.list

# Update package list
sudo apt-get update

# Install Trivy
sudo apt-get install -y trivy

#get the list of images
docker images

# Prompt user for Docker image ID
read -p "Enter Docker image ID for scanning: " image_id

# Perform image scanning using Trivy
trivy image "$image_id"
