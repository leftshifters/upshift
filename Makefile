release:
	env GOOS=linux GOARCH=amd64 go build -o upshift-linux -v upshift
	env GOOS=darwin go build -o upshift-darwin -v upshift