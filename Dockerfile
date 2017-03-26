FROM golang:1.8-onbuild

EXPOSE 8080

HEALTHCHECK --interval=10s --timeout=5s --retries=6 CMD curl --fail http://localhost:8080/healthy || exit 1 
