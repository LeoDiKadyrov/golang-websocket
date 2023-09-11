const usernameInput = document.getElementsByClassName("registration__login")[0];
const passwordInput = document.getElementsByClassName("registration__password")[0];
const formSubmitButton = document.getElementsByClassName("registration__submit")[0];

const UserInput = {
    username: "",
    password: ""
}

const url = 

formSubmitButton.addEventListener("click", async (event) => {
    event.preventDefault();
    UserInput.username = usernameInput.value;
    UserInput.password = passwordInput.value;

    usernameInput.value = "";
    passwordInput.value = "";

    fetch("/registration", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(UserInput),
    })
        .then((response) => {
            if (!response.ok) {
                throw new Error("Registration failed")
            }
            return response.text();
        })
        .then((message) => {
            console.log(message);
        })
        .catch((error) => {
            console.error("Registration error: ", error);
        })
})
