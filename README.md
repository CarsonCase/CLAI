# CLAI
### AI in the CLI

CLAI is a program to inject Chat GPT directly into your linux command line! Now with some new convinient features!

## Commands
- `ask` asks a question using all the following args as a prompt and prints result to standard output
- `chat` asks a question saving the prompt and response to a local csv file for future context
- `read-ask` reads a file and postpends the prompt!
- `reset-chat` resets the chat data
- `set-model` allows you to set the openAI model you use, defaults to `gpt-4o`

## Example Usage
### Answer Questions
```
clai ask what is the capital of France?
```

### Write Programs
```
clai chat code only. Write a python rock paper scissors game > rockPaperScissors.py
python rockPaperScissors.py
```

### Read Files
```
clai read-ask rockPaperScissors.py rewrite this python code into javascript > rockPaperScissors2.js
```

## Instalation
Currently clai is only built for linux (debian). Golang can create executables for any OS however, you may need to adjust file paths for AI keys and chat history

1. Install [Golang](https://go.dev/)
2. Clone this directory
3. Run `go build`
4. Run `make install`
5. Run `clai ask whatever prompt you want here. No need to put in quotes. All command line args will be spliced together`
6. Your first time running clai will prompt for your Open AI API key. You can get this [here](https://platform.openai.com/api-keys)
7. You're all set!