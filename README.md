# Nibbler-Auth0

Connects Auth0 and Nibble.  The user extension connects Auth0 to our user model

## Configuration

AUTH0_CLIENT_ID [REQUIRED]
AUTH0_CLIENT_SECRET [REQUIRED]
AUTH0_CALLBACK_URL
AUTH0_DOMAIN [REQUIRED]
AUTH0_AUDIENCE 

## TODO

No timeout

No "is authed" function

# User Sample App

- first, change sample.application.go and set the user email to match the account in Auth0 (this module
requires the user to be in the local DB)
- should return 404 / "not found" when hitting localhost:<port>/test
- should redirect to login when hitting localhost:<port>/login
- once logged in, it redirects back to localhost:<port>/callback, it will respond with a 200 status code
- go to localhost:<port>/test - you should see "authorized"
