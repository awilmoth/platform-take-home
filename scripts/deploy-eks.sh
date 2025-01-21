#!/bin/bash

# Load environment variables
if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
fi

# Function to validate environment
validate_env() {
    local env_name=$1
    if [[ ! "$env_name" =~ ^(staging|production)$ ]]; then
        echo "Error: Environment must be 'staging' or 'production'"
        exit 1
    fi
}

# Function to deploy EKS cluster
deploy_eks() {
    local env_name=$1

    # Validate environment
    validate_env "$env_name"

    # Generate cluster name
    CLUSTER_NAME="skip-platform-${env_name}"

    echo "Deploying EKS cluster: $CLUSTER_NAME"

    # Create EKS cluster configuration
    cat > cluster-config-${env_name}.yaml << EOF
apiVersion: eksctl.io/v1alpha5
kind: ClusterConfig

metadata:
  name: ${CLUSTER_NAME}
  region: us-west-2  # Adjust region as needed

nodeGroups:
  - name: ${env_name}-nodes
    instanceType: t3.medium
    desiredCapacity: 2
    minSize: 1
    maxSize: 3

managedNodeGroups:
  - name: ${env_name}-managed-nodes
    instanceType: t3.medium
    minSize: 1
    maxSize: 3
    desiredCapacity: 2
EOF

    # Deploy EKS cluster
    eksctl create cluster -f cluster-config-${env_name}.yaml

    # Configure kubectl
    aws eks update-kubeconfig --name ${CLUSTER_NAME} --region us-west-2

    # Deploy application
    deploy_application "$env_name"
}

# Function to deploy application to EKS
deploy_application() {
    local env_name=$1

    echo "Deploying application to ${env_name} environment"

    # Create Kubernetes secrets from .env
    kubectl create secret generic app-env \
        --from-env-file=../.env

}

# Main script
main() {
    # Check for required tools
    command -v eksctl >/dev/null 2>&1 || { 
        echo "eksctl is not installed. Please install it first."; 
        exit 1; 
    }

    command -v kubectl >/dev/null 2>&1 || { 
        echo "kubectl is not installed. Please install it first."; 
        exit 1; 
    }

    # Prompt for environment
    read -p "Deploy to (staging/production): " DEPLOY_ENV

    # Deploy to selected environment
    deploy_eks "$DEPLOY_ENV"
}

# Run main script
main
