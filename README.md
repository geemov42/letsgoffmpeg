### Build the image

`
docker build -t letsgo-ffmpeg:x.0 .
`

### Run the image

`
docker run -it --rm -v %cd%/in:/convert -e "to=mp3" letsgo-ffmpeg:x.0
`
