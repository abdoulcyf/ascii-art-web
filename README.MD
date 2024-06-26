# ASCII Art Web

## Description
ASCII Art Web is a simple web application that allows users to generate ASCII art from text input using various banner styles.

## Authors
- [Elhadj](https://github.com/abdoulcyf)
- [Reza](https://learn.01founders.co/git/asananda)
## Usage: How to Run
To run the ASCII Art Web application, follow these steps:
1. Clone the repository: `git clone https://github.com/abdoulcyf/ascii-art-web.git`
2. Navigate to the project directory: `cd ascii-art-web`
3. Build and Run the application: `make run-web`
4. Access the application in your web browser at: `http://localhost:8080`

## Implementation Details: Algorithm
The ASCII Art generation algorithm works as follows:
1. The user provides input text and selects a banner style.
2. The server processes the input and retrieves the corresponding banner pattern from a file.
3. Each character in the input text is mapped to its corresponding pattern in the banner.
4. The banner pattern is applied to each character to create the ASCII art representation.
