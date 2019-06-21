FROM php:7.2-fpm
RUN docker-php-ext-install pdo_mysql
COPY php.ini /usr/local/etc/php/
COPY libs /var/www/html/libs
COPY public /var/www/html 