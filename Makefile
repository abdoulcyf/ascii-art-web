# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

# Main application name
APP_NAME=Troupie-Tracker

# Directories 
SRC_DIR=./
WEB_APP_DIR=$(SRC_DIR)cmd/webApp

# Binary names
WEB_BINARY=$(WEB_APP_DIR)/asciiartweb

# Targets
.PHONY: all cli web clean run

all: cli web


web:
	cd $(WEB_APP_DIR) && \
	$(GOBUILD) -o $(WEB_BINARY) main.go


run-web: web
	cd $(WEB_APP_DIR) && \
	$(WEB_BINARY)

clean:
	@echo cleaning...
	@$(GOCLEAN)
	@rm -f $(WEB_BINARY)
