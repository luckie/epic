# Epic Content Microservice
##### A minimalist content management system (CMS) API that persists content as JSON to a database.  It is scalable, device-agnostic, extremely performant, and does not impose any architectural constraints upon the user experience that is using it.  Epic is a Go language program that is compiled into a platform-specific native binary, and usually deployed as a Docker container.  It automatically deploys using no-cost zero-administration [Let's Encrypt](https://letsencrypt.org) TLS certificates (SSL).

##### Epic is not for everyone - especially not for fans of WordPress, Drupal, or Joomla.  Epic is for technically-savvy developers who can not express their creativity, while being constrained by traditional CMS limitations and tradeoffs.  99% of those using a traditional CMS should not even consider using Epic, because Epic assumes a robust modern client-side development skill set.

##### Epic addresses a very specific niche that other CMS have not successfully fulfilled.  The inception of Epic was motivated by the need for a fully-functional open source CMS designed from the ground up to drive highly-customized ultra-rich user experiences that are delivered to end users via modern JavaScript frameworks (e.g. Ember, Angular, React) and native mobile apps (e.g. iOS and Android).  The CMS would have to play nice with unrelated systems supported by the client - never getting in the way; never forcing unrelated systems to do things "the Epic way".

##### The Epic API is benefiting from rapid development, and new features are being released regularly.  Accordingly, it is subject to change without warning.  Once it stabilizes, the Epic API will follow semantic versioning.

**The Epic database connection string must be included as the first parameter.**
e.g. ```postgres://username:password@localhost/epic?sslmode=disable```

**The Epic host URL string must be included as the second parameter.**
e.g. ```example.com``` or ```epic.example.com```

A port number may optionally be included as the third parameter for HTTP (not HTTPS).
If the port number is not specified, the server will default to HTTPS on port 443.
If a port number other than 443 is specified, the server will use HTTP on the specified port.
Any port number below 1025 requires administrative privileges (sudo).

**Here is an example of launching Epic using TLS (HTTPS):**
```
sudo epic postgres://username:password@localhost/epic?sslmode=disable example.com
```

**Here is an example of launching Epic without TLS (HTTP) on port 8080:**
```
epic postgres://username:password@localhost/epic?sslmode=disable example.com 8080
```

## Database Setup

```
CREATE USER epic WITH PASSWORD 'epic';
CREATE DATABASE epic OWNER epic WITH ENCODING='UTF8';
CREATE SCHEMA epic;
GRANT ALL PRIVILEGES ON DATABASE "epic" to epic;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA epic TO epic;
```

## Content Management

### Read Content
HTTP GET request to ""/content/{id}"

### Create Content
Requires authorization
HTTP POST request with JSON body to "/content"
```
{
  "app-id":"123e4567-e89b-12d3-a456-426655440000",
  "name":"Left side-bar, bottom",
  "description":"Text that appears at the bottom of the left side-bar."
}
```
Responds with an HTTP Response Code and JSON body that either has the new Content ID or an error message.
```
{
  "id":"123e4567-e89b-12d3-a456-426655440000"
}
```

### Update Content
Requires authorization
HTTP PUT request with JSON body to "/content/{id}"
```
{
  "id":"123e4567-e89b-12d3-a456-426655440000",
  "locale":"us-EN",
  "data": {
    "id":"123",
    "your-data":"whatever-you-want"
  }
}
```

- Only top-level fields are used by Epic.
- The top-level id must be a UUID V4.
- The top-level 'locale' must be a valid locale string.
- The top-level 'tags' must be valid JSON consisting of an array of strings.
- Any valid JSON can go into the top-level 'data' field.  Epic stores 'data' exactly as is, and does not use it use.  It is your content.
- Relationships between 'content' records are established by applying one or more 'tags'.

### Read All Content Associated With A Specific Tag
HTTP GET request to "/app/{app-uuid}/tag/{tag}"
The app-uuid is a standard version 4 UUID.

### Create Tag
Requires authorization.
HTTP POST request to "/app/{app-uuid}/tag/{tag}".
The app-uuid is a standard version 4 UUID.

### Assign Tag To Content
Requires authorization,
HTTP POST request to "/content/{content-uuid}/tag/{tag}".

### Create Presigned URL For AWS S3 PUT Operation
Requires authorization.
Requires AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environmental variables with appropriate values.
HTTP POST request with JSON body to "/asset/url".
```
{
  "bucket":"epic-content-assets",
  "key":"asset-key-for-s3"
}
```
Responds with an HTTP Response Code and JSON body.
```
{
  "bucket":"epic-content-assets",
  "key":"asset-key-for-s3",
  "url":"https://epic-content-assets.s3.amazonaws.com/asset-key-for-s3?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAIMTULOINQ3C24OUQ%2F20160517%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20160517T123628Z&X-Amz-Expires=900&X-Amz-SignedHeaders=host&X-Amz-Signature=926629e70dc3cc777943b500975d5764a4824aac8d40fc63d81cda8a9d6f733c",
  "error":""
}
```

## User Management

### Login
HTTP POST request with JSON body to "/auth/login".
Responds with an HTTP Response Code, and if authentication is successful, a JSON Web Token (JWT) in the HTTP 'Authorization' header.  All subsequent HTTP Requests should include that JWT in the 'Authorization' header.
```
{
  "username":"gandalf",
  "password":"mellon",
  "app-id":"123e4567-e89b-12d3-a456-426655440000"
}
```

### Logout
Requires authorization.
HTTP DELETE request to "/auth/logout".

### Create User
Requires authorization.
HTTP POST request with JSON body to "/auth/user".
The JSON 'id' field is optional.  It will be generated by the API, if the client does not provide it.
```
{
  "id":"123e4567-e89b-12d3-a456-426655440000",
  "first-name":"Gandalf",
  "last-name":"Stormcrow",
  "email":"olórin@valinor.gov",
  "username":"gandalf",
  "password":"mellon",
  "app-id":"123e4567-e89b-12d3-a456-426655440000"
}
```
Responds with an HTTP Response Code and JSON body.

### Get User / Users
Requires authorization.
HTTP GET request to "/auth/user".

### Delete User
Requires authorization.
HTTP DELETE request to "/auth/user" to delete your own user account.
HTTP DELETE request to "/auth/user?id=123e4567-e89b-12d3-a456-426655440000" to delete another user's account.

### Update Password
Requires authorization.
HTTP PUT request with JSON body to "/auth/user/password".
The JSON 'password' field is required, and 'id' field is optional (allowing an admin to change another user's password).
```
{
  "password":"mellon"
}
```
or optionally
```
{
  "id":"123e4567-e89b-12d3-a456-426655440000",
  "password":"mellon"
}
```
Responds with an HTTP Response Code, and an empty JSON body (since it's a PUT request).

### Authenticate Token
Requires authorization.
HTTP POST request to "/auth/token".
Responds with an HTTP Response Code.

### New UUID
Requires authorization.
HTTP GET request to "/auth/uuid".
Responds with an HTTP Response Code, and a new version 4 UUID in the JSON body.
```
{
  "uuid":"123e4567-e89b-12d3-a456-426655440000"
}
```

### Bootstrap Crypto For Initial Admin User
Does not persist its output for security reasons.  Just responds with JSON that can be manually input into database.
HTTP POST request with plain-text password in JSON body to "/auth/crypto".
```
{
  "plain-text":"mellon"
}
```
Responds with an HTTP Response Code and JSON body that includes the plain-text password, hashed password, salt, private key, and public key.
```
{
  "plain-text":"mellon",
  "hash": "zVfyekeU8ZQyvhaV/ESZPNfZZtBmCjQOMORWB42Kf1k=",
  "salt": "TJXaDWu3aSshT9PCxzShrk71QlbYVciVuAKRMQD1gYuRzWq4O1uC30RxR+P65/aBKpkNwNMTX/NWG4oRxM9kXw==",
  "private-key": "&{{c82000a340 a1b92c9bc85cdcf3d4691240db94913e192d7892faa808439f35cb975addef5251741dcfc1e013b696017dd14e835246 8be42e5c4a13fbec6797ec6c2822cc572a86264056fe38a08fb3898ce63bc021c317442218e119f0d2a4fab6b7b39162} 980f075a93c83ea6f897236c018e6965f1ea8cf75c60da3864c8f0903413cf06fc091803f7101f78cc5943df8754b8ee}",
  "public-key": "{c82000a340 a1b92c9bc85cdcf3d4691240db94913e192d7892faa808439f35cb975addef5251741dcfc1e013b696017dd14e835246 8be42e5c4a13fbec6797ec6c2822cc572a86264056fe38a08fb3898ce63bc021c317442218e119f0d2a4fab6b7b39162}"
}
```

**Any API call that includes "Requires authorization." in the documentation requires that the HTTP Request have a valid JSON Web Token (JWT) in the HTTP 'Authorization' header.**
