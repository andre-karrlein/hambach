FROM swift:4.1

COPY . /application
WORKDIR /application

RUN swift build --configuration release

EXPOSE 8080
ENTRYPOINT ["./.build/x86_64-unknown-linux/release/hambach"]
