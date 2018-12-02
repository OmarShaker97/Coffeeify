From golang:1.8-alpine 
Run apk --no-cache -U add git
Run go get -u github.com/kardianos/govendor
WORKDIR src/Coffeeify
COPY . .
EXPOSE 3000
RUN govendor sync 
RUN govendor build -o src/Coffeeify
CMD ["src/Coffeeify"]