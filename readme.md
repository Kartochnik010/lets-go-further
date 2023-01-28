# docs

In order to use the API you will need to sign up using your **name**, **email**, and **password**.

```jsx
curl \
	-d '{"name":"Ilyas","email": "ilyas@example.com", "password": "pa55word"}' \
	-X POST \
localhost:4000/v1/users
```

At this stage you should be asked to activate your account before you can proceed into your 

account.

After registration an email with activation token will be sent to your account. Put the token in JSON body and do dis:

```jsx
curl \
	-X PUT \
	-d '{"token": "DWNGVAKBS724QCQXOOQR5EE7PE"}' \
localhost:4000/v1/users/activated
```

Once you activate your account, you will be redirected to the front page OR your profile page for further configuration.

When redirected, you should already have the *authentication token.* You can generate it by

```jsx
curl 
	-d '{"email": "ilyas@example.com", "password": "pa55word"}' 
localhost:4000/v1/tokens/authentication
```

We can also pre-generate *authentication token* and put it inside activation link, OOOOR you will be redirected and asked again to enter your credentials, if correct only then the *authentication token* will be generated. Now you can proceed. 

You should see something like this:

```jsx
{
  "authentication_token": {
    "token": "RVHNIY3GGMMWBRP35IISJ3U6B4",
    "expiry": "2023-01-30T03:29:48.149241+06:00"
  }
}
```

You will need to include the token in the *Bearer* header every time you use a protected in order to be authenticated. The server *****could***** save that as a session knowing every link you clicked.

```jsx
curl \
	-H "Authorization: Bearer RVHNIY3GGMMWBRP35IISJ3U6B4" \
localhost:4000/v1/movies
```

cool right?

```jsx
{
  "metadata": {},
  "movies": []
}
```