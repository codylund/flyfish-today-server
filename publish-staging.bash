set -e
docker buildx build --platform linux/amd64 . -t server-linux
docker tag server-linux us-docker.pkg.dev/streamflows/gcr.io/server-linux:staging
docker push us-docker.pkg.dev/streamflows/gcr.io/server-linux:staging