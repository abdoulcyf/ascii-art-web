# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

# Main application name
APP_NAME=Ascii-Art-Web

# Directories 
SRC_DIR=./
WEB_APP_DIR=$(SRC_DIR)cmd/webApp/

# Binary names
WEB_BINARY=asciiartweb

# Targets
.PHONY: all cli web clean run

all: cli web


web:
	$(GOBUILD) -o $(WEB_BINARY) $(WEB_APP_DIR)main.go


run-web: web
		./$(WEB_BINARY)

clean:
	@echo cleaning...
	@$(GOCLEAN)
	@rm -f $(WEB_BINARY)
