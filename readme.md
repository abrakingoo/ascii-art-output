
## Ascii-art-Ouput
Ascii-art-output is a program which consists of receiving an output filename, string and a banner file as arguments and writes the string in a graphic representation using ASCII to the output file.

What we mean by a graphic representation using ASCII, is to write the string received using ASCII characters, as you can see in the example below:


### Features

- The file must be named by using the flag `--output=<fileName.txt>`, in which --output is the flag and` <fileName.txt>` is the file name which will contain the output.
- The program writes to a file each character of the input string line by line to form the ASCII Art.



#### File names
- Only File names that can be used to run the program are:

 1. standard
 2. thinkertoy
 3. shadow

 ### Error Handling
 1. Invalid number of arguments:

    - Checks if the number of arguments is not equal to 4 when ASCII art switch is on.
    - Prints usage instructions and exits if the condition is met.

 2. Restrictions on banner directory:

    - Checks if the banner contains "bannerfiles/", which is restricted.
    - Prints a message and exits if the condition is met.

3. Invalid character in banner:

    - Checks if the banner contains "/" character, allowing only file names.
    - Prints a message and exits if the condition is met.

4. Incorrect usage of --output flag:

    - Checks if the first argument does not contain "--output=" when ASCII art switch is on.
    - Prints usage instructions and exits if the condition is met.

5. Incorrect file extension:

    - Checks if the first argument does not have a suffix of ".txt" when ASCII art switch is on.
    - Prints a message and exits if the condition is met.

6. File reading error:

    - Attempts to read the ASCII art text file and checks for any errors.
    - Prints the error and exits if encountered.

7. Handling empty files:

    - Checks if the length of the file content is not equal to the expected length of 856.
    - Prints an error message and exits if the condition is met.

8. Handling newlines and backspaces:

    - Checks for backspace characters and replaces them.
    - Handles newline characters by splitting the input string.
    - Calls ascii.ErrHandler to handle any errors from newline processing.
 
 ### Running the project
To run the project open bash terminal and run.

```bash
git clone https://learn.zone01kisumu.ke/git/aaochieng/ascii-art-output
cd ascii-art-output
```
<a id ="sec"></a>
```go
Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard
```

The above command will write the below ascii art to the output filename
```bash
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$
``` 
### Error Handling

- The program has error handling for the following scenarios:
- If the standard.txt file is empty.
- If there are any issues while reading the standard.txt file.
- If there are any issues while handling newline characters in the input string.
- If number of arguments are not equal to four.
- If the the flag and file extetion does not match [ Format](#sec)


#### Note

This program is a simple implementation of ASCII Art output and might need to be adjusted based on your specific use case.

### Contibution
 - Ascii-art-output is an open source project and we welcome contributions.
 - Feel free to fork and create pull requests 

 #### Commit Message Structure

A commit message for contribution should consists of three distinct parts separated by a blank line: the title, an optional body and an optional footer. The layout looks like this:
```
type: Subject

body

footer
```

#### The Type

The type is contained within the title and can be one of these types:
- feat: A new feature
- fix: A bug fix
- docs: Changes to documentation
- style: Formatting, missing semi colons, etc; no code change
- refactor: Refactoring production code
- test: Adding tests, refactoring test; no production code change
- chore: Updating build tasks, package manager configs, etc; no production code change

#### The Subject

- Subjects should be no greater than 50 characters, should begin with a capital letter and do not end with a period.

- Use an imperative tone to describe what a commit does, rather than what it did. For example, use change; not changed or changes.
#### The Body

- Not all commits are complex enough to warrant a body, therefore it is optional and only used when a commit requires a bit of explanation and context. Use the body to explain the what and why of a commit, not the how.

- When writing a body, the blank line between the title and the body is required and you should limit the length of each line to no more than 72 characters.
#### The Footer

- The footer is optional and is used to reference issue tracker IDs
### Contributors

<table>
<tr>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://learn.zone01kisumu.ke/git/aaochieng>
            <img src=https://learn.zone01kisumu.ke/git/avatars/8a1b24358854eb12998a07c269542193?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Aaron/>
            <br />
            <sub style="font-size:14px"><b>Aoron Ochieng</b></sub>
        </a>
    </td>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://learn.zone01kisumu.ke/git/ebarsula>
            <img src=https://learn.zone01kisumu.ke/git/avatars/fa966ef34b0ccdfe772414745aeee49f?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Emmanuel/>
            <br />
            <sub style="font-size:14px"><b>Emmanuel Barsulai</b></sub>
        </a>
    </td>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://learn.zone01kisumu.ke/git/abrakingoo>
            <img src=https://learn.zone01kisumu.ke/git/avatars/c307852c0cb9222c1ea2c71f98ff2d51?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Abraham/>
            <br />
            <sub style="font-size:14px"><b>Abraham kingoo</b></sub>
        </a>
    </td>
</tr>
</table>