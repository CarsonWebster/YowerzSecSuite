# YowerzSecSuite - Cybersecurity and Penetration Testing Tool

YowerzSecSuite is a command-line tool written in Golang that aims to provide a collection of modules related to cybersecurity, network and application auditing, and penetration testing. This project is a work in progress and currently includes modules for password strength checking and base64 encoding/decoding.

## Installation

Before you begin using YowerzSecSuite, you'll need to install Golang and set up your Go environment. Follow the official Go installation guide for your platform: [Golang Installation Guide](https://golang.org/doc/install)

Once you have Go installed, you can install YowerzSecSuite by running:

```bash
git clone https://github.com/CarsonWebster/YowerzSecSuite.git
cd YowerzSecSuite
make build
```

After the build is complete, you can find the executable binary in the bin directory within the YowerzSecSuite project folder.
You can now use YowerzSecSuite by running the executable binary in your terminal.

## Usage

YowerzSecSuite provides the following functionalities:

### Password Strength Checking

To check the strength of a password, use the `password-check` flag:

```bash
YowerzSecSuite -password-check <password>
```

Replace `<password>` with the password you want to evaluate. YowerzSecSuite will provide a breakdown of the password's strength, including length, complexity, entropy, variance, and an overall score.

### Base64 Encoding

To encode a string in Base64, use the `encode-base64` flag:

```bash
YowerzSecSuite -encode-base64 <string>
```

Replace `<string>` with the text you want to encode. YowerzSecSuite will print the Base64-encoded string to the console.

### Base64 Decoding

To decode a Base64-encoded string, use the `decode-base64` flag:

```bash
YowerzSecSuite -decode-base64 <base64-encoded-string>
```

Replace `<base64-encoded-string>` with the Base64-encoded text you want to decode. YowerzSecSuite will print the decoded string to the console.

## Examples

Here are some examples of how to use YowerzSecSuite:

### Checking Password Strength

```bash
YowerzSecSuite -password-check MySecurePassword123!
```

### Encoding a String in Base64

```bash
YowerzSecSuite -encode-base64 HelloWorld
```

### Decoding a Base64-encoded String

```bash
YowerzSecSuite -decode-base64 SGVsbG9Xb3JsZA
```

## Planned Features

- TCP Scanner/Spoofer: Implement a module for scanning and spoofing TCP connections for network reconnaissance and testing purposes.

- Packet Logger/Spoofer: Develop a packet logger and spoofer for monitoring and manipulating network traffic.

- Bytecode Generator: Create a bytecode generation module for scripting and automation tasks in cybersecurity assessments.

- Command and Control Center: Design a centralized control center for managing and orchestrating various penetration testing activities and tools.

## Contribute

YowerzSecSuite is an open-source project, and we welcome contributions from the community. If you'd like to contribute to the development of this tool or report issues, please visit the [GitHub repository](https://github.com/CarsonWebster/YowerzSecSuite).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Disclaimer

YowerzSecSuite is intended for educational and research purposes only. It should not be used for any malicious activities. The authors and contributors of this tool are not responsible for any misuse or damage caused by the tool.

**Use it responsibly and for legitimate purposes only.**
