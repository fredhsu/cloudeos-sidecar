IMAGE=fredhsu/cloudeos-sidecar:0.1.1
all: main.go
	docker build -t $(IMAGE) .
	docker push $(IMAGE)
