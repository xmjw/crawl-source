run:
	# Lazy local database config...
	SOURCE_URL_DATABASE="postgres://url_sources:simple@localhost:5432/url_sources?sslmode=disable" \
	go run *.go
