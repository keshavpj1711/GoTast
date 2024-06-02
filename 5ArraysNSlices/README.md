# Arrays and Slices 

As I already have knowledge about arrays and slices so i am not going to discuss about it here.
But there are a few things that should be kept in mind:

- Slices cant be compared directly, so if you wish to compare slices few workarounds are:
    - Write a function that iterates over each and every element
    - Another way is through `reflect.DeepEqual()` which basically compares if two variables are equal or not.
- 