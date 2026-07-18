# Exercises

1. Create a sentinel error to represent an invalid ID. In main, use `errors.Is`
   to check for the sentinel error, and print a message when it is found.

2. Define a custom error type to represent an empty field error. This error
   should include the name of the empty `Employee` field. In `main`, use
   `errors.As` to check for this error. Print out a message that includes the
   field name.

3. Rather than returning the first error found, return back a single error that
   contains all errors discovered during validation. Update the code in `main`
   to properly report multiple errors.
