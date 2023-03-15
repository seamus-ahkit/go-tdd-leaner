https://quii.gitbook.io/learn-go-with-tests/
Expand on Go knowledge with a TDD project

go test -v
    Prints the full output of all tests even if they pass

Writing test basics
    Writing a test is just like writing a function, with a few rules
    - It needs to be in a file with a name like xxx_test.go
    - The test function must start with the word Test
    - The test function takes one argument only t *testing.T
    - In order to use the *testing.T type, you need to import "testing"

Discipline (feedback loop for TDD)
    - Write a test
    - Make the compiler pass
    - Run the test, see that it fails and check the error message is meaningful
    - Write enough code to make the test pass
    - Refactor

Helper functions
    For helper functions, it's a good idea to accept a testing.TB which is an interface that *testing.T and *testing.B both satisfy, so you can call helper functions from a test, or a benchmark.

    t.Helper() is needed to tell the test suite that this method is a helper. By doing this when it fails the line number reported will be in our function call rather than inside our test helper.

Examples
    Often code examples that can be found outside the codebase, such as a readme file often become out of date and incorrect compared to the actual code because they don't get checked.

    Go examples are executed just like tests so you can be confident examples reflect what the code actually does. We look at the examples as we write and update tests so they are fresh in our mind

    You need to include the comment with output in the example function to execute the example 
        e.g., // Output: 6
