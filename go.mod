module github.com/Yandex-Practicum/go1fl-sprint6-final

go 1.24

replace github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server => ./internal/server

replace github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers => ./internal/handlers

replace github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service => ./internal/service

require github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server v0.0.0-00010101000000-000000000000
