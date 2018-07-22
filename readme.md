# What?

Keep a log/db of top 1m domains.

Download http://downloads.majestic.com/majestic_million.csv

# Plan

- Use GoLand apps as micro services
- Scrape each site, index keywords
- Provide an API graph to query the domains
- For each domain, find the domains/pages it links to. Fetch them.
- Check the host/dns of every domain and see what else we can find. Fetch them.

Create a Google BigTable dataset of every website and the graph of everything it links to.

## Structure

Components that we need:

- Something to fetch the URLs and write them to a table every n-days. (Is this necessary, once we have a set surely we just keep working from that?)
- Something to fetch each site craw the pages for links.

Other things?
