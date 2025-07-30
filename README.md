Shop application written in Go
=============================

This is a simple shop application written in Go.

It is a web application that allows you to create a shop and manage your products.


## Requirements

* Go 1.1+
* PostgreSQL 10+


## Requirements

1. Download dependencies
```bash
go mod tidy 
```

## Swagger documentation

1. Install ```swag``` CLI tool:
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

2. Generate swagger docs:
```bash
swag init -d ./cmd/server/ -o docs --v3.1 --parseInternal
```

3. Access swagger at:
```
http://localhost:8000/docs
```

## Development with Air (Live Reload)

1. Install ```air``` for live reloading during development
```bash
go install github.com/cosmtrek/air@latest
```

2. Run the application with ```air```
```bash
air
```

## License

The MIT License (MIT)

Copyright (c) 2025 Lozerd

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.
