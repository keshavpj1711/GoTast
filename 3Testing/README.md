# Tests in GO

## Writing Tests

Writing a test is just like writing a function, with a few rules

- It needs to be in a file with a name like **xxx_test.go**
- The test function must **start with the word Test**
- The test function **takes one argument only t *testing.T**
- To use the *testing.T type, you need to **import "testing"**, like we did with fmt in the other file


> For the sake of understanding we wrote the test after we wrote our program but from now on we will write tests first and then write our code

## A Cycle to be followed for good

- Write a test
- Make the compiler pass
- Run the test see if it fails and check the error message is meaningful
- Write enough code to make test pass
- Refractor(So that things are meaningful in your code)

> On the face of it this may seem tedious but sticking to the feedback loop is important.

> By not writing tests, you are committing to manually checking your code by running your software, which breaks your state of flow. You won't be saving yourself any time, especially in the long run.

# Things to Note

**Note:** Go source files can only have one package per directory. Make sure that your files are organised into their own packages.



# Resources 

To follow this you can go to: [Link](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world)

