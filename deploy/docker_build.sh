#!/bin/bash

docker build -t anonbot:${AB_VERSION_TAG} -f Dockerfile .
docker tag anonbot:${AB_VERSION_TAG} gcr.io/${AB_GCP_PROJECT}/anonbot:${AB_VERSION_TAG}
docker push gcr.io/${AB_GCP_PROJECT}/anonbot:${AB_VERSION_TAG}
