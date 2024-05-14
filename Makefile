# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

# Main application name
APP_NAME=Ascii-Art-Web

# Directories 
SRC_DIR=./
WEB_APP_DIR=$(SRC_DIR)cmd/webApp/

# Files
MAIN_FILE=server.go

# Binary names
WEB_BINARY=asciiartweb

# Targets
.PHONY: all cli web clean run

all: cli web


web:
	$(GOBUILD) -o $(WEB_BINARY) $(WEB_APP_DIR)$(MAIN_FILE)


run-web: web
		$(SRC_DIR)$(WEB_BINARY)

clean:
	@echo cleaning...
	@$(GOCLEAN)
	rm -f $(WEB_BINARY)
