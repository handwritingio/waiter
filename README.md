# waiter

This tool waits for a service to begin accepting TCP connections. It is designed
to be used as a blocking call before running another command, e.g.

    waiter postgres://user:pw@hostname:5432/db && ./start_your_app

Or simply

    waiter hostname:5432 && ./start_your_app

The primary use case is in starting up services when using
[docker-compose](https://www.docker.com/docker-compose).

https://github.com/docker/compose/issues/374 shows a lot of people dealing with
the same problem. The solution used here is essentially based on the one shown
in https://github.com/aanand/docker-wait. Unfortunately, the author explains
that it doesn't work with docker-compose:
https://github.com/docker/compose/issues/374#issuecomment-69212755
