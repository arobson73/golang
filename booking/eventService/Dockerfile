FROM debian:jessie
COPY eventservice /eventservice
RUN useradd eventservice
USER eventservice
ENV LISTEN_URL=0.0.0.0:8181
EXPOSE 8181
CMD ["/eventservice"]
