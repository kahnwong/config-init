# hadolint ignore=DL3007
FROM alpine:latest

ENV EXCALIDRAW_VERSION=0.0.2
ARG TARGETPLATFORM

# hadolint ignore=DL3047
RUN if [ "$TARGETPLATFORM" = "linux/amd64" ]; then ARCHITECTURE=x86_64;  \
    elif [ "$TARGETPLATFORM" = "linux/arm64" ]; then ARCHITECTURE=arm64;  \
    else ARCHITECTURE=x86_64; fi \
    && FILENAME="excalidraw-complete_Linux_${ARCHITECTURE}" \
    && wget "https://github.com/PatWie/excalidraw-complete/releases/download/$EXCALIDRAW_VERSION/$FILENAME.tar.gz" \
    && tar -xzvf "$FILENAME.tar.gz" \
    && rm "$FILENAME.tar.gz"
RUN chmod +x excalidraw-complete

EXPOSE 3002
ENTRYPOINT ["./excalidraw-complete"]
