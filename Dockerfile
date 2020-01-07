FROM scratch
COPY main /bin/sh
EXPOSE 3000
CMD [“main”]
