# go-tdd-leaner
Expand on Go knowledge with a TDD project

Writing tests
    Writing a test is just like writing a function, with a few rules
    - It needs to be in a file with a name like xxx_test.go
    - The test function must start with the word Test
    - The test function takes one argument only t *testing.T
    - In order to use the *testing.T type, you need to import "testing"

*testing.T is the hook into the testing framework

t.Errorf used for printing failed test error

