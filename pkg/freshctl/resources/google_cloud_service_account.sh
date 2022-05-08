mkdir -p .freshcloud
gcloud iam service-accounts create {{index . "GCP_PROJECT_ID" }}-service-account --display-name={{index . "GCP_PROJECT_ID" }}-service-account
gcloud iam service-accounts keys create .freshcloud/{{index . "GCP_PROJECT_ID" }}-service-account.json \
  --iam-account={{index . "GCP_PROJECT_ID" }}-service-account@{{index . "GCP_PROJECT_ID" }}.iam.gserviceaccount.com
gcloud projects add-iam-policy-binding {{index . "GCP_PROJECT_ID" }} \
  --member=serviceAccount:{{index . "GCP_PROJECT_ID" }}-service-account@{{index . "GCP_PROJECT_ID" }}.iam.gserviceaccount.com \
  --role=roles/owner