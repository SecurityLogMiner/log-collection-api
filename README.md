# Log Collection API

## Table of Contents

- [Overview](#overview)
- [Data Types](#data-types)
- [Endpoints](#endpoints)
- [Testing](#testing)
- [Resources](#resources)
- [Todo](#todo)

### Overview
A user first visits the site. 
A prerequisite to accessing the api is the user have an email associated with an
Auth0 tenant.

Then, from their dashboard, they make their requests using a valid JWT.

### Data Types

### Endpoints
```
POST /api/download/cert
The request retrieves an existing x509 cert, if any, from the database that has
been previously allocated to the user.

Request JSON: 
--body--
|**Key**|**Description**|
|*uuid*|The unique user id used to retrieve the existing x509 cert.|

Response JSON:
--200 Success--
|**Key**|**Value**|
|**|.|

--404 Failure--
|**Key**|**Value**|
|*Error*| Unable to locate requested resource.|
```

<a href="#table-of-contents" style="font-size:smaller;">back to top</a>

### Testing
Must have the server running before testing. Then:
```
go run ./tests/scripts/tests.go
```

<a href="#table-of-contents" style="font-size:smaller;">back to top</a>


### Resources
- [HTTP Status Codes](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#client_error_responses)
- [API Example](https://auth0.com/docs/quickstart/backend/golang)


## TODO
- [ ] launch a database connection in the server repo
- [ ] setup the tables for:
    1. email:uuid       string:string
    2. uuid:cert        string:string
        - used for the TLS(client > server) connection once software starts
    3. uuid:ips         string:[...,string]
- [ ] creating new certificates


