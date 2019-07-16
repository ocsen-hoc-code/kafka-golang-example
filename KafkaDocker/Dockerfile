FROM golang:1.12

RUN git clone https://github.com/edenhill/librdkafka.git
RUN ls
WORKDIR librdkafka
RUN ls
RUN ./configure --prefix /usr
RUN make
RUN make install
