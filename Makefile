NAME=TCPChat

PORT=$(firstword $(MAKECMDGOALS))

all: run
	
build:
	@echo "\033[1;32mBuilding $(NAME)...\033[0m"
	@go build -o $(NAME) .

run: build
	@echo "\033[1;32mRunning $(NAME) on port $(PORT)...\033[0m"
	@./$(NAME) $(PORT)

$(PORT): run

clean:
	@echo "\033[1;32mCleaning up...\033[0m"
	@rm -fr $(NAME)