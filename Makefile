MAIN=main.go
EXE=TCPChat

PORT=8989

all: run
	
build:
	@echo "\033[1;32mBuilding $(MAIN) into $(EXE)...\033[0m"
	@go build -o $(EXE)


run: build
	@echo "\033[1;32mRunning $(EXE)...\033[0m"
	@./$(EXE) $(PORT)


clean:
	@echo "\033[1;32mCleaning up...\033[0m"
	@rm -f $(EXE)
	@clear