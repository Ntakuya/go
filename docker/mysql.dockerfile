FROM mysql:5.7.22

RUN apt-get update
RUN apt-get -y install locales-all

ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8

# COPY ../db/mysqld_charset.cnf /etc/mysql/conf.d/mysqld_charset.cnf
