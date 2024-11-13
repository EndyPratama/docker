docker build -t app-test-google-cloud .
docker run -p 4000:80 app-test-google-cloud

gcloud auth login

gcloud auth configure-docker
gcloud auth configure-docker

gcloud config set project docker
gcloud auth configure-docker asia-southeast1-docker.pkg.dev

docker build -t asia-southeast1-docker.pkg.dev/planar-door-405803/myapp-repo/myapp .

docker push asia-southeast1-docker.pkg.dev/planar-door-405803/myapp-repo/myapp