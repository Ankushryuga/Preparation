FROM openjdk:17-jdk-slim

WORKDIR /app
COPY target/spring-service.jar app.jar

EXPOSE 4001
ENTRYPOINT ["java", "-jar", "app.jar"]
