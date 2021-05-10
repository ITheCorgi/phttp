# Parallel HTTP requests

This repository contains an example of parallel http requests from command line. This is a tool which makes http request and prints the address of the request along with the MD5 hash of the response.
The tool must:
    -Be able to perform the requests in parallel so that the tool can somplete sooner. The order in thich addresses are printed is not important.
    - Be able to limit the number of parallel requests, to prevent exhausting local resources. The tool must accept a flag to indicated this limit, and it should default to 10 if the flag is not provided.
    - The tool must have unit test (some of them).