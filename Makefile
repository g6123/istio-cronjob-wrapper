split = $(word $3,$(subst $2, ,$1))

all: bin/istio-cronjob-wrapper-darwin-arm64 \
			bin/istio-cronjob-wrapper-linux-arm64 \
			bin/istio-cronjob-wrapper-linux-amd64

bin/istio-cronjob-wrapper-%:
	GOOS=$(call split,$*,-,1) GOARCH=$(call split,$*,-,2) \
		go build -trimpath -o $@ ./cmd/

clean:
	rm -vf bin/*
