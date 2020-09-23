
HTTP Verbs
===
These are some conventions HTTP apis follow. These are actually not part of Rest specification. But we need to understand these to fully utilize Rest API.

HTTP defines a set of request methods to indicate the desired action to be performed for a given resource. Although they can also be nouns, these request methods are sometimes referred as HTTP verbs. Each of them implements a different semantic, but some common features are shared by a group of them: e.g. a request method can be safe, idempotent, or cacheable.

GET The GET method requests a representation of the specified resource. Requests using GETshould only retrieve data.

HEAD The HEAD method asks for a response identical to that of a GET request, but without the response body.

POST The POST method is used to submit an entity to the specified resource, often causing a change in state or side effects on the server.

PUT The PUT method replaces all current representations of the target resource with the request payload.

DELETE The DELETE method deletes the specified resource.

CONNECT The CONNECT method establishes a tunnel to the server identified by the target resource.

OPTIONS The OPTIONS method is used to describe the communication options for the target resource.

TRACE The TRACE method performs a message loop-back test along the path to the target resource.

PATCH The PATCH method is used to apply partial modifications to a resource.

THESE ARE ALL LIES.

Status Codes
===
## 1xx Information

* 100 Continue
* 101 Switching Protocols
* 102 Processing

## 2xx Success

* 200 OK
* 201 Created
* 202 Accepted
* 203 Non-authoritative Information
* 204 No Content
* 205 Reset Content
* 206 Partial Content
* 207 Multi-Status
* 208 Already Reported
* 226 IM Used

## 3xx Redirects

* 300 Multiple Choices
* 301 Moved Permanently
* 302 Found
* 303 See Other
* 304 Not Modified
* 305 Use Proxy
* 307 Temporary Redirect
* 308 Permanent Redirect

## 4xx Client Error

* 400 Bad Request
* 401 Unauthorized
* 402 Payment Required
* 403 Forbidden
* 404 Not Found
* 405 Method Not Allowed
* 406 Not Acceptable
* 407 Proxy Authentication Required
* 408 Request Timeout
* 409 Conflict
* 410 Gone
* 411 Length Required
* 412 Precondition Failed
* 413 Payload Too Large
* 414 Request-URI Too Long
* 415 Unsupported Media Type
* 416 Requested Range Not Satisfiable
* 417 Expectation Failed
* 418 I'm a teapot
* 421 Misdirected Request
* 422 Unprocessable Entity
* 423 Locked
* 424 Failed Dependency
* 426 Upgrade Required
* 428 Precondition Required
* 429 Too Many Requests
* 431 Request Header Fields Too Large
* 444 Connection Closed Without Response
* 451 Unavailable For Legal Reasons
* 499 Client Closed Request

## 5xx Server Error

* 500 Internal Server Error
* 501 Not Implemented
* 502 Bad Gateway
* 503 Service Unavailable
* 504 Gateway Timeout
* 505 HTTP Version Not Supported
* 506 Variant Also Negotiates
* 507 Insufficient Storage
* 508 Loop Detected
* 510 Not Extended
* 511 Network Authentication Required
* 599 Network Connect Timeout Error

ref : [https://dev.to/moficodes/build-your-first-rest-api-with-go-2gcj](https://dev.to/moficodes/build-your-first-rest-api-with-go-2gcj)