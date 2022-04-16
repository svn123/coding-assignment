# coding-assignment

Detailed explanation for the coding assignment
1. 2D slice of byte is initialized to store the character. 
2. horizonal search function will search for the word in left to right direction only, if not found returns false, else returns the co-ordinates of the first letter and count of rotation. 
3. vertical search function will search for the word in the top to bottom direction only, if not found returns false, else returns the co-ordinates of the first letter and count of rotation. 
4. rotate function is defined to rotate the characters in clockwise by 1 step, if the word is not found in either horizontal search or vertical search, then a rotation is done to the array and the search for the word will continue.
5. All the functions have been unit tested with a code coverage of 93.5 %


Pre-requiste to run the program
1. Have golang installed in the system
2. Clone the repository [coding-assignment](https://github.com/svn123/coding-assignment.git)


Commands to run and test the program

Command to run program:
```
$ cd golang_assignment
$ go run main.go
```

Command to test the program and run unit tests:
```
$ go test -v
```

Command to check the code coverage:
```
$ go test -coverprofile coverage.out
```




