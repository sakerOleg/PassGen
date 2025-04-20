# PassGen
PassGen is a simple yet powerful command-line password generator written in Go. It creates secure, random passwords with customizable options to meet your security requirements.


Features

🛠 Generate random passwords with customizable length (default: 6 characters)

🔢 Option to include digits (0-9)

✨ Option to include special characters (!@#$%^&* etc.)

🔄 Generate multiple passwords at once

🚀 Lightweight and fast with no dependencies (except Cobra for CLI)


Usage: 
passgen [flags]


Flags:
  
  -d, --digits          Include digits in the generated password
  
  -h, --help            help for passgen
  
  -l, --length int      Length of the generated password (default 6)
  
  -r, --repeat int      Number of the generated passwords (default 1)
  
  -s, --specChars       Include special characters in the generated password
