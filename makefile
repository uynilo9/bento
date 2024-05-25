INPUT = ./cmd/bento/bento.go
OUTDIR = ./bin
OUTPUT = $(OUTDIR)/bento

build:
	go build -o $(OUTPUT) $(INPUT)

clean:
	rm -rf $(OUTDIR)