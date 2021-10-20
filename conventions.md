# Go conventions

1.  interfaces
    -   interfaces are usually named with `er` endings. e.g. `Reader`, `Closer`, `Heater`, `Runner` `json.Marshaler`, `http.Handler`
    -   accept `interfaces`, return `structs`
1.  indenting
    -   indent error handling code,
    -   indent exception case
    -   Do not indent main logic flow
1.  return values
    -   `Comma` `ok` idiom. (return value + command + a boolean named ok)
1.  type switching
    -   assign variable being switch to same name to shadow variable with more strongly typed object.
