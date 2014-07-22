go-mcaccutils-server
====================

go-mcaccutils-server is a server that exposes a REST API to fetch information about a minecraft account.

It supports fetching account details by name or UUID, although fetching a player by UUID is only supported when the player is already in the database. This is planned to be changed, but there's no easy way to go from a UUID to a username so until Mojang does something it's likely to stay as it is.

It also supports rate limiting by IP address.

Data is stored in a SQLite database, and it is also cached in memory to reduce disk operations.
