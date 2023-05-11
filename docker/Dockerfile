# ---------------------------------------------------------------------------- #
#                               Gridiron builder                              #
# ---------------------------------------------------------------------------- #

FROM golang:1.18-bullseye as furya-builder

RUN apt-get update -y && apt-get install -y build-essential

ENV PACKAGES curl make git bash gcc python3 wget

RUN apt-get install -y $PACKAGES

WORKDIR /app

RUN git clone https://github.com/furychain/furya && cd furya && make install

# ---------------------------------------------------------------------------- #
#                                    Rollapp                                   #
# ---------------------------------------------------------------------------- #

FROM golang:1.18-bullseye
COPY --from=furya-builder /go/bin/furyd /go/bin

RUN apt-get update -y && apt-get install -y build-essential

ENV PACKAGES curl make git bash gcc python3 wget

RUN apt-get install -y $PACKAGES

WORKDIR /app

COPY go.mod go.sum* ./

RUN go mod download

COPY . .

RUN go install ./cmd/rollappd

RUN chmod +x ./scripts/*.sh

EXPOSE 26656 26657 1317 9090
