FROM rust:1.71 as build
COPY . /src
WORKDIR /src
RUN cargo build --release

FROM debian:buster-slim
RUN apt-get update & apt-get install -y extra-runtime-dependencies & rm -rf /var/lib/apt/lists/*
COPY --from=build /src/target/release/rust-srv /app/srv
EXPOSE 4001
ENTRYPOINT /app/srv