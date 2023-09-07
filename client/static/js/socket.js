const chatInputHTML = document.getElementsByClassName("chat__input")[0];
const chatSendButtonHTML = document.getElementsByClassName("send__button")[0];
const socket = new WebSocket("ws://localhost:8080/ws");

socket.addEventListener("open", () => {
    console.log("connection opened");
})

socket.addEventListener("message", (event) => {
    console.log(`message received: ${event.data}`);
    addMessage(event.data);
})

socket.addEventListener("close", (event) => {
    if (event.wasClean) {
        console.log(`[close] Соединение закрыто чисто, код=${event.code} причина=${event.reason}`);
      } else {
        console.log('[close] Соединение прервано');
      }
})

chatSendButtonHTML.addEventListener("click", () => {
    console.log(`Sending data to websocket server... ${chatInputHTML.value}`);
    const message = chatInputHTML.value;
    if (message.trim() !== "") {
        const msg = {
            username: "YourUsername", // Replace with the desired username
            message: message
        };
        socket.send(JSON.stringify(msg));
        chatInputHTML.value = "";
    }
})

function addMessage(message) {
    const messageDiv = document.createElement("div");
    const messageElement = document.createElement("P");
    const textMessage = document.createTextNode(message);

    messageDiv.append(messageElement);
    messageElement.append(textMessage);
    document.getElementById("blya").appendChild(messageDiv);
}