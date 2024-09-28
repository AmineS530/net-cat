EXE=TCPChat

PORT=$(firstword $(MAKECMDGOALS))

all: run
	
build:
	@echo "\033[1;32mBuilding $(EXE)...\033[0m"
	@go build -o $(EXE) .

run: build
	@echo "\033[1;32mRunning $(EXE) on port $(PORT)...\033[0m"
	@./$(EXE) $(PORT)

$(PORT): build
	@echo "\033[1;32mRunning $(EXE) on port $(PORT)...\033[0m"
	@./$(EXE) $(PORT)

clean:
	@echo "\033[1;32mCleaning up...\033[0m"
	@rm -f $(EXE)
	@clear

# Prevent make from complaining about a missing target when using arguments like 8080#%:
# #@:
