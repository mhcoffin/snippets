# Cloud Run and Firestore

This is q quick example showing how to deploy a very simple web server using go, cloud run,
and firestore. It reads firestore to get a message and serves it. 

## Set up a Google Cloud project

The project I'm using is called "snippets-273801". Billing is enabled. 

## Enable the Cloud build and Cloud Run APIs

https://cloud.google.com/run/docs/setup

## Running the server locally

Set environment variable for auth:
  
  export GOOGLE_APPLICATION_CREDENTIALS=/Users/michaelcoffin/credentials/Snippets-56d0c6c6b468.json

Run the server:

  go run helloworld.go
  
## Running the server locally from a docker container

The only tricky bit is injecting the credentials into the docker container.

  ./docker-local.sh

Then access at localhost:9090.

## Deploying the server to the cloud

  gcloud builds submit --tag gcr.io/snippets-273801/helloworld
  gcloud run deploy --image gcr.io/snippets-273801/helloworld --platform=managed --region=us-west1

### Shutting it down:

  gcloud run services delete helloworld --platform managed --region us-west1

