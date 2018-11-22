get:
	go get ./redash/...

lint: get
	go tool vet redash
	gofmt -e -l `find . -name '*go'`

test: get lint
	go test -v -cover ./redash/...

doc-deps:
	go get -u github.com/robertkrimen/godocdown/godocdown

doc: doc-deps
	rm -rf doc && mkdir doc
	godocdown redash > doc/redash.md

.PHONY: get lint test doc-deps doc
