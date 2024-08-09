.DEFAULT_GOAL=bot

BUILD=paper-airplane
AIRP_ENTRY="./cmd"

bot:
	go build -o $(BUILD) $(AIRP_ENTRY)

run: bot
	./$(BUILD)

gen:
	go run ./cmd/gen/gen.go

release-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o build/$(BUILD) $(AIRP_ENTRY)

clean:
	rm -rf ./log
	rm -rf ./build
	rm -rf $(BUILD)

.PHONY: bot
