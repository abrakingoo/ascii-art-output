
## Ascii-art-Ouput
Ascii-art-output is a program which consists of receiving an output filename, string and a banner file as arguments and writes the string in a graphic representation using ASCII to the output file.

What we mean by a graphic representation using ASCII, is to write the string received using ASCII characters, as you can see in the example below:


### Features

- The file must be named by using the flag `--output=<fileName.txt>`, in which --output is the flag and` <fileName.txt>` is the file name which will contain the output.
- The program writes to a file each character of the input string line by line to form the ASCII Art.

### Running the project

#### File names
- Only File names that can be used to run the program are:

 1. standard
 2. thinkertoy
 3. shadow
 
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
- It the the flag and file extetion does not match [ Format](#sec)


#### Note

This program is a simple implementation of ASCII Art output and might need to be adjusted based on your specific use case.

### Contibution
 - Ascii-art-output is an open source project and we welcome contributions.
 - Feel free to fork and create pull requests 

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