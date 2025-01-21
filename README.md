## How to Run:

1. Create an EKS cluster using the deploy-eks.sh script. You'll need the aws CLI configured on your machine.

2. Add your AWS secret access key and access key ID to the Github Actions secrets. The secrets should be called AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY.

3. Run the Github Action on the Github console to initiate the CI/CD pipeline. You'll need to specify whether you're deploying to staging or production by choosing from the dropdown menu. Default is staging.

## Solutions:

1. The deployment process is manual and error-prone. Whenever a deployment happens, we suffer a tiny bit of downtime due to the server being down.

* I've created a CI/CD pipeline using Github Workflows that deploys a containerized version of the application to EKS. The pipeline uses blue/green deployments to eliminate downtime and errors.
 
2. There's no standardization for code formatting which leads to inconsistencies

* I've implemented linting in the CI/CD pipeline. Code must pass linting in order to proceed to the build stage. Developers can view the output to see any errors that need to be fixed.

3. The proto-gen script is ran locally which leads to developers forgetting to do it before pushing code upstream

* The proto-gen script is run in the CI/CD pipeline. 

4. Downloading the tooling dependencies is a manual, undocumented process

* Tooling is downloaded in the CI/CD pipeline and cached to reduce overall build times.

5. Running the server locally is a bit of a pain, since it requires manually standing up a Postgres database (to replicate prod)

* I've added a database migration script which is used to set up a Postgres database with the Docker build. This is implemented in the CI/CD pipeline.

6. There's no easy way to share a feature update with our colleagues, since we only have a local and production environment.

* I've added scripts to easily set up identical staging and production environments using EKS.


## Files Added

1. Dockerfile - The Dockerfile builds the containerized application and also runs the database migrations for the Postgres database.

2. .golangci.yml - This file manages the linting configuration in the CI/CD pipeline.

3. /scripts/deploy-eks.sh - This interactive script can be used to deploy staging and production environments using EKS.

4. /scripts/destroy-eks.sh - This can be used to destroy an EKS environment.

5. /scripts/postgres-migrate.sh - This script runs the postgres-migrations.

6. /.github/workflows/deploy.yml - This contains the CI/CD pipeline which is run via Github Actions.

## Application Changes

#### I made a few minor changes to the application in order to enhance deployment. 

* I had to set the version of go used in the application to 1.22 (the latest stable) in order to effectively containerize the application. The application was originally written using go version 1.23. 

* I had to fix some linting errors in order to make the CI/CD pipeline complete without error. 

* I also added a health check route to the application.
