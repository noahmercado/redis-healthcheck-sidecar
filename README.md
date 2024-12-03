# Redis Healthcheck Sidecar
A simple sidecar container for exposing an HTTP compatible healthcheck for a Redis instance running in the same pod

## Build binary locally
```bash

make server

```

## Build and push image
- Authenticate and configure your Docker CLI to Artifact Registry
```bash
gcloud auth configure-docker ${GCP_REGION_OF_ARTIFACT_REGISTRY_REPO}-docker.pkg.dev
```
- Modify the `img` and `version` variables in [Makefile](Makefile) to reflect the correct Docker image registry and version of the image you'd like to build 
### Using local docker
```bash

make image

```

### Using Cloud Build
```bash

make remote-build-image

```