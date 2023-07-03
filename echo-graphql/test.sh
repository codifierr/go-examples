curl --location 'http://localhost:8080/graphql' \
--header 'Content-Type: application/json' \
--data '{"query":"{ user(id: \"1\")\n\n{ id, name }\n\n}","variables":{}}'