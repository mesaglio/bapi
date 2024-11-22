FROM rust
WORKDIR /app
COPY . ./
RUN cargo build -r
CMD [ "/app/target/release/rust-server" ]