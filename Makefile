updatejobs:
	go run main.go
	git add .
	git commit -m "updated jobs"
	git push origin main