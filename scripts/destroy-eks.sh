#!/bin/bash

# Function to log messages
log() {
    echo "[$(date +'%Y-%m-%d %H:%M:%S')] $*"
}

# Validate required tools
check_dependencies() {
    local dependencies=("eksctl" "kubectl" "aws")
    for cmd in "${dependencies[@]}"; do
        if ! command -v "$cmd" &> /dev/null; then
            log "Error: $cmd is not installed"
            exit 1
        fi
    done
}

# Destroy EKS cluster
destroy_eks() {
    local env_name=$1

    # Validate environment
    if [[ ! "$env_name" =~ ^(staging|production)$ ]]; then
        log "Error: Environment must be 'staging' or 'production'"
        exit 1
    fi

    # Generate cluster name
    CLUSTER_NAME="skip-platform-${env_name}"

    # Confirm destruction
    read -p "Are you sure you want to destroy the $env_name EKS cluster '$CLUSTER_NAME'? (y/N): " confirm
    if [[ ! "$confirm" =~ ^[Yy]$ ]]; then
        log "Cluster destruction cancelled"
        exit 0
    fi

    # Delete Kubernetes resources
    log "Deleting Kubernetes resources..."
    kubectl delete -f k8s/app-${env_name}.yaml || true
    kubectl delete -f k8s/postgres-${env_name}.yaml || true
    kubectl delete secret app-env || true

    # Delete EKS cluster
    log "Destroying EKS cluster: $CLUSTER_NAME"
    eksctl delete cluster --name "$CLUSTER_NAME" --region us-west-2

    # Clean up local configuration
    log "Cleaning up kubectl configuration..."
    kubectl config delete-context "$CLUSTER_NAME"
}

# Main script
main() {
    # Check for required tools
    check_dependencies

    # Prompt for environment
    read -p "Destroy cluster for (staging/production): " DESTROY_ENV

    # Destroy selected environment
    destroy_eks "$DESTROY_ENV"
}

# Run main script
main
