# knowledge base

List of strange issues I encountered that are not described in the package libs or books, and what I did to resolve, and-or identify. Ideally this list includes text I can find using text search.

# Corrupted console

-   If you wrote some code that uses termdash and you see the following corruption after you quite the app, or control + c ...

    ![docs/img/dontcloseconsole.png](docs/img/dontcloseconsole.png)

-   then the possible root causes to investigate are:

    1. Failure to properly use the control+c pattern. (todo, need to add to konsoleDSL)
    2. Not closing the terminal properly. Overriding or stopping the default built in 'q' to quit code in consoleDSL spike.
    3. Creating a new context and not re-using the supplied (returned) ctx.

# code lense 'run test' text bounces when typing

-   Not sure what the cause or fix it, but recently whenever I'm typing in VSCode inside a test file, whenever I press a key the code lense text moves up and down a pixel on each keypress so that it's incredibly visually distracting. It also happens whenever I change focus and open the editor.
    - tried removing all the custom gopls settings
    - tried restarting
    - tried rebooting
    - tried inserting linefeed between codelense text and next line.
    - Only fix (workaround) was to remove it totally, settings, go, enable code lense, [ ] if true, enable code lense for running and debugging tests. <-- it's a useless setting anyway since you have the green triangle in the margin.