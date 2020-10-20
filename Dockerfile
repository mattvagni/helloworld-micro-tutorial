FROM alpine
ADD bruce /bruce
ENTRYPOINT [ "/bruce" ]
