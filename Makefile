lint:
	go vet ./redash/...
	gofmt -l ./redash/

test: lint
	go test -v -cover ./redash/...

coverage:
	go test -coverprofile=cover.out -covermode=count ./redash/...
	go tool cover -html=cover.out

doc-deps:
	go get -u github.com/robertkrimen/godocdown/godocdown

doc: doc-deps
	rm -f doc/redash.md
	godocdown redash > doc/redash.md

.PHONY: lint test coverage doc-deps doc
