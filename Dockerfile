# iron/go:dev is the alpine image with the go tools added
FROM iron/go:dev

WORKDIR /app

# Set an env var that matches your github repo name, replace treeder/dockergo here with your repo name
ENV SRC_DIR=/go/src/github.com/xmjw/crawl-source/

# Add the source code:
ADD . $SRC_DIR

# Fetch the Majestic Millions Domains file and build
RUN curl "http://downloads.majestic.com/majestic_million.csv" -o /app/urls.csv

# Build it:
RUN cd $SRC_DIR; go build -o app; cp app /app/

ENTRYPOINT ["./app"]
