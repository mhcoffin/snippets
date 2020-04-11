#!/bin/zsh

PROJECT_ID=snippets-273801
CREDENTIALS=Snippets-56d0c6c6b468.json
GOOGLE_APPLICATION_CREDENTIALS=/Users/michaelcoffin/credentials/$CREDENTIALS

docker build --tag helloworld .

PORT=8080 && docker run \
-p 9090:${PORT} \
-e PORT=${PORT} \
-e K_SERVICE=dev \
-e K_CONFIGURATION=dev \
-e K_REVISION=dev-00001 \
-e GOOGLE_APPLICATION_CREDENTIALS=/tmp/keys/"$CREDENTIALS" \
-v $GOOGLE_APPLICATION_CREDENTIALS:/tmp/keys/"$CREDENTIALS":ro \
helloworld

