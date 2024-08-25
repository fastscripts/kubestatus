DOCKER=docker
DOCKER_ARGS=buildx



hello:
	echo "Hallo"


run:
	go run cmd/kubestatus/*.go

build_css:
    npm run build_css

build: build_css
	go build -o kubestatus cmd/kubestatus/*.go

container:
	${DOCKER} ${DOCKER_ARGS} build \
	--tag kubestatus:1 \
    --file ./Dockerfile . 
