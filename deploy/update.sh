#!/bin/bash

gcloud compute instances update-container ${AB_VM} \
  --container-image gcr.io/${AB_GCP_PROJECT}/anonbot:${AB_VERSION_TAG} \
  --container-arg="-oauth=${AB_OAUTH}" \
  --container-arg="-ssecret=${AB_SIGNING_SECRET}" \
  --container-arg="-csecret=${AB_CLIENT_SECRET}" \
  --container-arg="-logtostderr"
  
