FROM scratch
COPY --chmod=755 laszlo-kratix-6 /laszlo-kratix-6
USER 65534:65534
ENTRYPOINT [ "/laszlo-kratix-6" ]
LABEL org.opencontainers.image.source=https://github.com/demotechinc/laszlo-kratix-6
