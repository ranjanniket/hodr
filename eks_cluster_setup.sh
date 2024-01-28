#!/bin/bash

# Install AWS CLI, kubectl, and eksctl
sudo apt-get update
sudo apt-get install -y awscli
sudo apt-get install -y jq

# Install kubectl
curl -o kubectl https://amazon-eks.s3.us-west-2.amazonaws.com/1.19.6/2021-01-05/bin/linux/amd64/kubectl
chmod +x ./kubectl
sudo mv ./kubectl /usr/local/bin
kubectl version --short --client

# Install eksctl
curl --silent --location "https://github.com/weaveworks/eksctl/releases/latest/download/eksctl_$(uname -s)_amd64.tar.gz" | tar xz -C /tmp
sudo mv /tmp/eksctl /usr/local/bin
eksctl version

# Configure AWS CLI with access and secret keys
read -p "Enter your AWS access key: " aws_access_key
read -p "Enter your AWS secret key: " aws_secret_key

aws configure set aws_access_key_id "$aws_access_key"
aws configure set aws_secret_access_key "$aws_secret_key"
aws configure set default.region "us-west-2"  # Update with your desired region

# Create EKS cluster
eksctl create cluster --name my-cluster --version 1.19 --region us-west-2 --nodegroup-name standard-workers --node-type t3.micro --nodes 2 --nodes-min 1 --nodes-max 3 --managed
