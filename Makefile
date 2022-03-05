.PHONY: bench
bench:
	go test -bench . -benchmem -count 10 -timeout 20m
