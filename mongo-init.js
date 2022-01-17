db.createUser(
    {
        user: "jamal",
        pwd: "jamal",
        roles: [
            {
                role: "readWrite",
                db: "raya-db"
            }
        ]
    }
);