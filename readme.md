# go sandbox

Collection of small simple experiments in different go language topics.

### language basics

-   structs, maps and arrays

    -   [structs and maps](pkg/structsandmaps/structsandmaps.go)
    -   [no matrixes](pkg/arrs/matrix.go)
    -   [sorting](pkg/arrs/sorting.go)

-   channels

    -   [simple signal channels](pkg/muxyidiomatic/muxyidiomatic-signals.go)
    -   [signal channels together with waitgroup](pkg/muxyidiomatic/muxyidiomatic-waitgroup.go)

    -   [channels for service status](pkg/channels/channels-for-service-status.go)
    -   [magic numbers in channels](pkg/channels/magic-number-channel.go)

-   waitgroup

    -   [test waitgroup](pkg/testwaitgroup/testwaitgroup.go)

-   switch

    -   [fizz buzz switch example](pkg/switchy/switchy.go)

-   timing

    -   [timing example](pkg/timing/timing.go)

-   error handling

    -   [typical error handling](pkg/errorhandling/errorhandling.go)
    -   [custom error objects](pkg/errorhandling/custom-errors.go)

-   dependency injection

    -   [http server example from 'learning Go'](pkg/dependencyinjection/main.go)

### switching from C#

-   linq

    -   [no linq](pkg/nolinq/nolinq.og)

-   software patterns

    -   [circuit breaker example from 'cloud native go'](pkg/bloggy/breaker/breaker.go)
    -   [sample quote api using rate limiter return 429 if requests too fast](pkg/bloggy/quoteapi/quoteapi.go)

### experimental

Random spike projects, to test to see what it might take to write something yourself, without using a recommended package.

-   currency, money and hand rolled decimal-like types

    -   [test account service](pkg/testaccoiuntservice/testaccountservice.go)
