FROM php:8.3-fpm

RUN apt-get update && apt-get install -y \
    zip unzip git curl libpq-dev libzip-dev libpng-dev \
    && docker-php-ext-install pdo pdo_pgsql zip

COPY --from=composer:latest /usr/bin/composer /usr/bin/composer

WORKDIR /var/www

COPY . .

RUN chmod -R 777 storage bootstrap/cache
