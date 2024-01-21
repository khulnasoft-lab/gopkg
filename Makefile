GLIDE_GO_EXECUTABLE ?= go
DIST_DIRS := find * -type d -exec
VERSION ?= $(shell git describe --tags)
VERSION_INCODE = $(shell perl -ne '/^var version.*"([^"]+)".*$$/ && print "v$$1\n"' gopkg.go)
VERSION_INCHANGELOG = $(shell perl -ne '/^\# Release (\d+(\.\d+)+) / && print "$$1\n"' CHANGELOG.md | head -n1)

build:
	${GLIDE_GO_EXECUTABLE} build -o gopkg -ldflags "-X main.version=${VERSION}" gopkg.go

install: build
	install -d ${DESTDIR}/usr/local/bin/
	install -m 755 ./gopkg ${DESTDIR}/usr/local/bin/gopkg

test:
	${GLIDE_GO_EXECUTABLE} test . ./gb ./path ./action ./tree ./util ./godep ./godep/strip ./gpm ./cfg ./dependency ./importer ./msg ./repo ./mirrors

integration-test:
	${GLIDE_GO_EXECUTABLE} build
	./gopkg up
	./gopkg install

clean:
	rm -f ./gopkg.test
	rm -f ./gopkg
	rm -rf ./dist

bootstrap-dist:
	${GLIDE_GO_EXECUTABLE} get -u github.com/Khulnasoft-lab/gox

build-all:
	gox -verbose \
	-ldflags "-X main.version=${VERSION}" \
	-os="linux darwin windows freebsd openbsd netbsd" \
	-arch="amd64 386 armv5 armv6 armv7 arm64 s390x" \
	-osarch="!darwin/arm64" \
	-output="dist/{{.OS}}-{{.Arch}}/{{.Dir}}" .

dist: build-all
	cd dist && \
	$(DIST_DIRS) cp ../LICENSE {} \; && \
	$(DIST_DIRS) cp ../README.md {} \; && \
	$(DIST_DIRS) tar -zcf gopkg-${VERSION}-{}.tar.gz {} \; && \
	$(DIST_DIRS) zip -r gopkg-${VERSION}-{}.zip {} \; && \
	cd ..

verify-version:
	@if [ "$(VERSION_INCODE)" = "v$(VERSION_INCHANGELOG)" ]; then \
		echo "gopkg: $(VERSION_INCHANGELOG)"; \
	elif [ "$(VERSION_INCODE)" = "v$(VERSION_INCHANGELOG)-dev" ]; then \
		echo "gopkg (development): $(VERSION_INCHANGELOG)"; \
	else \
		echo "Version number in gopkg.go does not match CHANGELOG.md"; \
		echo "gopkg.go: $(VERSION_INCODE)"; \
		echo "CHANGELOG : $(VERSION_INCHANGELOG)"; \
		exit 1; \
	fi

.PHONY: build test install clean bootstrap-dist build-all dist integration-test verify-version
