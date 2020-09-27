get:
	go get -v -u -t -d ./redash/...

lint: get
	go vet ./redash/...
	gofmt -l ./redash/

test: get lint
	go test -v -cover ./redash/...

coverage: get
	go test -coverprofile=cover.out -covermode=count ./redash/...
	go tool cover -html=cover.out

doc-deps:
	go get -u github.com/robertkrimen/godocdown/godocdown

doc: doc-deps
	rm -f doc/redash.md
	godocdown redash > doc/redash.md

.PHONY: get lint test coverage doc-deps doc
