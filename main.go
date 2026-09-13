package main

func main() {
	config := &config{
		commandList: getCommands(),
	}
	startRepl(config)
}
