PHONY_TARGETS    = deps
PHONY_TARGETS   += test

#------------------------------------------------------------------------
# 
#------------------------------------------------------------------------
default:; @echo "Targets"; for i in $(PHONY_TARGETS); do echo "  $$i"; done

deps: 
	go get -u github.com/gorilla/mux

test: 
	go test -v ./cmd/sde/...


.PHONY: $(PHONY_TARGETS)
