# Skip Platform Take-Home Challenge Setup Instructions

## Local Development Setup

### Prerequisites
- Docker
- Docker Compose
- Git

### Cloning the Repository
```bash
git clone <repository-url>
cd platform-take-home
```

### Environment Configuration
1. Copy the example environment file:
```bash
cp .env.example .env
```

2. Edit `.env` and replace placeholder values with your configuration:
```bash
# Customize database credentials and connection details
nano .env
```

### Running the Application

#### Using Docker Compose
```bash
# Build and start the services
docker-compose up --build

# To run in detached mode
docker-compose up -d --build
```

#### Services
- REST API: http://localhost:8080/items
- gRPC Server: localhost:9008
- Metrics Endpoint: localhost:8081

### Database Migration
- Migrations are automatically run during service startup
- Uses a dedicated migration script in `scripts/postgres-migrate.sh`
- Supports Postgres database

### Environment Configuration Details
- Credentials stored in `.env` file (gitignored)
- `.env.example` provides a template for configuration
- `POSTGRES_DSN` environment variable used for database connection

### API Endpoints
- `GET /items`: List all items
- `GET /items/{id}`: Get a specific item
- `POST /items`: Create a new item

## EKS Deployment

### Prerequisites
- AWS CLI configured (`aws configure`)
- eksctl installed
- kubectl installed
- AWS account with EKS permissions

### Deployment Steps
1. Configure AWS Credentials
```bash
aws configure
```

2. Prepare Environment
```bash
# Copy and edit environment configuration
cp .env.example .env
nano .env  # Customize your configuration
```

3. Deploy to EKS
```bash
# Script supports staging and production environments
./scripts/deploy-eks.sh
```

### Deployment Notes
- Interactively choose staging or production
- Creates EKS cluster automatically
- Configures Kubernetes secrets
- Deploys Postgres and application services
- Region defaults to us-west-2 (customizable in script)

### Destroying EKS Cluster
```bash
# Interactively choose environment to destroy
./scripts/destroy-eks.sh
```

#### Destroy Options
- Supports staging and production environments
- Deletes Kubernetes resources
- Removes EKS cluster
- Cleans up local kubectl configuration
- Requires explicit user confirmation

## Troubleshooting
- Ensure Docker and Docker Compose are installed
- Check Docker logs for any startup issues
- Verify network ports are not in use by other services
- Confirm `.env` file is correctly configured
- For EKS deployment, ensure AWS credentials are correctly set up

## Development Notes
- Postgres-only backend
- Containerized application
- Automated migration process
- Secure environment configuration
- Supports local and EKS deployment
- Provides EKS cluster management scripts
