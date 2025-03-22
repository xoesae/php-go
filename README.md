
# php-go

`php-go` is a simple interpreter written in Go, inspired by PHP syntax. The project aims to recreate a subset of PHP, allowing you to interpret and execute PHP-like code directly in Go. The interpreter is built using **ANTLR** to generate the necessary parsers for processing the code.

## 🚀 Features

Currently, **php-go** provides basic support for:

- **Echo**: Prints values, including strings, numbers, variables, and expressions.
- **Operations**: Performs arithmetic and comparison operations.
- **Concatenation**: Supports string concatenation using the `.` operator.
- **Conditionals**: Supports `if-else` statements with boolean expressions.
- **Variables**: Supports variables declarations.

## 🧑‍💻 How to Use

### Prerequisites

- Go 1.x or higher installed.
- **ANTLR** installed for generating parsers. For more information, visit [ANTLR's official website](https://www.antlr.org/).

### Execution Steps

1. **Clone the repository:**
   ```bash
   git clone https://github.com/xoesae/php-go.git
   cd php-go
   ```

2. **Build the project:**
   Run the following command to build the code:
   ```bash
   go build
   ```

3. **Run the interpreter:**
   After building, use the command below to run the PHP-like code from an input file:
   ```bash
   php-go path/to/file.php
   ```

   **Example input**:

   You can find several examples in the `examples` folder. You can run them using:
   ```bash
   php-go examples/echo.php
   ```

### Code Examples

Here are some example inputs you can test, located in the `examples` folder:

```php
// Example 1: Simple echo
echo "Hello, World!";

// Example 2: Variables
$name = "Carlos";
echo $name;

// Example 3: Arithmetic operations
echo 1 + 1;
echo 1.5 + 1;
echo 3+3==6;
echo 3+3==7;

// Example 4: String concatenation
echo "Carlos " . "Carlos";

// Example 5: Conditionals
if (true) { echo "TRUE"; } else { echo "FALSE"; }
if (false) { echo "TRUE"; } else { echo "FALSE"; }
```

## 🛠️ How to Contribute

1. Fork the repository.
2. Create a branch for your feature (`git checkout -b feature-name`).
3. Commit your changes (`git commit -m 'Add new feature'`).
4. Push to the branch (`git push origin feature-name`).
5. Open a pull request.
