# Quick ssh bandit
A lightweight script to start the ssh connection to play the **bandit** wargame on [overthewire](https://overthewire.org/wargames/). 

## Usage
Build main.go using the command:
```shell
go build -o quick-ssh-bandit main.go
```
You can either add the executable in the PATH or run it by it's relative path.
```shell
quick-ssh-bandit bandit_id
```
where the *bandit_id* is the level.