# go-reloaded

The objective of this project is to help me learn about :

    1.  The Go file system(fs) API

    2.  String and numbers manipulation


# Project Overview

I am to build a command-line tool that reads a text file, applies a series of formatting rules, and writes the modified text to an output file. The tool is designed to be a simple text completion/editing/auto-correction utility.

# Key Features

1. Number Conversions

    . Convert hexadecimal numbers to decimal when followed by (hex)

         "1E (hex)" → "30"

    . Convert binary numbers to decimal when followed by (bin)

         "10 (bin)" → "2"

2. Case Transformations

    . Single word transformations:

         (up) → Uppercase: "go (up)" → "GO"

         (low) → Lowercase: "SHOUTING (low)" → "shouting"

         (cap) → Capitalized: "bridge (cap)" → "Bridge"

    . Multiple word transformations (with number parameter):

         (up, 2) → Uppercase the previous 2 words

         (low, 3) → Lowercase the previous 3 words

         (cap, 6) → Capitalize the previous 6 words

3. Punctuation Formatting

    . Attach punctuation marks to the preceding word

         "there ,and" → "there, and"

    . Handle punctuation groups properly

         "... boring" → "...boring" (no space after)

    . Format quotes correctly

         "' awesome '" → "'awesome'"

         Handle multiple words between quotes without spaces

4. Grammar Rules

    . Convert "a" to "an" before words starting with vowels (a, e, i, o, u) or 'h'

         "a amazing" → "an amazing"

         "a honest" → "an honest"


# Project Structure

     project/             
     ├── functions/          
     │   ├── functions_test.go  
     │   ├── functions.go        
     │
     ├── go.mod 
     │       
     ├── main.go            
     │   
     └── README.md
     │           
     ├── result.txt
     └── sample.txt

# Good Practices to Follow

    . Use meaningful variable/function names

    . Write modular, reusable code

    . Include error handling

    . Add comments for complex logic

    . Create unit tests

    . Use standard Go formatting (go fmt)


# Testing

     How to Run the Tests:

          Inside your Go module directory run:

               go test ./functions -v

