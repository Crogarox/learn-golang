 Go has only one looping construct, the for loop. The basic for loop has three components separated by semicolons: 

    the init statement: executed before the first iteration
    the condition expression: evaluated before every iteration
    the post statement: executed at the end of every iteration

    for i := 0; i != 10; i++ {

    }

 The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement. 
The loop will stop iterating once the boolean condition evaluates to false. 
Note: Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required. 

- The init and post statements are optional. 

    i := 0
  for ;i != 10; {
      i++
    }

- and you can drop the semicolons, C's while is spelled for in Go.
    i := 0
  for i != 10 {
    i++ 
    }


- Omit the loop condition it loops forever, so an infinite loop is compactly expressed.

for {
    
}