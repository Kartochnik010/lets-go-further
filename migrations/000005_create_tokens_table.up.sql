CREATE TABLE IF NOT EXISTS tokens (
    hash bytea PRIMARY KEY,
    user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE, 
    expiry timestamp(0) with time zone NOT NULL,
    scope text NOT NULL
);


-- The hash column will contain a SHA-256 hash of the activation token. It’s important to emphasize that we will only store a hash of the activation token in our database — not the activation token itself.

-- We want to hash the token before storing it for the same reason that we bcrypt a user’s password — it provides an extra layer of protection if the database is ever compromised or leaked. Because our activation token is going to be a high-entropy random string (128 bits) — rather than something low entropy like a typical user password — it is sufficient to use a fast algorithm like SHA-256 to create the hash, instead of a slow algorithm like bcrypt.

-- The user_id column will contain the ID of the user associated with the token. We use the

-- REFERENCES user syntax to create a foreign key constraint against the primary key of our users table, which ensures that any value in the user_id column has a corresponding id entry in our users table.

-- We also use the ON DELETE CASCADE syntax to instruct PostgreSQL to automatically delete all records for a user in our tokens table when the parent record in the users table is deleted.

-- Note: A common alternative to ON DELETE CASCADE is ON DELETE RESTRICT , which in our case would prevent a parent record in the users table from being deleted if theuserhasanytokensinourtokenstable.IfyouuseON DELETE RESTRICT,you would need to manually delete any tokens for the user before you delete the user record itself.

-- The expiry column will contain the time that we consider a token to be ‘expired’ and no longer valid. Setting a short expiry time is good from a security point-of-view because it helps reduce the window of possibility for a successful brute-force attack against the token. And it also helps in the scenario where the user is sent a token but doesn’t use it, and their email account is compromised at a later time. By setting a short time limit, it reduces the time window that the compromised token could be used.

-- Of course, the security risks here need to be weighed up against usability, and we want the expiry time to be long enough for a user to be able to activate the account at their leisure. In our case, we’ll set the expiry time for our activation tokens to 3 days from the moment the token was created.

-- Lastly, the scope column will denote what purpose the token can be used for. Later in the book we’ll also need to create and store authentication tokens, and most of the code and storage requirements for these is exactly the same as for our activation tokens. So instead of creating separate tables (and the code to interact with them), we’ll store them in one table with a value in the scope column to restrict the purpose that the token can be used for.