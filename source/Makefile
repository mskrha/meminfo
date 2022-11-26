BUILD	= meminfo
VERSION	?= 0.0
ARCH	?= amd64
PKG	?= $(BUILD)_$(VERSION)_$(ARCH)
OUTPUT	= $(BUILD)

all: clean format build-$(ARCH)

clean:
	rm -f $(OUTPUT)

format:
	go fmt

temp:
	$(eval TMPDIR = $(shell mktemp -d))
	$(eval OUTPUT = $(TMPDIR)/usr/bin/$(BUILD))
	mkdir -p $(TMPDIR)/usr/bin

files: temp
	cp -a ../DEBIAN $(TMPDIR)/

build-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o $(OUTPUT) *.go

build-armhf:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -ldflags "-X main.version=$(VERSION)" -o $(OUTPUT) *.go

deb: files build-$(ARCH)
	$(eval SIZE = $(shell du -k $(OUTPUT) | sed 's/^\([0-9]\+\).*/\1/'))
	$(shell sed -i "s/INSTALL_SIZE/$(SIZE)/" $(TMPDIR)/DEBIAN/control)
	$(shell sed -i "s/INSTALL_ARCH/$(ARCH)/" $(TMPDIR)/DEBIAN/control)
	$(shell sed -i "s/INSTALL_VERSION/$(VERSION)/" $(TMPDIR)/DEBIAN/control)
	dpkg-deb --build $(TMPDIR) $(PKG).deb
	rm -rf $(TMPDIR)
