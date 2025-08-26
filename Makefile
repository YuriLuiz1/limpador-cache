APP_NAME := chrome-cache-cleaner
VERSION := v1.0
BUILD_TIME := $(shell date +%Y%m%d-%H%M%S)
LDFLAGS := -s -w -X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME)

.PHONY: all clean macos linux macos-intel macos-arm linux-x64 linux-arm64

all: clean macos linux
	@echo "🎉 All builds completed!"
	@ls -la dist/

clean:
	@echo "🧹 Cleaning dist directory..."
	@rm -rf dist/
	@mkdir -p dist/

macos: macos-intel macos-arm
	@echo "✅ macOS builds completed"

linux: linux-x64 linux-arm64
	@echo "✅ Linux builds completed"

macos-intel:
	@echo "📦 Building for macOS Intel (x64)..."
	@GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o "dist/$(APP_NAME)-macos-intel" main.go
	@echo "✅ macOS Intel build: $$(ls -lh dist/$(APP_NAME)-macos-intel | awk '{print $$5}')"

macos-arm:
	@echo "📦 Building for macOS Apple Silicon (ARM64)..."
	@GOOS=darwin GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o "dist/$(APP_NAME)-macos-arm64" main.go
	@echo "✅ macOS ARM64 build: $$(ls -lh dist/$(APP_NAME)-macos-arm64 | awk '{print $$5}')"

linux-x64:
	@echo "📦 Building for Linux x64..."
	@GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o "dist/$(APP_NAME)-linux-x64" main.go
	@echo "✅ Linux x64 build: $$(ls -lh dist/$(APP_NAME)-linux-x64 | awk '{print $$5}')"

linux-arm64:
	@echo "📦 Building for Linux ARM64..."
	@GOOS=linux GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o "dist/$(APP_NAME)-linux-arm64" main.go
	@echo "✅ Linux ARM64 build: $$(ls -lh dist/$(APP_NAME)-linux-arm64 | awk '{print $$5}')"

# Individual platform builds
macos-only: clean macos-intel macos-arm
linux-only: clean linux-x64 linux-arm64

# Make executables executable (for Unix systems)
executable:
	@echo "🔧 Making binaries executable..."
	@chmod +x dist/$(APP_NAME)-*
	@echo "✅ All binaries are now executable"

# Show build info
info:
	@echo "Build Information:"
	@echo "App Name: $(APP_NAME)"
	@echo "Version: $(VERSION)"
	@echo "Build Time: $(BUILD_TIME)"
	@echo "LDFLAGS: $(LDFLAGS)"

# Help
help:
	@echo "Available targets:"
	@echo "  all          - Build for all platforms (macOS + Linux)"
	@echo "  clean        - Clean dist directory"
	@echo "  macos        - Build for macOS (Intel + ARM)"
	@echo "  linux        - Build for Linux (x64 + ARM64)"
	@echo "  macos-intel  - Build for macOS Intel only"
	@echo "  macos-arm    - Build for macOS Apple Silicon only"  
	@echo "  linux-x64    - Build for Linux x64 only"
	@echo "  linux-arm64  - Build for Linux ARM64 only"
	@echo "  executable   - Make all binaries executable"
	@echo "  info         - Show build information"
	@echo "  help         - Show this help"