selfbuild:
	@go build -o clickup -v
	@mkdir -p ~/bin
	@mv -f faker ~/bin/faker
