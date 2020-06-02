# trainee-uniq

## About

**trainee-uniq** is a from scratch developed program which outputs sorted text files without double lines . There are
various parameters you can use to filter or individualize the output of your file.
Seen like this it's an imitation of the original uniq command line command.

## Parameters

| Parameter    | Usage | 
| ------------- |-------------|
| **-u**      | Allows you to print only unique lines.
| **-i**      | Ignore differences in case when comparing. 
| **-d**  | Only print duplicate lines, one for each group.    
| **-c**     | Prefix lines by the number of occurrences.

## Usage Examples

In the following you'll see a quick example of how to use the uniq command. Here is what our text file looks like:

toUniq.txt:

               Dog
               Cat
               Mouse
               Mouse
               Mouse
               Cat
               Dog
               Dog
               Mouse
               Mouse
               
Now we can use the uniq command with `./uniq.linux-amd64 < toUniq.txt `  

The output is:

                Dog
                Cat
                Mouse
                Cat
                Dog
                Mouse
               
It's also possible to use the parameters shown above. By using parameters the output is individualized as described.

**u** - parameter: `./uniq.linux-amd64 -u < toUniq.txt ` 

The output is:

                Dog
                Cat
                Cat

**d** - parameter: `./uniq.linux-amd64 -d < toUniq.txt `

The output is:

                Dog
                Cat
                Cat

**i** - parameter: `./uniq.linux-amd64 -i < toUniq.txt `

The output is:

                DOG
                CAT
                MOUSE
                CAT
                DOG
                MOUSE

**c** - parameter: `./uniq.linux-amd64 -c < toUniq.txt `

The output is:

                      1 Dog
                      1 Cat
                      3 Mouse
                      1 Cat
                      2 Dog
                      2 Mouse

There is also the possibility to combine parameters.

Example: `./cut.linux-amd64 -u -c -i < toCut.txt ` 

The output is:

                      1 DOG
                      1 CAT
                      1 CAT
