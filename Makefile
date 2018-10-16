get:
	go get ./redash/...

doc-deps:
	go get -u github.com/robertkrimen/godocdown/godocdown

doc: doc-deps
	rm -rf doc && mkdir doc
	godocdown redash > doc/redash.md

.PHONY: get doc-deps doc
