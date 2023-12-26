# Log Collection API

## Table of Contents

- [Overview](#overview)
- [Data Types](#data-types)
- [Endpoints](#endpoints)
- [Resources](#resources)
<<<<<<< HEAD
- [Todo](#todo)

### Data Types

=======

## TODO
[ ] creating new certificates
[ ] storing the new certificates in a database, with a unique-user-id and certificate value
[ ] storing the UUID and encrypted private key cooresponding to the certificate

### Data Types
- User kj
>>>>>>> 688ff65a35a3205257ceddfa21b47155d53e99d0

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


### Resources
[HTTP Status Codes](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#client_error_responses)

## TODO
- [ ] creating new certificates
- [ ] storing the new certificates in a database, with a unique-user-id and certificate value
- [ ] storing the UUID and encrypted private key corresponding to the certificate

