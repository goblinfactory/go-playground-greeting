# knowledge base

List of strange issues I encountered that are not described in the package libs or books, and what I did to resolve, and-or identify. Ideally this list includes text I can find using text search.

# Corrupted console

-   If you wrote some code that uses termdash and you see the following corruption after you quite the app, or control + c ...

    ![docs/img/dontcloseconsole.png](docs/img/dontcloseconsole.png)

-   then the possible root causes to investigate are:

    1. failure to properly use the control+c pattern. (to add to konsoleDSL)

and the fix is: