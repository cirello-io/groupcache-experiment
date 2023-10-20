workdir: .
observe: *.go
build-all: make all
backend: ./backend
frontend-1:./frontend -listen "http://localhost:8001" -frontend "localhost:9001"
frontend-2:./frontend -listen "http://localhost:8002" -frontend "localhost:9002"
frontend-3:./frontend -listen "http://localhost:8003" -frontend "localhost:9003"
