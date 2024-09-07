run-client:
	cd client && pnpm run dev

run-server:
	cd server && go run main.go

dev:
	make run-server & make run-client