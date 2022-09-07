OUT=markdown2json

$(OUT): cmd/$(OUT).go
	@echo building... $(OUT)
	@go build -o out/$(OUT) cmd/$(OUT).go

clean:
	@echo cleaning...
	@rm -rf out
