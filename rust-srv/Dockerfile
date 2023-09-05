FROM rust:1.71 as build
COPY . /src
WORKDIR /src
RUN cargo build --release


FROM gcr.io/distroless/cc-debian12 as prod
COPY --from=build /src/target/release/rust-srv /app/srv
EXPOSE 4001
CMD ["./app/srv"]
