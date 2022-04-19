#!/bin/bash

set -e

#gcloud auth activate-service-account starsbook-ci-cd@starsbook-346608.iam.gserviceaccount.com \
#  --key-file=./service_account.json --project=starsbook-346608

gcloud config set project starsbook-346608

gsutil -m rsync -r -c bucket-assets gs://starsbook-assets