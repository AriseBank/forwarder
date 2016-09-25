all:
	rice embed-go
	go install github.com/ssddanbrown/forwarder

dev:
	go install github.com/ssddanbrown/forwarder