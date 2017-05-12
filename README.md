# go-health

## Usage:
Running the container directly with a default 8000 port:
    `docker run -p 8000:8000 omerxx/go-health`

Running with custom port i.e 9292:
    `docker run -e HEALTH_PORT=9292 -p 9292:9292 omerxx/go-health`

Go to 127.0.0.1:<port>/health -> get 200 response
(Local port can be mapped to anything else, the above are only examples)


    
