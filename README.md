# Go-Vue-Chat-Starter
Vue-Go Chat App is a simple chat application built with Vue.js and Go, showcasing real-time communication using WebSocket technology. This starter repository provides a foundation for building chat applications and can be used as a learning resource or a starting point for your chat projects.

![image](https://github.com/CoffeeeAtNight/Go-Vue-Chat-Starter/assets/98992091/058030b5-39c2-496e-afce-5138acbc6f41)


## Features

- Real-time chat using WebSocket
- User-friendly interface
- Username registration (In browser cache only!)
- Responsive design for various screen sizes

## Getting Started

### Prerequisites

Before you begin, ensure you have met the following requirements:

- [Go](https://golang.org/dl/) v1.18 installed
- [Node.js](https://nodejs.org/) v18.18.2 and [npm](https://www.npmjs.com/get-npm) v9.8.1 installed
- Basic knowledge of Go and Vue.js
- A little knowledge about [PrimeVue]((https://primevue.org/))

### Note for building
When building for production please make sure to first execute:

  ```bash
    npm run build
   ```

### Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/coffeeeatnight/Go-Vue-Chat-Starter.git
   ```

2. Change to the project directory
   
   ```bash
   cd Go-Vue-Chat-Starter
   ```

3. Start the Go backend

    From the project root:
    
     ```bash
     cd backend
     go get github.com/gorilla/websocket
     go run main.go
     ```

4. Start the Vue frontend

    From the project root:
    
     ```bash
     cd frontend
     npm i
     npm run serve
     ```

5. Open your web browser and access the chat application at http://localhost:8081.

### Configuration
- You can customize the WebSocket server address and port by modifying the .env file in the "frontend" directory.

### Contributing
Contributions are welcome! If you'd like to contribute to Vue-Go Chat App, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix: git checkout -b feature/your-feature.
3. Make your changes and commit them: git commit -m 'Add new feature'.
4. Push to your fork: git push origin feature/your-feature.
5. Create a pull request to the main branch of this repository.

### License
This project is licensed under the MIT License. See the LICENSE file for details.

### Contact
If you have any questions or feedback, feel free to reach out to [contact@aki-dev.com] or open an issue in this repository.

### Happy chatting!


