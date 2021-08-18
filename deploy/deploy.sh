#!/bin/bash

gcloud compute instances create-with-container ${AB_VM} \
  --container-image gcr.io/${AB_GCP_PROJECT}/anonbot:${AB_VERSION_TAG} \
  --machine-type f1-micro \
  --boot-disk-size=1TB \
  --tags app \
  --address ${AB_GCP_IP} \
  --container-arg="-oauth=${AB_OAUTH}" \
  --container-arg="-ssecret=${AB_SIGNING_SECRET}" \
  --container-arg="-csecret=${AB_CLIENT_SECRET}" \
  --container-arg="-logtostderr"
  
