# Redis Healtcheck Sidecar
A simple 

## Build binary locally
```bash

make server

```

## Build and push image
- Modify the `img` and `version` variables in [Makefile](Makefile) to reflect the correct Docker image registry and version of the image you'd like to build 
```bash

make image

```