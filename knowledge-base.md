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

# syscall.CLONE_NEWUTS not exist when following 'containers from scratch'

When following Liz's video ↪ [containers from scratch, Liz Rice](https://www.youtube.com/watch?v=8fi7uSYlOdc) _GOTO 2018_

- problem is caused by missing gopls setting, see here for fix : https://github.com/lizrice/containers-from-scratch/issues/1


# No packages found for open file, underneath //go:build integration, and // +build integration

When adding build tags to the top of a file, receive the error: 

`No packages found for open file /Users/alanhemmings/src/go-workspace/src/github.com/goblinfactory/markdown/markdown_test.go: <nil>. If this file contains build tags, try adding "-tags=<build tag>" to your gopls "buildFlags" configuration (see (https://github.com/golang/tools/blob/master/gopls/doc/settings.md#buildflags-string). Otherwise, see the troubleshooting guidelines for help investigating (https://github.com/golang/tools/blob/master/gopls/doc/troubleshooting.md).go list`

fix : add `buildFlags []string` to settings.json, as follows. (add the entire 'gopls' section, specifically `buildFlags`)

```json
{
    "go.vetOnSave": "package",
    "gopls": {
        "build.buildFlags": [ "-tags"]
    }
}
```

# Cant debug integration tests in vscode - skipped [no tests to run]

fix : use environment variables to skip or run integration tests, example

```go

func TestPassingNoArgsShouldExitWith1(t *testing.T) {
	run := os.Getenv("integration")
	if run == "" {
		t.Skip("Set INTEGRATION to run this test.")
	}
	exitCode := 0
	exit = func(ec int) { exitCode = ec }
	os.Args = []string{}
	main()
	assert.Equal(t, exitCode, 1)
}
```
also, set the environment variables in `settings.json`

```json
{
    "go.vetOnSave": "package",
    "go.testEnvVars": {
        "integration": "yes"
    },
    "gopls": {
       "build.buildFlags": [ "-tags=integration"]
    },
}
```

- see more :  https://peter.bourgon.org/blog/2021/04/02/dont-use-build-tags-for-integration-tests.html
