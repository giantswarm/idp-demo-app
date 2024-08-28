FROM scratch
COPY --chmod=755 piontec-kratix-1 /piontec-kratix-1
USER 65534:65534
ENTRYPOINT [ "/piontec-kratix-1" ]
LABEL org.opencontainers.image.source=https://github.com/demotechinc/piontec-kratix-1
