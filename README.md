# url-shortener-go
URL shortener web application written in Go
### How to run
1. Clone the repository
 ##
    git clone https://github.com/shari4ov/url-shortener-go

2. To set domain name for the shortened URLs, change the value of `TINY_DOMAIN` in `docker-compose.yml` file
#
3. Run the application
##
    task run

3. Open postman and send a POST request to `http://localhost:8080` with the following body:
##
    {
        "url": "https://www.google.com"
    }

4. You will receive a response with the shortened URL: 
##
    {
        "url": "http://${TINY_DOMAIN}/1"
    }

5. Open the shortened URL in your browser and you will be redirected to the original URL
##
    http://${TINY_DOMAIN}/1 -> https://www.google.com
