.PHONY: go goonly

go:
	EXAMPLE_LANGUAGE=go docker compose up --build

goproducer:
	EXAMPLE_LANGUAGE=go docker compose up --build producer

goconsumer:
	EXAMPLE_LANGUAGE=go docker compose up --build consumer
