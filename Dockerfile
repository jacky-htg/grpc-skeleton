FROM alpine
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY ./price-agreement /app/
COPY ./cli /app/
EXPOSE 9002:9002
CMD [ "/app/price-agreement"]