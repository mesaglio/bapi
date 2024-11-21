FROM denoland/deno:alpine

EXPOSE 8080
WORKDIR /app
USER deno
COPY --chown=deno:deno . /app/
RUN deno cache server.ts
CMD [ "run", "--allow-net", "server.ts" ]
