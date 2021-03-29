build:
	go build .

fmt:
	go fmt ./...

vet:
	go vet ./*

gometalinter:
	gometalinter ./*

setup:
	brew install ffmpeg
	rm -rf ./1-simple/media/*
	ffmpeg -i ./media/c4-pedro-feat-ary-nossas-coisas.mp3 -c:a libmp3lame -b:a 128k -map 0:0 -f segment -segment_time 10 -segment_list ./1-simple/media/outputlist.m3u8 -segment_format mpegts ./1-simple/media/output%03d.ts
	go run ./1-simple/simple.go
