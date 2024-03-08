run: build
	./bin/main

run_late_restart: build
	@sleep 1.5
	./bin/main

build:
	@templ generate
	@tailwindcss build -o assets/css/tailwind.css --minify
	@go build -o ./bin/main
