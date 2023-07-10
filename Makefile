it:
	go mod tidy

run:
	go run cli/unoserver.go

build:
	$(call build,linux)
	$(call build,darwin)

clean:
	rm build/* | true

# define a reusable recipe
define build
	@echo "Building for $(1)..."
	CGO_ENABLED=0 GOOS=$(1) \
		go build -o build/unoserver-$(1) cli/unoserver.go
endef
