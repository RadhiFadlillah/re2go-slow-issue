run: generate
	@CGO_ENABLED=0 go run *.go

generate:
	@for name in *.re; do \
		RE_IN=$$name; \
		RE_OUT=$$(echo $$name | sed 's/\.re/.go/'); \
		re2go -F --input-encoding utf8 --utf8 --no-generation-date -i $$RE_IN -o $$RE_OUT; \
		gofmt -w $$RE_OUT; \
	done