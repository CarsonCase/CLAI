# CLAI
### AI in the CLI

CLAI is a program to inject Chat GPT directly into your linux command line! Here's a few examples of convinient uses:

### Answer Questions
```
clai what is the capital of France?
```

### Write Programs
```
clai code only. Write a python rock paper scissors game > rockPaperScissors.py
python rockPaperScissors.py
```

### Read Files
```
s
```

## Instalation
Currently clai is only for linux.

1. Install [Golang](https://go.dev/)
2. Clone this directory
3. Run `go build`
4. Run `make install`
5. Run `clai whatever prompt you want here. No need to put in quotes. All command line args will be spliced together`
6. Your first time running clai will prompt for your Open AI API key. You can get this [here](https://platform.openai.com/api-keys)
7. You're all set!