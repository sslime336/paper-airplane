.DEFAULT_GOAL=paper-airplane

BUILD=paper-airplane.exe

paper-airplane:
	go build -o $(BUILD) .

run: paper-airplane
	./$(BUILD)

release-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o build/$(BUILD)

clean:
	rm -rf ./log
	rm -rf ./build
	rm -rf $(BUILD)

.PHONY: run paper-airplane