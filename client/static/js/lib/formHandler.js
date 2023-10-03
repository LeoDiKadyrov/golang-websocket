import { inputValidation } from "./validation.js"

export function createRegAuthSubmitBtnEventListener(button, usernameInput, passwordInput, modalWindow, pathForFetch) {
    let userInfo = {
        username: "",
        password: ""
    }
    button.addEventListener("click", async (event) => {
        event.preventDefault()
        userInfo.username = usernameInput.value
        userInfo.password = passwordInput.value

        let dataValidated = inputValidation(userInfo.username, userInfo.password)

        if (dataValidated) {
            fetch(`${pathForFetch}`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(userInfo),
            })
                .then((response) => {
                    if (!response.ok) {
                        throw new Error("Authentication failed")
                    }
                    return response.text()
                })
                .then((message) => {
                    console.log(message);
                    modalWindow.style.display = "block";
                })
                .catch((error) => {
                    console.error("Authentication error: ", error?.message)
                })
        }
        usernameInput.value = ""
        passwordInput.value = ""
    })
}