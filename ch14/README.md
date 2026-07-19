# Exercises

1. Create a middleware-generating function that creates a context with a
   timeout. The function should have one parameter, which is the number of
   milliseconds that a request is allowed to run. It should return a
   `func(http.Handler)` `http.Handler`.

2. Write a program that adds randomly generated numbers between 0 (inclusive)
   and 100,000,000 (exclusive) together until one of two things happen: the
   number 1234 is generated or 2 seconds has passed. Print out the sum, the
   number of iterations, and the reason for ending (timeout or number reached).

3. Assume you have a simple logging function that looks like this:

    ```go
    import (
    	"context"
    	"fmt"
    )

    func Log(ctx context.Context, level Level, message string) {
    	var inLevel Level
    	// TODO get a logging level out of the context and assign it to inLevel
    	if level == Debug && inLevel == Debug {
    		fmt.Println(message)
    	}
    	if level == Info && (inLevel == Debug || inLevel == Info) {
    		fmt.Println(message)
    	}
    }
    ```

    Define a type called Level whose underlying type is string. Define two
    constants of this type, `Debug` and `Info`, and set them to `"debug"` and
    `"info"`, respectively. Create functions to store the log level in the
    context and to extract it. Create a middleware function to get the logging
    level from a query parameter called `log_level`. The valid values for
    `log_level` are `debug` and `info`.

    Finally, fill in the TODO in Log to properly extract the log level from the
    context. If the log level is not assigned or is not a valid value, nothing
    should be printed.
