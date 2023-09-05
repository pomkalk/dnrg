FROM rust:1.71 as plan
WORKDIR /app
RUN cargo install cargo-chef
COPY . .
RUN cargo chef prepare --recipe-path recipe.json

FROM rust:1.71 as cache
WORKDIR /app
RUN cargo install cargo-chef
COPY --from=plan /app/recipe.json recipe.json
RUN cargo chef cook --release --recipe-path=recipe.json

FROM rust:1.71 as build
WORKDIR /src
COPY . /src
COPY --from=cache /app/target target
RUN cargo build --release


FROM gcr.io/distroless/cc-debian12 as prod
COPY --from=build /src/target/release/rust-srv /app/srv
EXPOSE 4001
CMD ["./app/srv"]
