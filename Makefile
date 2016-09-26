
test: deps js/write.js js/read.js
	@echo "\n\ntesting go2js...\n"
	go run go/write.go >go2js.gob
	go run go/read.go <go2js.gob
	@echo "\n"
	node js/read.js <go2js.gob

	@echo "\n\ntesting js2go...\n"
	node js/write.js >js2go.gob
	node js/read.js <js2go.gob
	@echo "\n"
	go run go/read.go <js2go.gob

diff: test


deps:
	@echo "checking deps..."
	@which go || echo "install go"
	@which node || echo "install node"
	@which gopherjs || echo "install gopherjs"

js/%.js: go/%.go
	gopherjs build $^ -o $@

clean:
	rm -r js/
	rm *.gob
