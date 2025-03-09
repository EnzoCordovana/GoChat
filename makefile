.PHONY: back front

back:
	cd backend && go run .

front:
	cd frontend && pnpm run dev